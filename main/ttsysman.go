package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (i *Trans) newS3Object() {
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"), // Replace with your AWS region
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Create an S3 client
	s3Client := s3.New(sess)

	fbytes := []byte(i.HTTP.RequestBody)
	// Create an S3 object
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(i.EnvVars.S3_PUBLISH_FILES),
		Key:    aws.String(i.Query.Name),
		Body:   bytes.NewReader(fbytes),
	})
	if err != nil {
		log.Println(err.Error())
	}
}

func (i *Trans) newCodeMap() {
	switch i.Query.Action {
	case XDW_OPERATION_EDIT_CODEMAP:
		i.IdMaps = updateUserTerminology(i.Query.User, GetIntFromString(i.Query.Id), i.Query.Lid, i.Query.Mid)
	case XDW_OPERATION_INSERT_CODEMAP:
		i.IdMaps = newUserTerminology(i.Query.User, i.Query.Lid, i.Query.Mid)
	case XDW_OPERATION_DELETE_CODEMAP:
		i.IdMaps = deleteUserTerminology(i.Query.User, GetIntFromString(i.Query.Id))
	}
	if i.HTTP.ReqContentType == APPLICATION_JSON || i.Query.Format == APPLICATION_JSON {
		bodybytes, _ := json.Marshal(i.IdMaps)
		i.HTTP.ResponseBody = string(bodybytes)
	} else {
		if i.Query.Template == "" {
			i.Query.Template = "codemaps2t_tmplt"
		}
		i.setResponseFromTemplate()
	}
}
func (i *Trans) persistDefinitions() {
	if xdwconfigFiles, err := GetFolderFiles("./xdwconfig/"); err == nil {
		for _, file := range xdwconfigFiles {
			if strings.HasSuffix(file.Name(), ".json") {
				filebytes := loadFile(file, "./xdwconfig/")
				if len(filebytes) > 0 {
					i.HTTP.RequestBody = string(filebytes)
					log.Printf("Registering %s", file.Name())
					if strings.HasSuffix(strings.Split(file.Name(), ".")[0], "_def") {
						i.newXDWDefinition()
					} else {
						if strings.HasSuffix(strings.Split(file.Name(), ".")[0], "_meta") {
							i.newXDWMETA()
						} else {
							log.Printf("Invalid Definition filename - %s", file.Name())
						}

					}
				}
			}
		}
	} else {
		log.Println(err.Error())
	}
}
func (i *Trans) newXDWDefinition() {
	var newexpressions = make(map[string]string)
	i.XDWState.Definition = WorkflowDefinition{}
	if i.Error = json.Unmarshal([]byte(i.HTTP.RequestBody), &i.XDWState.Definition); i.Error == nil {
		log.Printf("Marshalled %s XDW Definition", i.XDWState.Definition.Ref)
		i.Query.Pathway = i.XDWState.Definition.Ref
		log.Printf("Processing %s", i.XDWState.Definition.Ref)
		for _, task := range i.XDWState.Definition.Tasks {
			for _, inp := range task.Input {
				log.Printf("Processing Task %v Task Name %s Input Part Name %s", task.ID, task.Name, inp.Name)
				if inp.AccessType == XDS_REGISTERED {
					newexpressions[inp.Name] = i.XDWState.Definition.Ref
					log.Printf("Task %v Task Name %s Input Part %s included in potential DSUB Broker subscriptions", task.ID, task.Name, inp.Name)
				}
			}
			for _, out := range task.Output {
				log.Printf("Checking Task %v %s output %s", task.ID, task.Name, out.Name)
				if out.AccessType == XDS_REGISTERED {
					newexpressions[out.Name] = i.XDWState.Definition.Ref
					log.Printf("Task %v %s output %s included in potential DSUB Broker subscriptions", task.ID, task.Name, out.Name)
				}
			}
		}
		if len(newexpressions) > 0 {
			var expressions []string
			for expression := range newexpressions {
				expressions = append(expressions, expression)
			}
			log.Printf("Creating %v DSUB Broker Subscriptions for Document Type Codes %v", len(expressions), expressions)
			i.XDWState.Expressions = expressions
			i.newBrokerSubscriptions()
		}
	}
	i.Query.Name = i.XDWState.Definition.Ref
	i.persistWorkflowDefinition(string(i.HTTP.RequestBody), false)
	i.HTTP.ResponseBody = i.HTTP.RequestBody
}
func (i *Trans) newXDWMETA() {
	i.XDWState.Meta = WorkflowMeta{}
	if i.Error = json.Unmarshal([]byte(i.HTTP.RequestBody), &i.XDWState.Meta); i.Error == nil {
		i.Query.Name = i.XDWState.Meta.Id
		i.persistWorkflowDefinition(string(i.HTTP.RequestBody), true)
		i.HTTP.StatusCode = http.StatusOK
		i.HTTP.ResponseBody = i.HTTP.RequestBody
	}
}
func (i *Trans) setWorkflowDocument() {
	vers := GetIntFromString(i.Query.Vers)
	if vers < 1 {
		vers = 1
	}
	wfs := Workflows{Action: SELECT}
	wf := Workflow{Pathway: i.Query.Pathway, NHSId: i.Query.Nhs, Version: vers}
	wfs.Workflows = append(wfs.Workflows, wf)
	log.Printf("Searching for Workflows for Pathway=" + i.Query.Pathway + " NHS=" + i.Query.Nhs + " Version=" + GetStringFromInt(vers))
	if i.Error = wfs.newEvent(); i.Error == nil {
		if wfs.Count == 1 {
			i.Error = json.Unmarshal([]byte(wfs.Workflows[1].XDW_Doc), &i.XDWState.WorkflowDocument)
			i.Error = json.Unmarshal([]byte(wfs.Workflows[1].XDW_Def), &i.XDWState.Definition)
			i.setWorkflowXDSMeta()
		} else {
			i.Error = errors.New("no workflow found for Pathway " + i.Query.Pathway + " NHS " + i.Query.Nhs + " Version " + GetStringFromInt(vers))
			log.Println(i.Error.Error())
		}
	} else {
		log.Println(i.Error.Error())
	}
}
func (i *Trans) setWorkflowDefinition() {
	if i.XDWState.Definition.Ref == "" {
		log.Println("No Definition currently Loaded")
		xdws := XDWS{Action: SELECT}
		xdw := XDW{Name: i.Query.Pathway + "_def"}
		log.Printf("Retrieving XDW Definition %s", xdw.Name)
		xdws.XDW = append(xdws.XDW, xdw)
		if i.Error = xdws.newEvent(); i.Error == nil {
			if xdws.Count == 1 {
				i.Error = json.Unmarshal([]byte(xdws.XDW[1].XDW), &i.XDWState.Definition)
			} else {
				i.Error = errors.New("no xdw registered for " + xdw.Name)
			}
		}
	}
}
func (i *Trans) setWorkflowXDSMeta() {
	if i.XDWState.Meta.Id == "" {
		xdws := XDWS{Action: SELECT}
		xdw := XDW{Name: i.Query.Name + "_meta", IsXDSMeta: true}
		xdws.XDW = append(xdws.XDW, xdw)
		if i.Error = xdws.newEvent(); i.Error == nil {
			if xdws.Count == 1 {
				i.Error = json.Unmarshal([]byte(xdws.XDW[1].XDW), &i.XDWState.Meta)
			} else {
				i.Error = errors.New("no xdw meta registered for " + i.Query.Name)
			}
		}
	}
}
func (i *Trans) persistWorkflowDefinition(config string, isxdsmeta bool) {
	log.Printf("Persisting %s", i.Query.Name)
	xdws := XDWS{Action: DELETE}
	xdw := XDW{Name: i.Query.Name, IsXDSMeta: isxdsmeta}
	xdws.XDW = append(xdws.XDW, xdw)
	xdws.newEvent()
	xdws = XDWS{Action: INSERT}
	xdw = XDW{Name: i.Query.Name, IsXDSMeta: isxdsmeta, XDW: config}
	xdws.XDW = append(xdws.XDW, xdw)
	i.Error = xdws.newEvent()
	if i.Error != nil {
		log.Println(i.Error.Error())
		i.setError()
	} else {
		log.Printf("Persisted %s XDW Definition", i.Query.Name)
	}
}
func (i *Trans) newTemplate() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		i.Error = persistTemplate(i.Query.User, i.Query.Template, string(i.HTTP.RequestBody))
	} else {
		var files []fs.DirEntry
		if files, i.Error = GetFolderFiles("./templates/"); i.Error == nil {
			for _, file := range files {
				if i.Error == nil && strings.Contains(file.Name(), "_tmplt.") {
					var filebytes []byte
					if filebytes, i.Error = GetFileBytes("./templates/" + file.Name()); i.Error == nil {
						i.Error = persistTemplate(i.Query.User, strings.Split(file.Name(), ".")[0], string(filebytes))
					}
				}
			}
		}
	}
	if i.Error == nil {
		i.HTTP.ResponseBody = i.HTTP.RequestBody
	}
}
func (i *Trans) newImage() {
	statics := Statics{Action: INSERT}
	static := Static{Name: i.Query.Name, Content: string(i.HTTP.RequestBody)}
	statics.Static = append(statics.Static, static)
	if i.Error = statics.newEvent(); i.Error == nil {
		i.HTTP.ResponseBody = i.HTTP.RequestBody
	}
}

func (i *Trans) getUserMappings() {
	i.IdMaps = getUserTerminology(i.Query.User)
	i.Query.Template = "codemaps2t_tmplt"
	i.setResponseFromTemplate()

}
func updateUserTerminology(user string, id int, lid string, mid string) IdMaps {
	idmaps := IdMaps{Action: UPDATE}
	idmap := IdMap{Id: id, Lid: strings.TrimSpace(lid), Mid: strings.TrimSpace(mid)}
	idmaps.LidMap = append(idmaps.LidMap, idmap)
	idmaps.newEvent()
	return getUserTerminology(user)
}
func newUserTerminology(user string, lid string, mid string) IdMaps {
	idmaps := IdMaps{Action: INSERT}
	idmap := IdMap{Lid: strings.TrimSpace(lid), Mid: strings.TrimSpace(mid), User: user}
	idmaps.LidMap = append(idmaps.LidMap, idmap)
	idmaps.newEvent()
	return getUserTerminology(user)
}
func deleteUserTerminology(user string, id int) IdMaps {
	idmaps := IdMaps{Action: DELETE}
	idmap := IdMap{Id: id}
	idmaps.LidMap = append(idmaps.LidMap, idmap)
	idmaps.newEvent()
	return getUserTerminology(user)
}
func getUserTerminology(user string) IdMaps {
	idmaps := IdMaps{Action: SELECT}
	idmap := IdMap{User: user}
	idmaps.LidMap = append(idmaps.LidMap, idmap)
	if err := idmaps.newEvent(); err != nil {
		log.Println(err.Error())
	}
	return idmaps
}
func (i *Trans) getImage() {
	if i.Query.Name == "" {
		i.Query.Name = "favicon.ico"
	}
	log.Printf("Retrieving %s", i.Query.Name)
	bodybytes := i.getStaticContent()
	if i.Query.Name == "favicon.ico" {
		i.HTTP.RspContentType = "image/x-icon"
		decodedBytes, err := base64.StdEncoding.DecodeString(string(bodybytes))
		if err != nil {
			log.Printf("Error decoding favicon.ico - %s:", err.Error())
			return
		}
		i.HTTP.ResponseBody = string(decodedBytes)
	} else {
		i.HTTP.RspContentType = "image/png"
		i.HTTP.IsBase64 = true
		i.HTTP.ResponseBody = string(bodybytes)
	}
}
func (i *Trans) getStaticContent() []byte {
	if i.Query.Name == "" {
		i.Query.Name = "event-logo.png"
	}
	statics := Statics{Action: SELECT}
	static := Static{Name: i.Query.Name}
	statics.Static = append(statics.Static, static)
	statics.newEvent()
	if statics.Count == 1 {
		return []byte(statics.Static[1].Content)
	}
	return []byte("")
}
func (i *Trans) newUploadFileForm() {
	i.Query.Template = "uploadfile2t_tmplt"
	i.setResponseFromTemplate()
}

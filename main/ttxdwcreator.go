package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"strings"
)

// func (i *Trans) newDocument() {
// 	log.Printf("Publishing %s Workflow Document for %s Expression %s", i.Query.Pathway, i.Query.Nhs, i.Query.Expression)
// 	var httpReq = HTTPRequest{Method: http.MethodPost, URL: i.EnvVars.SCR_URL}
// 	var req []byte

//		if i.Error = json.Unmarshal([]byte(i.HTTP.RequestBody), &req); i.Error == nil {
//			i.persistWorkflowDefinition(string(i.HTTP.RequestBody), true)
//			i.HTTP.StatusCode = http.StatusOK
//			i.HTTP.ResponseBody = i.HTTP.RequestBody
//		}
//	}
func (i *Trans) newXDWCreator() {
	var bb []byte
	log.Printf("Checking for OPEN %s Workflow for NHS %s", i.Query.Pathway, i.Query.Nhs)
	wfs, _ := GetWorkflows(i.Query.Pathway, i.Query.Nhs, 1, STATUS_OPEN)
	if wfs.Count == 1 {
		log.Println("Found Open Workflow")
		json.Unmarshal([]byte(wfs.Workflows[1].XDW_Doc), &i.XDWState.WorkflowDocument)
	} else {
		log.Printf("Creating New %s Workflow for Nhs ID %s", i.Query.Pathway, i.Query.Nhs)
		i.Query.Name = i.Query.Pathway
		i.setWorkflowDefinition()
		if i.Error == nil {
			i.setWorkflowXDSMeta()
			if i.Error == nil {
				i.deprecateWorkflows()
				if i.Error == nil {
					i.deprecateEvents()
					if i.Error == nil {
						i.createWorkflow()
						if i.Error == nil {
							i.persistWorkflow()
						} else {
							log.Println(i.Error.Error())
						}
					} else {
						log.Println(i.Error.Error())
					}
				} else {
					log.Println(i.Error.Error())
				}
			} else {
				log.Println(i.Error.Error())
			}
		} else {
			log.Println(i.Error.Error())
		}
	}
	i.XDWState.Events, i.Error = getEvents(i.Query.Pathway, i.Query.Nhs, 1, -1)
	if i.Error != nil {
		i.setError()
	}
	switch i.HTTP.RspContentType {
	case APPLICATION_XML:
		bb, i.Error = xml.MarshalIndent(i.XDWState.WorkflowDocument, "", "  ")
		i.HTTP.ResponseBody = string(bb)
	case APPLICATION_JSON:
		bb, i.Error = json.MarshalIndent(i.XDWState.WorkflowDocument, "", "  ")
		i.HTTP.ResponseBody = string(bb)
	default:
		i.Query.Template = "xdw2t_tmplt"
		i.setResponseFromTemplate()
	}
}
func (i *Trans) createWorkflow() {
	var authoid = GetLocalValue(i.Query.User, i.Query.Org)
	var wfid = Newid()
	var effectiveTime = Time_Now()
	var confcode = i.XDWState.Definition.Confidentialitycode
	i.XDWState.WorkflowDocument.Xdw = XDWNameSpace
	i.XDWState.WorkflowDocument.Hl7 = HL7NameSpace
	i.XDWState.WorkflowDocument.WsHt = WHTNameSpace
	i.XDWState.WorkflowDocument.Xsi = XMLNS_XSI
	i.XDWState.WorkflowDocument.XMLName.Local = XDWNameLocal
	i.XDWState.WorkflowDocument.SchemaLocation = WorkflowDocumentSchemaLocation
	i.XDWState.WorkflowDocument.ID.Root = strings.ReplaceAll(WorkflowInstanceId, "^", "")
	i.XDWState.WorkflowDocument.ID.Extension = wfid
	i.XDWState.WorkflowDocument.ID.AssigningAuthorityName = "NHS"
	i.XDWState.WorkflowDocument.EffectiveTime.Value = effectiveTime
	log.Printf("Set Workflow Creation Time %s", i.XDWState.WorkflowDocument.EffectiveTime.Value)
	i.XDWState.WorkflowDocument.ConfidentialityCode.Code = confcode
	i.XDWState.WorkflowDocument.Patient.Root = NHS_OID_DEFAULT
	i.XDWState.WorkflowDocument.Patient.Extension = i.Query.Nhs
	i.XDWState.WorkflowDocument.Patient.AssigningAuthorityName = "NHS"
	i.XDWState.WorkflowDocument.Author.AssignedAuthor.ID.Root = authoid
	i.XDWState.WorkflowDocument.Author.AssignedAuthor.ID.Extension = strings.ToUpper(i.Query.Org)
	i.XDWState.WorkflowDocument.Author.AssignedAuthor.ID.AssigningAuthorityName = authoid
	i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Family = i.Query.User
	i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Prefix = i.Query.Role
	i.XDWState.WorkflowDocument.WorkflowInstanceId = wfid + WorkflowInstanceId
	i.XDWState.WorkflowDocument.WorkflowDocumentSequenceNumber = "1"
	i.XDWState.WorkflowDocument.WorkflowStatus = STATUS_OPEN
	i.XDWState.WorkflowDocument.WorkflowDefinitionReference = strings.ToUpper(i.Query.Pathway)
	i.Query.Eventtype = XDW_OPERATION_CREATE_WORKFLOW
	i.Query.Expression = i.XDWState.WorkflowDocument.WorkflowDefinitionReference
	i.Query.Topic = XDW_OPERATION_WORKFLOW
	i.Query.Taskid = "0"
	i.Query.Vers = "1"
	i.newXDWEvent()
	docevent := DocumentEvent{}
	docevent.Author = i.Query.User + " " + i.Query.Org + " " + i.Query.Role
	docevent.TaskEventIdentifier = "0"
	docevent.EventTime = effectiveTime
	log.Printf("Set XDW DocumentEvent Time %s", docevent.EventTime)
	docevent.EventType = i.Query.Eventtype
	docevent.ActualStatus = STATUS_OPEN
	i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent = append(i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent, docevent)
	i.Query.Eventtype = XDW_OPERATION_CREATE_TASK
	for _, t := range i.XDWState.Definition.Tasks {
		i.Query.Taskid = t.ID
		i.Query.Expression = t.Name
		i.newXDWEvent()
		log.Printf("Creating XDW Task Id - %v Name - %s", t.ID, t.Name)
		task := XDWTask{}
		task.TaskData.TaskDetails.ID = t.ID
		task.TaskData.TaskDetails.TaskType = t.Tasktype
		task.TaskData.TaskDetails.Name = t.Name
		task.TaskData.TaskDetails.ActualOwner = t.ActualOwner
		task.TaskData.TaskDetails.CreatedBy = i.Query.Role + " " + i.Query.User
		task.TaskData.TaskDetails.CreatedTime = effectiveTime
		log.Printf("Set Task Creation Time %s", task.TaskData.TaskDetails.CreatedTime)
		task.TaskData.TaskDetails.RenderingMethodExists = "false"
		task.TaskData.TaskDetails.LastModifiedTime = effectiveTime
		log.Printf("Set Task Last Modified Time %s", task.TaskData.TaskDetails.LastModifiedTime)
		task.TaskData.Description = t.Description
		task.TaskData.TaskDetails.Status = STATUS_CREATED
		for _, inp := range t.Input {
			docinput := Input{}
			docinput.Part.Name = inp.Name
			docinput.Part.AttachmentInfo.Name = inp.Name
			docinput.Part.AttachmentInfo.AccessType = inp.AccessType
			docinput.Part.AttachmentInfo.ContentType = inp.Contenttype
			docinput.Part.AttachmentInfo.ContentCategory = MEDIA_TYPES
			task.TaskData.Input = append(task.TaskData.Input, docinput)
			log.Printf("Created Input Part - %s", inp.Name)
		}
		for _, outp := range t.Output {
			docoutput := Output{}
			docoutput.Part.Name = outp.Name
			docoutput.Part.AttachmentInfo.Name = outp.Name
			docoutput.Part.AttachmentInfo.AccessType = outp.AccessType
			docoutput.Part.AttachmentInfo.ContentType = outp.Contenttype
			docoutput.Part.AttachmentInfo.ContentCategory = MEDIA_TYPES
			task.TaskData.Output = append(task.TaskData.Output, docoutput)
			log.Printf("Created Output Part - %s", outp.Name)
		}
		tev := TaskEvent{}
		tev.EventTime = effectiveTime
		log.Printf("Set Task Event Time %s", tev.EventTime)
		tev.ID = i.Query.Taskid
		tev.Identifier = t.ID
		tev.EventType = XDW_OPERATION_CREATE_TASK
		tev.Status = STATUS_COMPLETE
		task.TaskEventHistory.TaskEvent = append(task.TaskEventHistory.TaskEvent, tev)
		i.XDWState.WorkflowDocument.TaskList.XDWTask = append(i.XDWState.WorkflowDocument.TaskList.XDWTask, task)
		log.Printf("Set Workflow Task Event %s %s status to %s", t.ID, tev.EventType, tev.Status)
	}
	log.Printf("%s Created new %s Workflow for Patient %s", i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Family, i.XDWState.WorkflowDocument.WorkflowDefinitionReference, i.Query.Nhs)
}
func (i *Trans) persistWorkflow() error {
	xdwDocBytes, _ := json.MarshalIndent(i.XDWState.WorkflowDocument, "", "  ")
	xdwDefBytes, _ := json.MarshalIndent(i.XDWState.Definition, "", "  ")
	wfs := Workflows{Action: INSERT}
	wf := Workflow{Pathway: i.Query.Pathway, NHSId: i.Query.Nhs, XDW_Key: i.Query.Pathway + i.Query.Nhs, XDW_UID: i.XDWState.WorkflowDocument.ID.Extension, Version: 1, Status: STATUS_OPEN, XDW_Doc: string(xdwDocBytes), XDW_Def: string(xdwDefBytes)}
	wfs.Workflows = append(wfs.Workflows, wf)
	return wfs.newEvent()
}
func (i *Trans) deprecateWorkflows() {
	wf := Workflow{Pathway: i.Query.Pathway, NHSId: i.Query.Nhs}
	wfs := Workflows{Action: DEPRECATE}
	wfs.Workflows = append(wfs.Workflows, wf)
	i.Error = wfs.newEvent()
	log.Printf("Deprecated %v Workflows", i.XDWState.Events.LastInsertId)
}

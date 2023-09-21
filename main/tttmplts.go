package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var htmlTemplates *template.Template
var InitTemplates = true

func (i *Trans) setDocumentUploadTemplate() {
	if i.Query.Template == "" {
		log.Println("Using Default DocumentUpload template")
		i.Query.Template = "publishfile2t_tmplt"
	}
	i.setResponseFromTemplate()
}
func (i *Trans) setSpaTemplate() {
	if i.Query.Template == "" {
		log.Println("Using Default SPA template")
		i.Query.Template = "spa2t_tmplt"
	}
	i.setResponseFromTemplate()
}
func (i *Trans) pixmQueryThread(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("PIXm Worker starting")
	i.newPIXmReq()
	log.Println("PIXm Worker done")
	log.Printf("PIXm Worker Status Code %v", i.HTTP.StatusCode)
}
func (i *Trans) cglQueryThread(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("CGL Worker starting")
	i.newCglReq()
	log.Println("CGL Worker done")
	log.Printf("CGL Worker Status Code %v", i.HTTP.StatusCode)
}
func (i *Trans) setPatientTemplate() {
	i.Query.Pid = i.Query.Nhs
	i.Query.Pidoid = NHS_OID_DEFAULT
	var wg sync.WaitGroup
	wg.Add(2)
	go i.pixmQueryThread(&wg)
	go i.cglQueryThread(&wg)
	wg.Wait()
	i.newPDSReq()
	i.HTTP.StatusCode = http.StatusOK
	i.Query.Template = "allpatientSrvs_tmplt"
	i.setResponseFromTemplate()
}
func (i *Trans) setDsubAckTemplate() {
	i.Query.Template = "dsuback_tmplt"
	i.setResponseFromTemplate()
}
func (i *Trans) setDsubCancelTemplate() {
	i.Query.Template = "dsubcancel_tmplt"
	i.setResponseFromTemplate()
}
func (i *Trans) setDsubSubscribeTemplate() {
	i.Query.Template = "dsubsubscribe_tmplt"
	i.setResponseFromTemplate()
}
func (i *Trans) setNotificationTemplate() {
	i.Query.Template = "notification_tmplt"
	var tplReturn bytes.Buffer
	if InitTemplates {
		i.cacheTemplates()
	}
	if i.Error == nil {
		log.Printf("Processing Template %s", i.Query.Template)
		if i.Error = htmlTemplates.ExecuteTemplate(&tplReturn, i.Query.Template, i); i.Error == nil {
			log.Printf("Created Template \n%s", i.HTTP.ResponseBody)
			i.HTTP.ResponseBody = tplReturn.String()
			return
		}
	}
}
func (i *Trans) setResponseFromTemplate() {
	var tplReturn bytes.Buffer
	if InitTemplates {
		i.cacheTemplates()
	}
	log.Printf("Processing Template %s", i.Query.Template)

	if i.Error = htmlTemplates.ExecuteTemplate(&tplReturn, i.Query.Template, i); i.Error == nil {
		i.HTTP.ResponseBody = tplReturn.String()
		i.HTTP.RspContentType = TEXT_HTML
	} else {
		log.Println(i.Error.Error())
	}
}
func (i *Trans) cacheTemplates() {
	htmlTemplates = template.New(HTML)
	tmplts := Templates{Action: SELECT}
	if i.Error = NewDBEvent(&tmplts); i.Error == nil {
		log.Printf("Caching %v Templates", tmplts.Count)
		for _, tmplt := range tmplts.Templates {
			if tmplt.Id != 0 {
				if htmlTemplates, i.Error = htmlTemplates.New(tmplt.Name).Funcs(funcMap()).Parse(tmplt.Template); i.Error != nil {
					log.Printf("Template %s caused Error %s", tmplt.Name, i.Error.Error())
				}
			}
		}
		InitTemplates = false
	}
}
func funcMap() template.FuncMap {
	return template.FuncMap{
		"newuuid":    NewUuid,
		"newid":      Newid,
		"mappedid":   GetMappedValue,
		"prettytime": PrettyTime,
		"duration":   timeDuratipn,
		"docid":      getDocId,
	}
}

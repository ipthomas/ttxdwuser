package main

import (
	"bytes"
	"html/template"
	"log"
)

var htmlTemplates *template.Template
var InitTemplates = true

func (i *Trans) setUploadTemplate() {
	if i.Query.Template == "" {
		log.Println("Using Default Upload template")
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
func (i *Trans) setErrorTemplate() {
	if i.Query.Template == "" {
		log.Println("Using Default Error template")
		i.Query.Template = "error_tmplt"
	}
	i.setResponseFromTemplate()
}
func (i *Trans) setPatientTemplate() {
	if i.Query.Template == "" {
		log.Println("Using Default patient template")
		i.Query.Template = "allpatientSrvs_tmplt"
	}
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
	i.setResponseFromTemplate()
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
		"duration":   timeDuration,
		"docid":      getDocId,
	}
}

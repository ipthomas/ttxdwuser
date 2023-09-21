package main

import (
	"errors"
	"log"
	"strings"
)

func (i *Trans) newGetHandler() {
	log.Println("Processing new GET Request")
	i.Query, i.HTTP.RspContentType = GetQueryVars(i.Req)
	i.logRequest()
	if i.Error == nil {
		path := i.HTTP.Path
		log.Printf("Path = %s", path)
		switch {
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_TASK_STATUS):
			i.getTaskState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_TASKS_STATUS):
			i.getTasksState()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_EVENTS):
			i.getEvents()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_UPLOAD):
			i.setDocumentUploadTemplate()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_CLEAR_CACHE):
			InitTemplates = true
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_PATIENT):
			i.setPatientTemplate()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_CODEMAP):
			i.getUserMappings()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_PATHWAYS), strings.HasSuffix(path, HTTP_PATH_API_STATE_PATHWAYS):
			i.getPathways()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_EXPRESSIONS), strings.HasSuffix(path, HTTP_PATH_API_STATE_EXPRESSIONS):
			i.getExpressionsState()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_STATIC), strings.HasSuffix(path, "favicon.ico"):
			i.getImage()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_COMMENTS), strings.HasSuffix(path, HTTP_PATH_API_STATE_COMMENTS):
			i.getCommentsState()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_XDWS):
			i.setXDWs()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_XDW):
			i.setXDW()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_SPA):
			i.setSpaTemplate()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_MYSUBS):
			i.getUserSubscriptions()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_NEWSUB):
			i.newUserSubscription()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_DELSUB):
			i.cancelUserSubscription()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_CREATOR):
			i.newXDWCreator()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_TEMPLATES):
			PersistTemplates()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_DEFINITIONS):
			i.persistDefinitions()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_DEFINITION), strings.HasSuffix(path, HTTP_PATH_CONSUMER_TEMPLATE), strings.HasSuffix(path, HTTP_PATH_PUBLISH):
			i.newUploadFileForm()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE):
			i.getState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_DASHBOARD):
			i.getDashboardState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_WORKFLOWS):
			i.getWorkflowStates()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_WORKFLOW):
			i.getWorkflowState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_EVENTS):
			i.getEventStates()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_TERMINOLOGY):
			i.getTerminology()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_SUBSCRIPTIONS):
			i.getSubscriptions()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_DEFINITION):
			i.getXDWDefinition()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_META):
			i.getXDWMeta()
		default:
			i.Error = errors.New("Unsupported Consumer query path - " + i.HTTP.Path)
			i.getErrorTemplate()
		}
	}
}
func (i *Trans) getErrorTemplate() {
	i.Query.Template = "error_tmplt"
	i.Query.Act = SELECT
	i.setResponseFromTemplate()
}

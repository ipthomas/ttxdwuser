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
			i.setTaskState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_TASKS_STATUS):
			i.setTasksState()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_EVENTS):
			i.setEvents()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_UPLOAD):
			i.setUploadTemplate()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_CLEAR_CACHE):
			InitTemplates = true
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_PATIENT):
			i.setPatient()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_CODEMAP):
			i.setUserMappings()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_PATHWAYS), strings.HasSuffix(path, HTTP_PATH_API_STATE_PATHWAYS):
			i.setPathways()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_EXPRESSIONS), strings.HasSuffix(path, HTTP_PATH_API_STATE_EXPRESSIONS):
			i.setExpressions()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_STATIC), strings.HasSuffix(path, "favicon.ico"):
			i.setImage()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_COMMENTS), strings.HasSuffix(path, HTTP_PATH_API_STATE_COMMENTS):
			i.setComments()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_XDWS):
			i.setXDWs()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_XDW):
			i.setXDW()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_SPA):
			i.setSpaTemplate()
		case strings.HasSuffix(path, HTTP_PATH_CONSUMER_MYSUBS):
			i.setUserSubscriptions()
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
			i.setState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_DASHBOARD):
			i.setDashboardState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_WORKFLOWS_COUNT):
			i.setWorkflowsCount()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_WORKFLOWS):
			i.setStates()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_WORKFLOW):
			i.setWorkflowState()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_EVENTS):
			i.setEventStates()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_TERMINOLOGY):
			i.setTerminology()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_SUBSCRIPTIONS):
			i.setSubscriptions()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_DEFINITION):
			i.setXDWDefinition()
		case strings.HasSuffix(path, HTTP_PATH_API_STATE_META):
			i.setXDWMeta()
		default:
			i.Error = errors.New("Unsupported Consumer query path - " + i.HTTP.Path)
			i.setErrorTemplate()
		}
	}
}

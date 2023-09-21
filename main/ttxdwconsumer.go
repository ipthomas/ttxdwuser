package main

import (
	"encoding/json"
	"log"
	"sort"
	"strings"
	"time"
)

func (i *Trans) getXDWDefinition() {
	xdwstr := getXDW(i.Query.Pathway+"_def", false)
	def := WorkflowDefinition{}
	json.Unmarshal([]byte(xdwstr), &def)
	var bb []byte
	bb, i.Error = json.MarshalIndent(def, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getXDWMeta() {
	metastr := getXDW(i.Query.Pathway+"_meta", true)
	meta := WorkflowMeta{}
	json.Unmarshal([]byte(metastr), &meta)
	var bb []byte
	bb, i.Error = json.MarshalIndent(meta, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getTerminology() {
	i.IdMaps = getUserTerminology(i.Query.User)
	var bb []byte
	bb, i.Error = json.MarshalIndent(i.IdMaps, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getSubscriptions() {
	i.Subscriptions = i.newSubscriptionRequest().getSubscriptions()
	var bb []byte
	bb, i.Error = json.MarshalIndent(i.Subscriptions, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getCommentsState() {
	bodybytes, err := GetTaskNotes(i.Query.Pathway, i.Query.Nhs, GetIntFromString(i.Query.Taskid), GetIntFromString(i.Query.Vers))
	i.HTTP.ResponseBody = string(bodybytes)
	i.HTTP.RspContentType = APPLICATION_JSON
	i.Error = err
}
func (i *Trans) getExpressionsState() {
	if i.Query.Pathway == "" {
		return
	}
	i.Query.Name = i.Query.Pathway
	i.setWorkflowDefinition()
	if i.Error == nil && i.XDWState.Definition.Name != "" {
		uniqueMap := make(map[string]bool)
		uniqueValues := []string{}
		for _, task := range i.XDWState.Definition.Tasks {
			for _, inp := range task.Input {
				if _, found := uniqueMap[inp.Name]; !found {
					uniqueMap[inp.Name] = true
					uniqueValues = append(uniqueValues, inp.Name)
					log.Printf("set expression key: %s", inp.Name)
				}
			}
			for _, out := range task.Output {
				if _, found := uniqueMap[out.Name]; !found {
					uniqueMap[out.Name] = true
					uniqueValues = append(uniqueValues, out.Name)
					log.Printf("set expression key: %s", out.Name)
				}
			}
		}
		wfexpressions := Expressions{}
		for _, val := range uniqueValues {
			exp := Expression{Text: GetMappedValue(i.Query.User, val), Value: val}
			wfexpressions.Expression = append(wfexpressions.Expression, exp)
		}
		bodybytes, err := json.MarshalIndent(wfexpressions.Expression, "", "  ")
		i.HTTP.ResponseBody = string(bodybytes)
		i.HTTP.RspContentType = APPLICATION_JSON
		i.Error = err
	}
}
func (i *Trans) getDashboardState() {
	i.setWorkflowStates()
	var bb []byte
	bb, i.Error = json.MarshalIndent(i.XDWState.Dashboard, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getWorkflowState() {
	vers := 1
	if i.Query.Vers != "" {
		vers = GetIntFromString(i.Query.Vers)
		if vers < 1 {
			vers = 1
		}
	}
	if i.Query.Pathway != "" && i.Query.Nhs != "" {
		i.XDWState.Workflows, i.Error = GetWorkflows(i.Query.Pathway, i.Query.Nhs, vers, "")
		if i.XDWState.Workflows.Count == 1 {
			var bb []byte
			wfdoc := WorkflowDocument{}
			wfdocstr := i.XDWState.Workflows.Workflows[1].XDW_Doc
			json.Unmarshal([]byte(wfdocstr), &wfdoc)
			bb, i.Error = json.MarshalIndent(wfdoc, "", "  ")
			i.HTTP.ResponseBody = string(bb)
			i.HTTP.RspContentType = APPLICATION_JSON
		}
	}
}
func (i *Trans) getWorkflowStates() {
	i.setWorkflowStates()
	var bb []byte
	bb, i.Error = json.MarshalIndent(i.XDWState.WorkflowStates, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getEventStates() {
	taskid := -1
	vers := -1
	if i.Query.Taskid != "" {
		taskid = GetIntFromString(i.Query.Taskid)
		if taskid < 1 {
			taskid = -1
		}
	}
	if i.Query.Vers != "" {
		vers = GetIntFromString(i.Query.Vers)
		if vers < 1 {
			vers = -1
		}
	}
	i.XDWState.Events, i.Error = getEvents(i.Query.Pathway, i.Query.Nhs, vers, taskid)
	var bb []byte
	bb, i.Error = json.MarshalIndent(i.XDWState.Events, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) getState() {
	i.setWorkflowStates()
	i.Subscriptions = i.newSubscriptionRequest().getSubscriptions()
	i.IdMaps = getUserTerminology(i.Query.User)
	type a struct {
		Dashboard Dashboard       `json:"dashboard"`
		Workflows []Workflowstate `json:"workflows"`
		Events    []Event         `json:"events"`
	}

	rsp := a{Dashboard: i.XDWState.Dashboard, Workflows: i.XDWState.WorkflowStates}

	for _, e := range i.XDWState.Events.Events {
		if e.Id > 0 {
			rsp.Events = append(rsp.Events, e)
		}
	}
	var bb []byte
	bb, i.Error = json.MarshalIndent(rsp, "", "  ")
	i.HTTP.ResponseBody = string(bb)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func (i *Trans) setXDW() {
	var bb []byte
	i.setWorkflowStates()
	if i.Error != nil {
		log.Println(i.Error.Error())
	}
	switch i.HTTP.RspContentType {
	case APPLICATION_JSON:
		bb, i.Error = json.MarshalIndent(i.XDWState, "", "  ")
		i.HTTP.ResponseBody = string(bb)
	default:
		i.Query.Template = "xdw2t_tmplt"
		i.setResponseFromTemplate()
	}
}
func (i *Trans) setXDWs() {
	var bb []byte
	i.setWorkflowStates()
	switch i.HTTP.RspContentType {
	case APPLICATION_JSON:
		bb, i.Error = json.MarshalIndent(i.XDWState, "", "  ")
		i.HTTP.ResponseBody = string(bb)
	default:
		i.Query.Template = "wfstate2t_tmplt"
		i.setResponseFromTemplate()
	}
}
func (i *Trans) setWorkflowStates() {
	if i.XDWState.Workflows.Count == 0 {
		taskid := -1
		vers := -1
		if i.Query.Taskid != "" {
			taskid = GetIntFromString(i.Query.Taskid)
			if taskid < 1 {
				taskid = -1
			}
		}
		if i.Query.Vers != "" {
			vers = GetIntFromString(i.Query.Vers)
			if vers < 1 {
				vers = -1
			}
		}
		if i.XDWState.Workflows, i.Error = GetWorkflows(i.Query.Pathway, i.Query.Nhs, vers, i.Query.Status); i.Error == nil {
			if i.XDWState.Events, i.Error = getEvents(i.Query.Pathway, i.Query.Nhs, vers, taskid); i.Error != nil {
				i.setError()
				return
			}
		}
	}
	log.Println("Setting Workflow States")
	i.XDWState.Dashboard.Total = i.XDWState.Workflows.Count
	for _, wf := range i.XDWState.Workflows.Workflows {
		if len(wf.XDW_Doc) > 0 {

			if i.Error = json.Unmarshal([]byte(wf.XDW_Doc), &i.XDWState.WorkflowDocument); i.Error != nil {
				i.setError()
				return
			}
			if i.Error = json.Unmarshal([]byte(wf.XDW_Def), &i.XDWState.Definition); i.Error != nil {
				i.setError()
				return
			}
			log.Printf("Setting %s Workflow state for Patient %s", i.XDWState.WorkflowDocument.WorkflowDefinitionReference, i.XDWState.WorkflowDocument.Patient.Extension)
			state := Workflowstate{}
			state.Created = wf.Created
			state.Status = wf.Status
			state.Published = wf.Published
			state.WorkflowId = wf.Id
			state.Pathway = wf.Pathway
			state.NHSId = wf.NHSId
			state.Version = wf.Version
			state.CreatedBy = i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Family + " " + i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Prefix
			state.LastUpdate = i.getLatestWorkflowEventTime().String()
			state.Owner = i.getWorkflowOwner()
			state.Overdue = "FALSE"
			state.Escalated = "FALSE"
			state.TargetMet = "TRUE"
			state.InProgress = "TRUE"
			state.Duration = i.getWorkflowDuration()
			if state.Status == STATUS_CLOSED {
				state.TimeRemaining = "0"
				i.XDWState.ClosedWorkflows.Count = i.XDWState.ClosedWorkflows.Count + 1
				i.XDWState.ClosedWorkflows.Workflows = append(i.XDWState.ClosedWorkflows.Workflows, wf)
			} else {
				state.TimeRemaining = i.getWorkflowTimeRemaining()
				i.XDWState.OpenWorkflows.Count = i.XDWState.OpenWorkflows.Count + 1
				i.XDWState.OpenWorkflows.Workflows = append(i.XDWState.OpenWorkflows.Workflows, wf)
			}
			workflowStartTime := GetTimeFromString(state.Created)
			if i.isWorkflowOverdue() {
				i.XDWState.Dashboard.TargetMissed = i.XDWState.Dashboard.TargetMissed + 1
				state.Overdue = "TRUE"
				state.TargetMet = "FALSE"
				i.XDWState.OverdueWorkflows.Count = i.XDWState.OverdueWorkflows.Count + 1
				i.XDWState.OverdueWorkflows.Workflows = append(i.XDWState.OverdueWorkflows.Workflows, wf)
			} else {
				if i.XDWState.WorkflowDocument.WorkflowStatus == STATUS_CLOSED {
					i.XDWState.Dashboard.TargetMet = i.XDWState.Dashboard.TargetMet + 1
					i.XDWState.MetWorkflows.Count = i.XDWState.MetWorkflows.Count + 1
					i.XDWState.MetWorkflows.Workflows = append(i.XDWState.MetWorkflows.Workflows, wf)
				}
			}
			if i.XDWState.Definition.CompleteByTime == "" {
				state.CompleteBy = "Non Specified"
			} else {
				period := strings.Split(i.XDWState.Definition.CompleteByTime, "(")[0]
				periodDuration := GetIntFromString(strings.Split(strings.Split(i.XDWState.Definition.CompleteByTime, "(")[1], ")")[0])
				switch period {
				case "month":
					state.CompleteBy = strings.Split(GetFutureDate(workflowStartTime, 0, periodDuration, 0, 0, 0, 0).String(), " +")[0]
				case "day":
					state.CompleteBy = strings.Split(GetFutureDate(workflowStartTime, 0, 0, periodDuration, 0, 0, 0).String(), " +")[0]
				case "hour":
					state.CompleteBy = strings.Split(GetFutureDate(workflowStartTime, 0, 0, 0, periodDuration, 0, 0).String(), " +")[0]
				case "min":
					state.CompleteBy = strings.Split(GetFutureDate(workflowStartTime, 0, 0, 0, 0, periodDuration, 0).String(), " +")[0]
				case "sec":
					state.CompleteBy = strings.Split(GetFutureDate(workflowStartTime, 0, 0, 0, 0, 0, periodDuration).String(), " +")[0]
				}
			}

			if i.XDWState.WorkflowDocument.WorkflowStatus == STATUS_OPEN {
				log.Printf("Workflow %s is OPEN", wf.XDW_Key)
				i.XDWState.Dashboard.InProgress = i.XDWState.Dashboard.InProgress + 1
				if i.isWorkflowEscalated() {
					log.Printf("Workflow %s is ESCALATED", wf.XDW_Key)
					i.XDWState.Dashboard.Escalated = i.XDWState.Dashboard.Escalated + 1
					state.Escalated = "TRUE"
					i.XDWState.EscalatedWorkflows.Count = i.XDWState.EscalatedWorkflows.Count + 1
					i.XDWState.EscalatedWorkflows.Workflows = append(i.XDWState.EscalatedWorkflows.Workflows, wf)
				}
			} else {
				log.Printf("Workflow %s is CLOSED", wf.XDW_Key)
				i.XDWState.Dashboard.Complete = i.XDWState.Dashboard.Complete + 1
				state.InProgress = "FALSE"
			}
			i.XDWState.WorkflowStates = append(i.XDWState.WorkflowStates, state)
		}
	}
}
func (i *Trans) getWorkflowOwner() string {
	owner := ""
	for _, ev := range i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent {
		if ev.Author != "" {
			owner = ev.Author
		}
	}
	return owner
}
func (i *Trans) getLatestWorkflowEventTime() time.Time {
	var we = GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value)
	for _, docevent := range i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent {
		if docevent.EventTime != "" {
			doceventtime := GetTimeFromString(docevent.EventTime)
			if doceventtime.After(we) {
				we = doceventtime
			}
		}
	}
	return we
}
func (i *Trans) getWorkflowDuration() string {
	ws := GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value)
	log.Printf("Workflow Started %s Status %s", ws.String(), i.XDWState.WorkflowDocument.WorkflowStatus)
	loc, _ := time.LoadLocation("Europe/London")
	we := time.Now().In(loc)
	log.Printf("Time Now %s", we.String())
	if i.XDWState.WorkflowDocument.WorkflowStatus == STATUS_CLOSED {
		we = i.getLatestWorkflowEventTime()
		log.Printf("Workflow is Complete. Latest Event Time was %s", we.String())
	}
	duration := GetDuration(we.Sub(ws))
	log.Println("Duration - " + duration)
	return duration
}
func (i *Trans) getWorkflowTimeRemaining() string {
	createTime := GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value)
	completeby := OHT_FutureDate(createTime, i.XDWState.Definition.CompleteByTime)
	log.Printf("Completion time %s", completeby.String())
	if time.Now().After(completeby) {
		return "0"
	}
	timeRemaining := time.Until(completeby)
	log.Println("Workflow Time Remaining : " + timeRemaining.String())
	return PrettyPrintDuration(timeRemaining)
}
func (i *Trans) isWorkflowOverdue() bool {
	log.Println("Checking if Workflow is Overdue")
	if i.XDWState.Definition.CompleteByTime != "" {
		completebyDate := i.getWorkflowCompleteByDate()
		log.Printf("Workflow Complete By Date is %s", completebyDate.String())
		if time.Now().After(completebyDate) {
			if i.XDWState.WorkflowDocument.WorkflowStatus == STATUS_CLOSED {
				levent := i.getLatestWorkflowEventTime()
				log.Printf("Workflow Latest Event Time %s. Workflow Target Met = %v", levent.String(), levent.Before(completebyDate))
				return levent.After(completebyDate)
			} else {
				log.Println("Workflow Target Met = false")
				return true
			}
		} else {
			log.Println("Workflow is not overdue")
			return false
		}
	}
	log.Printf("XDW Definition for %s Workflow does not specify a Complete By Time. Workflow Target Met = true", i.Query.Pathway)
	return false
}
func (i *Trans) getWorkflowCompleteByDate() time.Time {
	return OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.CompleteByTime)
}
func (i *Trans) isWorkflowEscalated() bool {
	if i.XDWState.Definition.ExpirationTime != "" {
		escalatedate := OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.ExpirationTime)
		log.Printf("Workflow Start Time %s Worklow Escalate Time %s Workflow Escaleted = %v", i.XDWState.WorkflowDocument.EffectiveTime.Value, escalatedate.String(), time.Now().After(escalatedate))
		return time.Now().After(escalatedate)
	}
	log.Println("No Escalate time defined for Workflow")
	return false
}
func (i *Trans) getPathways() {
	pwys := Pwys{}
	pathways := make(map[string]bool)
	xdws := getXDWs()
	for _, xdw := range xdws.XDW {
		if strings.HasSuffix(xdw.Name, "_def") {
			if ok := pathways[strings.TrimSuffix(xdw.Name, "_def")]; !ok {
				pathways[strings.TrimSuffix(xdw.Name, "_def")] = true
				p := Pwy{Text: strings.TrimSpace(GetMappedValue(i.Query.User, strings.TrimSuffix(xdw.Name, "_def"))), Value: strings.TrimSuffix(xdw.Name, "_def")}
				pwys.Pwy = append(pwys.Pwy, p)
			}
		}
	}
	sort.Sort(pwys)
	bodyBytes, _ := json.MarshalIndent(pwys.Pwy, "", "  ")
	i.HTTP.ResponseBody = string(bodyBytes)
	i.HTTP.RspContentType = APPLICATION_JSON
}
func GetWorkflows(pathway string, nhsid string, version int, status string) (Workflows, error) {
	wfs := Workflows{Action: SELECT}
	wf := Workflow{Pathway: pathway, NHSId: nhsid, Version: version, Status: status}
	wfs.Workflows = append(wfs.Workflows, wf)
	err := wfs.newEvent()
	log.Printf("selected %v Workflows for Pathway %s NHS %s Version %v", wfs.Count, pathway, nhsid, version)
	return wfs, err
}
func (i *Trans) getTasksState() {
}
func (i *Trans) getTaskState() {
	log.Printf("Setting Task %s Task Index %v State", i.Query.Taskid, GetIntFromString(i.Query.Taskid))
	i.Query.Name = i.Query.Pathway
	i.setWorkflowDocument()
	if i.Error == nil {
		i.setWorkflowDefinition()
		if i.Error == nil {
			taskstate := TaskState{Taskid: GetIntFromString(i.Query.Taskid)}
			if i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].CompleteByTime == "" {
				taskstate.CompleteBy = OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.CompleteByTime).String()
			} else {
				taskstate.CompleteBy = OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].CompleteByTime).String()
			}
			if i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].StartByTime == "" {
				taskstate.StartBy = OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.CompleteByTime).String()
			} else {
				taskstate.StartBy = OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].StartByTime).String()
			}
			if i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].ExpirationTime == "" {
				taskstate.EscalateOn = OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.CompleteByTime).String()
			} else {
				taskstate.EscalateOn = OHT_FutureDate(GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value), i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].ExpirationTime).String()
			}
			for _, v := range i.XDWState.Definition.Tasks[GetIntFromString(i.Query.Taskid)-1].CompletionBehavior {
				if v.Completion.Condition != "" {
					taskstate.CompletionConditions = append(taskstate.CompletionConditions, v.Completion.Condition)
				}
			}
			taskstate.Status = i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.Status
			if taskstate.Status == STATUS_COMPLETE {
				taskstate.CompletedOn = i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.LastModifiedTime
			}
			taskstate.StartedOn = i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.ActivationTime
			taskstate.Owner = i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.ActualOwner
			taskstate.Duration = i.getTaskDuration()
			if time.Now().After(GetTimeFromString(taskstate.CompleteBy)) {
				taskstate.TargetMet = i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.Status == STATUS_COMPLETE
			} else {
				taskstate.TargetMet = true
			}
			if time.Now().After(GetTimeFromString(taskstate.EscalateOn)) {
				taskstate.Escalated = i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.Status != STATUS_COMPLETE
			} else {
				taskstate.Escalated = false
			}
			var bb []byte
			bb, i.Error = json.MarshalIndent(taskstate, "", "  ")
			i.HTTP.ResponseBody = string(bb)
			i.HTTP.RspContentType = APPLICATION_JSON
		}
	}
	if i.Error != nil {
		log.Println(i.Error.Error())
	}
}
func (i *Trans) getTaskDuration() string {
	at := i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.ActivationTime
	isActive := at != ""
	lmt := i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.LastModifiedTime
	if isActive {
		if i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.Status == STATUS_COMPLETE {
			return timeDuratipn(i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(i.Query.Taskid)-1].TaskData.TaskDetails.ActivationTime, lmt)
		}
		return GetDurationSince(at)
	} else {
		return "0"
	}
}

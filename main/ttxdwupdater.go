package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func (i *Trans) newXDWUpdater() {
	i.XDWState.Workflows, i.Error = GetWorkflows(i.Query.Pathway, i.Query.Nhs, GetIntFromString(i.Query.Vers), STATUS_OPEN)
	for _, wf := range i.XDWState.Workflows.Workflows {
		if wf.Id > 0 {
			if i.Error = json.Unmarshal([]byte(wf.XDW_Def), &i.XDWState.Definition); i.Error == nil {
				log.Printf("Loaded %s Definition", wf.Pathway)
				if err := json.Unmarshal([]byte(wf.XDW_Doc), &i.XDWState.WorkflowDocument); err == nil {
					log.Printf("Loaded Version %v of %s XDW Document for NHS %s", wf.Version, wf.Pathway, wf.NHSId)
					i.XDWState.Events, i.Error = getEvents(wf.Pathway, wf.NHSId, wf.Version, -1)
					log.Printf("Selected %v Events for Pathway %s NHS %s Version %v", i.XDWState.Events.Count, wf.Pathway, wf.NHSId, wf.Version)
					i.Query.Vers = GetStringFromInt(wf.Version)
					i.updateXDW()
				}
			}
		}
	}
}
func (i *Trans) updateXDW() {
	log.Println("Updating Workflow State")
	newSequenceNumber := false
	i.Query.Topic = XDW_OPERATION_WORKFLOW
	for _, ev := range i.XDWState.Events.Events {
		if !strings.EqualFold(ev.Topic, WORKFLOW) {
			for k, wfdoctask := range i.XDWState.WorkflowDocument.TaskList.XDWTask {
				for inp, input := range wfdoctask.TaskData.Input {
					if ev.Expression == input.Part.Name {
						log.Printf("Matched Task %v %s Input with event %v", wfdoctask.TaskData.TaskDetails.ID, input.Part.Name, ev.Id)
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Input[inp].Part.AttachmentInfo.AttachedTime = ev.Creationtime
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Input[inp].Part.AttachmentInfo.AttachedBy = ev.User + " " + ev.Org + " " + ev.Role
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Input[inp].Part.AttachmentInfo.HomeCommunityId = i.EnvVars.REG_OID
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.LastModifiedTime = ev.Creationtime
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.Status = IN_PROGRESS
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActualOwner = ev.User + " " + ev.Org + " " + ev.Role
						if i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActivationTime == "" {
							i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActivationTime = ev.Creationtime
							log.Printf("Set Task %s Activation Time %s", wfdoctask.TaskData.TaskDetails.ID, i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActivationTime)
						}
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Input[inp].Part.AttachmentInfo.Identifier = GetStringFromInt(int(ev.Id))
						if i.isNewEvent(ev) {
							i.newXDWTaskEvent(ev)
							i.newXDWDocEvent(ev)
							newSequenceNumber = true
						}
					}
				}
				for oup, output := range i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Output {
					if ev.Expression == output.Part.Name {
						log.Println("Matched workflow document task " + wfdoctask.TaskData.TaskDetails.ID + " Output Part : " + output.Part.Name + " with Event Expression : " + ev.Expression + " Status : " + wfdoctask.TaskData.TaskDetails.Status)
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.LastModifiedTime = ev.Creationtime
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Output[oup].Part.AttachmentInfo.AttachedTime = ev.Creationtime
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Output[oup].Part.AttachmentInfo.AttachedBy = ev.User + " " + ev.Org + " " + ev.Role
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActualOwner = ev.User + " " + ev.Org + " " + ev.Role
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.Status = IN_PROGRESS
						if i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActivationTime == "" {
							i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActivationTime = ev.Creationtime
							log.Printf("Set Task %s Activation Time %s", wfdoctask.TaskData.TaskDetails.ID, i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.ActivationTime)
						}
						i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.Output[oup].Part.AttachmentInfo.Identifier = GetStringFromInt(int(ev.Id))
						if i.isNewEvent(ev) {
							i.newXDWTaskEvent(ev)
							i.newXDWDocEvent(ev)
							newSequenceNumber = true
						}
					}
				}
			}
		}
	}
	if newSequenceNumber {
		wfseqnum, _ := strconv.ParseInt(i.XDWState.WorkflowDocument.WorkflowDocumentSequenceNumber, 0, 0)
		wfseqnum = wfseqnum + 1
		i.XDWState.WorkflowDocument.WorkflowDocumentSequenceNumber = strconv.Itoa(int(wfseqnum))
	}
	for cnt := 0; cnt < 3; cnt++ {
		i.checkTasksCompletionBehaviour()
	}

	if complete, completiontask := i.isWorkflowCompleteBehaviorMet(); complete {
		i.XDWState.WorkflowDocument.WorkflowStatus = STATUS_CLOSED
		docevent := DocumentEvent{}
		docevent.Author = i.Query.User + " " + i.Query.Org + " " + i.Query.Role
		docevent.TaskEventIdentifier = completiontask
		docevent.EventTime = Time_Now()
		docevent.EventType = XDW_TASKEVENTTYPE_WORKFLOW_COMPLETED
		docevent.PreviousStatus = i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent[len(i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent)-1].ActualStatus
		docevent.ActualStatus = STATUS_CLOSED
		i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent = append(i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent, docevent)
		for k := range i.XDWState.WorkflowDocument.TaskList.XDWTask {
			i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.Status = STATUS_COMPLETE
		}
		i.Query.Eventtype = XDW_OPERATION_CLOSE_WORKFLOW
		i.Query.Expression = strings.ToUpper(i.XDWState.WorkflowDocument.WorkflowDefinitionReference)
		i.Query.Taskid = completiontask
		i.newXDWEvent()
		log.Println("Closed Workflow")
	}
	i.persistUpdatedXDW()
}
func (i *Trans) checkTasksCompletionBehaviour() {
	for k, task := range i.XDWState.WorkflowDocument.TaskList.XDWTask {
		if i.isTaskCompleteBehaviorMet(task.TaskData.TaskDetails.ID) {
			i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.Status = STATUS_COMPLETE
		} else {
			if i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.Status == STATUS_COMPLETE {
				i.XDWState.WorkflowDocument.TaskList.XDWTask[k].TaskData.TaskDetails.Status = IN_PROGRESS
			}
		}
	}
}
func (i *Trans) isNewEvent(ev Event) bool {
	for _, te := range i.XDWState.WorkflowDocument.TaskList.XDWTask[ev.Taskid-1].TaskEventHistory.TaskEvent {
		if GetIntFromString(te.ID) == ev.Id {
			log.Printf("Event ID %v is Registered", ev.Id)
			return false
		}
	}
	log.Printf("New Event Found. ID = %v", ev.Id)
	return true
}
func (i *Trans) newXDWTaskEvent(ev Event) {
	nte := TaskEvent{
		ID:         GetStringFromInt(ev.Id),
		Identifier: GetStringFromInt(ev.Taskid),
		EventTime:  Pretty_Time_Now(),
		EventType:  i.XDWState.WorkflowDocument.TaskList.XDWTask[ev.Taskid-1].TaskData.TaskDetails.TaskType,
		Status:     STATUS_COMPLETE,
	}
	i.XDWState.WorkflowDocument.TaskList.XDWTask[ev.Taskid-1].TaskEventHistory.TaskEvent = append(i.XDWState.WorkflowDocument.TaskList.XDWTask[ev.Taskid-1].TaskEventHistory.TaskEvent, nte)
}
func (i *Trans) isTaskCompleteBehaviorMet(taskid string) bool {
	for _, task := range i.XDWState.Definition.Tasks {
		if task.ID == taskid {
			log.Printf("Checking if Task %s is complete", taskid)
			var conditions []string
			var completedConditions = 0
			for _, cond := range task.CompletionBehavior {
				log.Printf("Task %s Completion Condition is %s", taskid, cond)
				if cond.Completion.Condition != "" {
					if strings.Contains(cond.Completion.Condition, " and ") {
						conditions = strings.Split(cond.Completion.Condition, " and ")
					} else {
						conditions = append(conditions, cond.Completion.Condition)
					}
					log.Printf("Checkiing Task %s %v completion conditions", taskid, len(conditions))

					for _, condition := range conditions {
						endMethodInd := strings.Index(condition, "(")
						if endMethodInd > 0 {
							method := condition[0:endMethodInd]
							endParamInd := strings.Index(condition, ")")
							if endParamInd < endMethodInd+2 {
								log.Println("Invalid Condition. End bracket index invalid")
								continue
							}
							param := condition[endMethodInd+1 : endParamInd]
							log.Printf("Completion condition is %s", method)

							switch method {
							case "output":
								for _, op := range i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(taskid)-1].TaskData.Output {
									if op.Part.AttachmentInfo.AttachedTime != "" {
										if op.Part.AttachmentInfo.Name == param {
											completedConditions = completedConditions + 1
											log.Printf("Task %s Output Part %s - Attached %s", taskid, op.Part.AttachmentInfo.Name, op.Part.AttachmentInfo.AttachedTime)
										}
									}
								}
							case "input":
								for _, in := range i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(taskid)-1].TaskData.Input {
									if in.Part.AttachmentInfo.AttachedTime != "" && in.Part.AttachmentInfo.Name == param {
										completedConditions = completedConditions + 1
									}
								}
							case "task":
								if i.XDWState.WorkflowDocument.TaskList.XDWTask[GetIntFromString(param)-1].TaskData.TaskDetails.Status == STATUS_COMPLETE {
									completedConditions = completedConditions + 1
								}
							case "latest":
								if i.getLatestXDWTaskEvent(GetIntFromString(taskid)-1) == param {
									completedConditions = completedConditions + 1
								}
							}
						}
					}
				}
			}
			if len(conditions) == completedConditions {
				log.Printf("Task %s is complete", taskid)
				return true
			}
		}
	}
	log.Printf("Task %s is not complete", taskid)
	return false
}
func (i *Trans) isWorkflowCompleteBehaviorMet() (bool, string) {
	var completionTask string
	for c, cc := range i.XDWState.Definition.CompletionBehavior {
		var conditions []string
		var completedConditions = 0
		if cc.Completion.Condition != "" {
			log.Printf("Workflow Completion Condition %v %s", c+1, cc.Completion.Condition)
			if strings.Contains(cc.Completion.Condition, " and ") {
				conditions = strings.Split(cc.Completion.Condition, " and ")
			} else {
				conditions = append(conditions, cc.Completion.Condition)
			}
			for _, condition := range conditions {
				endMethodInd := strings.Index(condition, "(")
				if endMethodInd > 0 {
					method := cc.Completion.Condition[0:endMethodInd]
					if method != "task" {
						log.Println(method + " is an Invalid Workflow Completion Behaviour Condition method. Ignoring Condition")
						continue
					}
					endParamInd := strings.Index(cc.Completion.Condition, ")")
					param := cc.Completion.Condition[endMethodInd+1 : endParamInd]
					for _, task := range i.XDWState.WorkflowDocument.TaskList.XDWTask {
						if task.TaskData.TaskDetails.ID == param {
							if task.TaskData.TaskDetails.Status == STATUS_COMPLETE {
								completedConditions = completedConditions + 1
								completionTask = param
							}
						}
					}
				}
			}
			if len(conditions) == completedConditions {
				log.Printf("%s Workflow for Nhs ID %s is complete", i.Query.Pathway, i.Query.Nhs)
				return true, completionTask
			}
		}
	}
	log.Printf("%s Workflow for Nhs ID %s is not complete", i.Query.Pathway, i.Query.Nhs)
	return false, ""
}
func (i *Trans) newXDWDocEvent(ev Event) {
	docevent := DocumentEvent{}
	docevent.Author = ev.User + " " + ev.Org + " " + ev.Role
	docevent.TaskEventIdentifier = GetStringFromInt(ev.Taskid)
	docevent.EventTime = ev.Creationtime
	docevent.EventType = i.XDWState.WorkflowDocument.TaskList.XDWTask[ev.Taskid-1].TaskData.TaskDetails.TaskType
	docevent.PreviousStatus = i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent[len(i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent)-1].ActualStatus
	docevent.ActualStatus = IN_PROGRESS
	i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent = append(i.XDWState.WorkflowDocument.WorkflowStatusHistory.DocumentEvent, docevent)
}
func (i *Trans) persistUpdatedXDW() {
	wfs := Workflows{Action: UPDATE}
	wf := Workflow{
		Pathway: i.Query.Pathway,
		NHSId:   i.Query.Nhs,
		XDW_Key: i.Query.Pathway + i.Query.Nhs,
		XDW_UID: i.XDWState.WorkflowDocument.ID.Extension,
		Version: GetIntFromString(i.Query.Vers),
		Status:  i.XDWState.WorkflowDocument.WorkflowStatus,
	}

	xdwDocBytes, _ := json.MarshalIndent(i.XDWState.WorkflowDocument, "", "  ")
	wf.XDW_Doc = string(xdwDocBytes)
	wfs.Workflows = append(wfs.Workflows, wf)
	if i.Error = NewDBEvent(&wfs); i.Error != nil {
		return
	}
	log.Printf("Upddated Workflow State for Pathway %s Nhs ID %s Version %v Status %s", i.Query.Pathway, i.Query.Nhs, i.Query.Vers, i.XDWState.WorkflowDocument.WorkflowStatus)
}
func (i *Trans) getLatestXDWTaskEvent(taskid int) string {
	var lasteventtime = GetTimeFromString(i.XDWState.WorkflowDocument.EffectiveTime.Value)
	var lastevent = ""
	for _, v := range i.XDWState.WorkflowDocument.TaskList.XDWTask[taskid].TaskData.Input {
		if v.Part.AttachmentInfo.AttachedTime != "" {
			et := GetTimeFromString(v.Part.AttachmentInfo.AttachedTime)
			if et.After(lasteventtime) {
				lasteventtime = et
				lastevent = v.Part.AttachmentInfo.Name
			}
		}
	}
	for _, v := range i.XDWState.WorkflowDocument.TaskList.XDWTask[taskid].TaskData.Output {
		if v.Part.AttachmentInfo.AttachedTime != "" {
			et := GetTimeFromString(v.Part.AttachmentInfo.AttachedTime)
			if et.After(lasteventtime) {
				lasteventtime = et
				lastevent = v.Part.AttachmentInfo.Name
			}
		}
	}
	return lastevent
}

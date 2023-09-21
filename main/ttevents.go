package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"strings"
	"time"
)

func (i *Trans) newXDWEvent() int {
	i.Query.Name = i.Query.Pathway
	if i.XDWState.Meta.Id != i.Query.Pathway+"_meta" {
		i.setWorkflowXDSMeta()
	}
	if i.Error == nil {
		ev := Event{
			Eventtype:      i.Query.Eventtype,
			Docname:        GetMappedValue(i.Query.User, i.Query.Pathway),
			Classcode:      i.XDWState.Meta.ClasscodeValue,
			Confcode:       i.XDWState.Meta.ConfcodeValue,
			Formatcode:     i.XDWState.Meta.FormatcodeValue,
			Facilitycode:   i.XDWState.Meta.FacilitycodeValue,
			Practicecode:   i.XDWState.Meta.PracticesettingcodeValue,
			Repositoryuid:  i.XDWState.Meta.Repositoryuniqueid,
			Expression:     i.Query.Expression,
			Authors:        i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Prefix + " " + i.XDWState.WorkflowDocument.Author.AssignedAuthor.AssignedPerson.Name.Family,
			Xdsdocentryuid: i.XDWState.WorkflowDocument.ID.Extension,
			Nhs:            i.Query.Nhs,
			User:           i.Query.User,
			Org:            i.Query.Org,
			Role:           i.Query.Role,
			Topic:          i.Query.Topic,
			Pathway:        i.Query.Pathway,
			Comments:       i.Query.Comments,
			Version:        GetIntFromString(i.Query.Vers),
			Taskid:         GetIntFromString(i.Query.Taskid),
		}
		evs := Events{Action: INSERT}
		evs.Events = append(evs.Events, ev)
		if i.Error = NewDBEvent(&evs); i.Error != nil {
			log.Println(i.Error.Error())
			return 0
		}
		log.Printf("Created Event ID = %v Pathway = %s Type = %s", evs.LastInsertId, ev.Pathway, ev.Eventtype)
		if i.EnvVars.DEBUG_MODE {
			logStruct(ev)
		}
		return evs.LastInsertId
	}
	return 0
}
func (i *Trans) newUserEvent() {
	log.Println("Processing New User Event")
	i.Query.Name = i.Query.Pathway
	i.setWorkflowXDSMeta()
	if i.Error != nil {
		return
	}
	event := Event{
		Eventtype:      i.Query.Eventtype,
		Docname:        i.XDWState.Meta.Docname,
		Classcode:      i.XDWState.Meta.ClasscodeValue,
		Confcode:       i.XDWState.Meta.ConfcodeValue,
		Formatcode:     i.Query.Formatcode,
		Facilitycode:   GetMappedValue(i.Query.User, i.Query.Org),
		Practicecode:   GetMappedValue(i.Query.User, i.Query.Role),
		Expression:     i.Query.Expression,
		Authors:        i.Query.User,
		Nhs:            i.Query.Nhs,
		User:           i.Query.User,
		Org:            i.Query.Org,
		Role:           i.Query.Role,
		Speciality:     i.Query.Role,
		Topic:          i.Query.Topic,
		Pathway:        i.Query.Pathway,
		Comments:       i.Query.Comments,
		Version:        GetIntFromString(i.Query.Vers),
		Taskid:         GetIntFromString(i.Query.Taskid),
		Repositoryuid:  i.Query.Repositoryuid,
		Xdsdocentryuid: i.Query.Xdsdocentryuid,
		Brokerref:      i.Query.Brokerref,
	}
	if event.Repositoryuid == "" {
		event.Repositoryuid = i.XDWState.Meta.Repositoryuniqueid
	}
	if i.Error = event.persistEvent(); i.Error == nil {
		log.Println("Persisted Event")
		logStruct(event)
		i.newXDWUpdater()
		event.Creationtime = time.Now().Local().String()
		i.XDWState.Events = Events{}
		i.XDWState.Events.Events = append(i.XDWState.Events.Events, event)
		i.notifyEmailSubscribers()
	} else {
		log.Println(i.Error.Error())
	}

	i.Query.Template = "xdw2t_tmplt"
	i.setResponseFromTemplate()
}
func (i *Trans) getEvents() {
	if i.XDWState.Events, i.Error = getEvents(i.Query.Pathway, i.Query.Nhs, GetIntFromString(i.Query.Vers), GetIntFromString(i.Query.Taskid)); i.Error != nil {
		log.Println(i.Error.Error())
		return
	}
	log.Printf("Found %v Events", i.XDWState.Events.Count)
	if i.HTTP.ReqContentType == APPLICATION_JSON || i.Query.Format == APPLICATION_JSON {
		eventBytes, _ := json.MarshalIndent(i.XDWState.Events.Events, "", "  ")
		i.HTTP.ResponseBody = string(eventBytes)
		i.HTTP.RspContentType = APPLICATION_JSON
	} else {
		if i.Query.Template == "" {
			i.Query.Template = "events_tmplt"
		}
		i.setResponseFromTemplate()
	}
}
func getEvents(pathway string, nhs string, vers int, taskid int) (Events, error) {
	events := Events{Action: SELECT}
	ev := Event{Pathway: pathway, Nhs: nhs, Version: vers, Taskid: taskid}
	events.Events = append(events.Events, ev)
	err := events.newEvent()
	return events, err
}
func (i *Trans) deprecateEvents() {
	events := Events{Action: DEPRECATE}
	ev := Event{Pathway: i.Query.Pathway, Nhs: i.Query.Nhs}
	events.Events = append(events.Events, ev)
	i.Error = events.newEvent()
	log.Printf("Deprecated %v Events", i.XDWState.Events.LastInsertId)
}

// creates a DSUBEventMessage and populates a new Event with the DSUBEventMessage values
// It then checks for Subscriptions matching the brokerref and creates an Event for each subscription
// A DSUB ack response is always returned regardless of success
// A DSUB Cancel response is sent if no brokerref subscriptions are found
func (i *Trans) newDSUBEvent() {
	i.setDsubAckTemplate()
	notifyMsg := DSUBNotifyMessage{}
	if len(i.HTTP.RequestBody) > 1 {
		if i.Error = xml.Unmarshal([]byte(i.HTTP.RequestBody), &notifyMsg); i.Error == nil {
			i.Query.Brokerref = notifyMsg.NotificationMessage.SubscriptionReference.Address.Text
			if i.Query.Brokerref != "" {
				event := i.initDSUBEvent(notifyMsg)
				if i.Query.Pid != "" {
					i.setNHS()
					if i.Query.Nhs != "" {
						event.Nhs = i.Query.Nhs
						i.setBrokerSubscriptions()
						if i.Error == nil {
							for _, dbsub := range i.Subscriptions.Subscriptions {
								if dbsub.Id > 0 && i.Error == nil {
									log.Printf("Creating Event for Dsub %s Notification", dbsub.BrokerRef)
									event.Pathway = dbsub.Pathway
									event.Topic = dbsub.Topic
									event.Eventtype = XDW_TASKEVENTTYPE_ATTACHMENT
									if i.Error = event.persistEvent(); i.Error == nil {
										log.Println("Persisted Event. Updating Applicable Workflow State")
										i.newXDWUpdater()
									}
								}
							}
							event.Creationtime = time.Now().Local().String()
							i.XDWState.Events = Events{}
							i.XDWState.Events.Events = append(i.XDWState.Events.Events, event)
							i.notifyEmailSubscribers()
						}
					}
				}
			}
		}
	}
}

// InitDSUBEvent initialise the DSUBEvent struc with values parsed from the DSUBEventMessage
func (i *Trans) initDSUBEvent(notify DSUBNotifyMessage) Event {
	event := Event{}
	event.Docname = notify.NotificationMessage.Message.SubmitObjectsRequest.RegistryObjectList.ExtrinsicObject.Name.LocalizedString.Value
	event.Brokerref = notify.NotificationMessage.SubscriptionReference.Address.Text
	for _, slot := range notify.NotificationMessage.Message.SubmitObjectsRequest.RegistryObjectList.ExtrinsicObject.Slot {
		switch slot.Name {
		case REPOSITORY_UID:
			event.Repositoryuid = slot.ValueList.Value[0]
			log.Printf("Set Repository UID %s", event.Repositoryuid)
		case SOURCE_PATIENT_ID:
			i.Query.Pid = strings.Split(slot.ValueList.Value[0], "^")[0]
			i.Query.Pidoid = strings.Split(slot.ValueList.Value[0], "&")[1]
			log.Printf("Set XDS PID %s OID %s", i.Query.Pid, i.Query.Pidoid)
		}
	}
	for _, c := range notify.NotificationMessage.Message.SubmitObjectsRequest.RegistryObjectList.ExtrinsicObject.Classification {
		val := c.Name.LocalizedString.Value
		switch c.ClassificationScheme {
		case URN_CLASS_CODE:
			event.Classcode = val
			log.Printf("Set Class Code %s", val)
		case URN_CONF_CODE:
			event.Confcode = val
			log.Printf("Set Conf Code %s", val)
		case URN_FORMAT_CODE:
			event.Formatcode = val
			log.Printf("Set Format Code %s", val)
		case URN_FACILITY_CODE:
			event.Facilitycode = val
			log.Printf("Set Facility Code %s", val)
		case URN_PRACTICE_CODE:
			event.Practicecode = val
			log.Printf("Set Practice Code %s", val)
		case URN_TYPE_CODE:
			event.Expression = val
			log.Printf("Set Type Code %s", val)
		case URN_AUTHOR:
			for _, slot := range c.Slot {
				switch slot.Name {
				case AUTHOR_PERSON:
					for _, slotval := range slot.ValueList.Value {
						event.User = event.User + strings.TrimSpace(strings.ReplaceAll(slotval, "^", " ")) + " "
					}
					event.User = strings.TrimSpace(event.User)
					log.Printf("Set User %s", event.User)
				case AUTHOR_INSTITUTION:
					for _, slotval := range slot.ValueList.Value {
						event.Org = event.Org + slotval + " "
					}
					event.Org = strings.TrimSpace(event.Org)
					log.Printf("Set Org %s", event.Org)
				case AUTHOR_ROLE:
					for _, slotval := range slot.ValueList.Value {
						event.Role = event.Role + slotval + " "
					}
					event.Role = strings.TrimSpace(event.Role)
					log.Printf("Set Role %s", event.Role)
				case AUTHOR_SPECIALITY:
					for _, slotval := range slot.ValueList.Value {
						event.Speciality = event.Speciality + slotval + ". "
					}
					event.Speciality = strings.TrimSpace(event.Speciality)
					log.Printf("Set Speciality %s", event.Speciality)
				}
			}
		case URN_EVENT_LIST:
			event.Comments = val + " "
			for _, slot := range c.Slot {
				for _, slotval := range slot.ValueList.Value {
					event.Comments = event.Comments + slotval + " "
				}
			}
			event.Comments = strings.TrimSpace(event.Comments)
			log.Printf("Set event codes %s", event.Comments)
		default:
			log.Printf("Unknown classication scheme %s. Skipping", c.ClassificationScheme)
		}
	}
	for _, exid := range notify.NotificationMessage.Message.SubmitObjectsRequest.RegistryObjectList.ExtrinsicObject.ExternalIdentifier {
		if exid.IdentificationScheme == URN_XDS_DOCUID {
			event.Xdsdocentryuid = exid.Value
			log.Printf("Set XDS Doc UID %s", event.Xdsdocentryuid)
		}
	}
	log.Println("Created workflow Event from DSUB Notify Message")
	return event
}
func (i Event) persistEvent() error {
	events := Events{Action: INSERT}
	events.Events = append(events.Events, i)
	return events.newEvent()
}

package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (i *Trans) setBrokerSubscriptions() {
	sub := Subscription{BrokerRef: i.Query.Brokerref}
	i.Subscriptions = sub.getSubscriptions()
	if i.Subscriptions.Count < 1 {
		log.Printf("No Subscriptions found with brokerref = %s. Sending Cancel request to Broker", i.Query.Brokerref)
		i.Error = errors.New("No Subscriptions found with brokerref = " + i.Query.Brokerref)
		i.newDsubCancelReq()
	}
}
func (i *Trans) getUserSubscriptions() {
	i.Subscriptions = i.newSubscriptionRequest().getSubscriptions()
	i.sortSubscriptions()
	if i.HTTP.RspContentType == "" || i.HTTP.RspContentType == TEXT_HTML {
		i.Query.Template = "subscriptions2t_tmplt"
		i.setResponseFromTemplate()
	} else {
		bb, _ := json.Marshal(i.Subscriptions)
		i.HTTP.ResponseBody = string(bb)
	}
}
func (i *Trans) newUserSubscription() {
	sub := i.newSubscriptionRequest()
	if ok := sub.hasUserSubscription(); !ok {
		if i.Error = sub.newSubscription(); i.Error != nil {
			log.Println(i.Error.Error())
		}
	}
	i.getUserSubscriptions()
}
func (i *Trans) cancelUserSubscription() {
	i.newSubscriptionRequest().cancelUserSubscription()
	i.Query.Id = "0"
	i.getUserSubscriptions()
}
func (i *Trans) newBrokerSubscriptions() {
	i.HTTP.Method = http.MethodPost
	i.HTTP.ReqURL = i.EnvVars.DSUB_BROKER_URL
	i.Query.User = "ICB"
	i.Query.Org = "LHSCR"
	i.Query.Role = "Broker"
	i.Query.Topic = DSUB_TOPIC_TYPE_CODE
	for _, expression := range i.XDWState.Expressions {
		i.Query.Expression = expression
		if ok, ref := i.newSubscriptionRequest().hasBrokerSubscription(); ok {
			log.Printf("Found existing Broker Subscription Reference %s", ref)
			continue
		}
		subrsp := DSUBSubscribeResponse{}
		log.Printf("Sending Subscribe Request for TypeCode %s to DSUB Broker %s", expression, i.HTTP.ReqURL)
		i.newDsubSubscribeReq()
		if i.Error = xml.Unmarshal([]byte(i.HTTP.ResponseBody), &subrsp); i.Error == nil {
			i.Query.Brokerref = subrsp.Body.SubscribeResponse.SubscriptionReference.Address
			if i.Query.Brokerref != "" {
				if i.Error = i.newSubscriptionRequest().newSubscription(); i.Error == nil {
					log.Printf("Registered Broker Subscription with Workflow Service for Pathway %s Document Type %s Availability Notifications", i.Query.Pathway, i.Query.Expression)
				}
			}
		}
		if i.Error != nil {
			log.Println(i.Error.Error())
			return
		}
	}
}
func isSubscriptionMatch(sub Subscription, event Event) bool {
	var shouldReturn bool
	if sub.Email != "" {
		shouldReturn = (sub.NhsId == "" || strings.EqualFold(event.Nhs, sub.NhsId)) && (sub.Expression == "" || strings.EqualFold(event.Expression, sub.Expression))
	}
	return shouldReturn
}
func (i *Trans) setEmailNotifySubscriptions() {
	sub := Subscription{
		Pathway: i.Query.Pathway,
		Topic:   "EMAIL",
	}
	subs := sub.getSubscriptions()
	for _, sub := range subs.Subscriptions {
		if sub.Id > 0 {
			if isSubscriptionMatch(sub, i.XDWState.Events.Events[0]) {
				log.Printf("Matched User %s %s %s Subscription for %s %s %s", sub.User, sub.Org, sub.Role, i.XDWState.Events.Events[0].Pathway, i.XDWState.Events.Events[0].Nhs, i.XDWState.Events.Events[0].Expression)
				i.Subscriptions.Subscriptions = append(i.Subscriptions.Subscriptions, sub)
			}
		}
	}
}
func (i *Trans) newSubscriptionRequest() Subscription {
	sub := Subscription{
		BrokerRef:  i.Query.Brokerref,
		Pathway:    i.Query.Pathway,
		Topic:      i.Query.Topic,
		Expression: i.Query.Expression,
		User:       i.Query.User,
		Org:        i.Query.Org,
		Role:       i.Query.Role,
		Email:      i.Query.Email,
		NhsId:      i.Query.Nhs,
		Id:         GetIntFromString(i.Query.Id),
	}
	return sub
}
func (i Subscription) getSubscriptions() Subscriptions {
	subs := Subscriptions{Action: SELECT}
	subs.Subscriptions = append(subs.Subscriptions, i)
	subs.newEvent()
	return subs
}
func (i Subscription) newSubscription() error {
	subs := Subscriptions{Action: INSERT}
	subs.Subscriptions = append(subs.Subscriptions, i)
	return subs.newEvent()
}
func (i Subscription) hasBrokerSubscription() (bool, string) {
	i.Topic = DSUB_TOPIC_TYPE_CODE
	log.Println("Retrieving Subscriptions - Query params")
	logStruct(i)
	subs := i.getSubscriptions()
	log.Printf("Retrieved %v Subscriptions", subs.Count)
	logStruct(subs)

	if subs.Count > 0 {
		for _, v := range subs.Subscriptions {
			if v.BrokerRef != "" {
				return true, v.BrokerRef
			}
		}
	}
	return false, ""
}
func (i Subscription) hasUserSubscription() bool {
	return i.getSubscriptions().Count == 1
}
func (i Subscription) cancelUserSubscription() {
	subs := Subscriptions{Action: DELETE}
	subs.Subscriptions = append(subs.Subscriptions, i)
	subs.newEvent()
}

package main

import (
	"fmt"
	"log"
	"sync"
)

func (i *Trans) notifyEmailSubscribers() {
	log.Println("Checking for Subscriptions")
	i.setEmailNotifySubscriptions()
	wg := sync.WaitGroup{}
	for _, sub := range i.Subscriptions.Subscriptions {
		i.Subscription = sub
		i.setNotificationTemplate()
		emailBody := fmt.Sprintf("Subject: %s\r\n\r\n%s", i.EnvVars.SMTP_SUBJECT, i.HTTP.ResponseBody)
		notifyEvent := i.NewEmailNotifyEvent(emailBody, sub.Email)
		wg.Add(1)
		go func(notifyEvent NotifyEvent) {
			defer wg.Done()
			if err := notifyEvent.sendEmailNotification(); err != nil {
				log.Println(err.Error())
			}
		}(notifyEvent)
	}
	wg.Wait()
}

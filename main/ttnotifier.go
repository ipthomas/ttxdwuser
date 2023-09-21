package main

import (
	"fmt"
	"log"
	"sync"
)

func (i *Trans) notifyEmailSubscribers() {
	i.setEmailNotifySubscriptions()
	wg := sync.WaitGroup{}
	for _, sub := range i.Subscriptions.Subscriptions {
		i.Subscription = sub
		i.setNotificationTemplate()
		emailBody := fmt.Sprintf("Subject: %s\r\n\r\n%s", i.EnvVars.SMTP_SUBJECT, i.HTTP.ResponseBody)
		ne := NotifyEvent{
			Body:     emailBody,
			From:     i.EnvVars.SMTP_USER,
			To:       sub.Email,
			Server:   i.EnvVars.SMTP_SERVER,
			Port:     i.EnvVars.SMTP_PORT,
			Password: i.EnvVars.SMTP_PASSWORD,
		}
		wg.Add(1)
		go func(ne NotifyEvent) {
			defer wg.Done()
			if err := ne.Notify(); err != nil {
				log.Println(err.Error())
			}
		}(ne)
	}
	wg.Wait()
}

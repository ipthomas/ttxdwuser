package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

func (i *Trans) NewEmailNotifyEvent(body string, to string) NotifyEvent {
	return NotifyEvent{
		Body:     body,
		From:     i.EnvVars.SMTP_USER,
		To:       to,
		Server:   i.EnvVars.SMTP_SERVER,
		Port:     i.EnvVars.SMTP_PORT,
		Password: i.EnvVars.SMTP_PASSWORD,
	}
}
func (i *NotifyEvent) sendEmailNotification() error {
	var err error
	auth := smtp.PlainAuth("", i.From, i.Password, i.Server)
	conn, err := smtp.Dial(i.Server + ":" + i.Port)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer conn.Close()
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if err = conn.StartTLS(tlsConfig); err != nil {
		log.Println(err.Error())
		return err
	}
	if err = conn.Auth(auth); err != nil {
		log.Println(err.Error())
		return err
	}
	if err = conn.Mail(i.From); err != nil {
		log.Println(err.Error())
		return err
	}
	if err = conn.Rcpt(i.To); err != nil {
		log.Println(err.Error())
		return err
	}
	wc, err := conn.Data()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer wc.Close()
	if _, err = fmt.Fprint(wc, i.Body); err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("Notification sent to %s", i.To)
	}

	return err
}

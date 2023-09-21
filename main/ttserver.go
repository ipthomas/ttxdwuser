package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer() {
	logfile := createLog("logs")
	var i = Trans{EnvVars: EnvState, DBVars: DBState}
	if i.Error != nil {
		log.Println(i.Error.Error())
		return
	}
	DEBUG_DB = i.EnvVars.DEBUG_DB
	DEBUG_DB_ERROR = i.EnvVars.DEBUG_DB_ERROR

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	server := &http.Server{
		Addr:    i.EnvVars.SERVER_PORT,
		Handler: http.HandlerFunc(Handle_HTTP_Request),
	}
	host, _ := os.Hostname()
	i.openDBConnection()
	if i.Error != nil {
		log.Println(i.Error.Error())
		return
	}
	if i.EnvVars.PERSIST_TEMPLATES {
		InitTemplates = true
		PersistTemplates()
	}
	go func() {
		log.Printf("Starting ICB Workflow Service. Service Endpoint http://%s%s/", host, server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if i.Error = server.Shutdown(ctx); i.Error != nil {
		log.Println(i.Error.Error())
	}
	log.Println("ICB Workflow Service stopped")
	DBConn.Close()
	logfile.Close()
}

package main

import (
	"log"
	"net/http"
	"sync"
)

func (i *Trans) pixmQueryThread(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("PIXm Worker starting")
	i.newPIXmReq()
	log.Println("PIXm Worker done")
	log.Printf("PIXm Worker Status Code %v", i.HTTP.StatusCode)
}
func (i *Trans) cglQueryThread(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("CGL Worker starting")
	i.newCglReq()
	log.Println("CGL Worker done")
	log.Printf("CGL Worker Status Code %v", i.HTTP.StatusCode)
}
func (i *Trans) setPatient() {
	i.Query.Pid = i.Query.Nhs
	i.Query.Pidoid = NHS_OID_DEFAULT
	var wg sync.WaitGroup
	wg.Add(2)
	go i.pixmQueryThread(&wg)
	go i.cglQueryThread(&wg)
	wg.Wait()
	i.newPDSReq()
	i.HTTP.StatusCode = http.StatusOK
	i.setPatientTemplate()
}

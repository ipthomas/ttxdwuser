package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func (i *Trans) newPIXmReq() {
	log.Printf("Creating PIXm PDQ - Server URL %s", i.EnvVars.PIXM_SERVER_URL+"?identifier="+i.Query.Pidoid+"%7C"+i.Query.Pid)
	header := http.Header{}
	header.Set(HTTP_HEADER_CONTENT_TYPE, APPLICATION_JSON)
	req := HTTPRequest{Method: http.MethodGet, Header: header, URL: i.EnvVars.PIXM_SERVER_URL + "?identifier=" + i.Query.Pidoid + "%7C" + i.Query.Pid}
	if i.Error = req.newHTTPRequest(); i.Error == nil {
		if !strings.Contains(string(req.Response), "context deadline exceeded") {
			i.HTTP.ResponseBody = string(req.Response)
			i.HTTP.StatusCode = req.StatusCode
			json.Unmarshal(req.Response, &i.PIXmResponse)
		} else {
			log.Println(req.Response)
			i.HTTP.ResponseBody = "ICB Patient Service Unavailable"
			i.HTTP.StatusCode = http.StatusRequestTimeout
		}
	} else {
		i.HTTP.ResponseBody = "ICB Patient Service Unavailable"
		i.HTTP.StatusCode = http.StatusRequestTimeout
	}
}
func (i *Trans) newPDSReq() {
	var dob, gender, family, url string
	isSearch := false
	if i.PIXmResponse.Total == 1 {
		dob = i.PIXmResponse.Entry[0].Resource.BirthDate
		gender = i.PIXmResponse.Entry[0].Resource.Gender
		family = i.PIXmResponse.Entry[0].Resource.Name[0].Family
	}
	if dob == "" || gender == "" || family == "" {
		dob = i.CGLResponse.Data.Client.BasicDetails.BirthDate
		if len(dob) > 7 {
			dob = dob[:4] + "-" + dob[4:6] + "-" + dob[6:]
		}
		gender = strings.ToLower(i.CGLResponse.Data.Client.BasicDetails.SexAtBirth)
		family = i.CGLResponse.Data.Client.BasicDetails.Name.Family
	}
	if dob != "" && gender != "" && family != "" {
		url = i.EnvVars.PDS_SERVER_URL + "/Patient?family=" + family + "&gender=" + gender + "&birthdate=" + dob
		isSearch = true
	} else {
		url = i.EnvVars.PDS_SERVER_URL + "/Patient/" + i.Query.Nhs
	}
	log.Printf("Creating PDS Request - Server URL %s", url)
	header := http.Header{}
	header.Set(HTTP_HEADER_CONTENT_TYPE, APPLICATION_JSON)
	header.Set("X-Request-ID", NewUuid())
	header.Set(ACCEPT, ALL)
	req := HTTPRequest{Method: http.MethodGet, Header: header, URL: url}
	if i.Error = req.newHTTPRequest(); i.Error == nil {
		if !strings.Contains(string(req.Response), "context deadline exceeded") {
			i.HTTP.ResponseBody = string(req.Response)
			i.HTTP.StatusCode = req.StatusCode
			if isSearch {
				json.Unmarshal(req.Response, &i.PDSSearchResponse)
				i.PDSResponse = "search"
			} else {
				json.Unmarshal(req.Response, &i.PDSRetrieveResponse)
				i.PDSResponse = "retrieve"
			}
		} else {
			log.Println(req.Response)
			i.HTTP.ResponseBody = "PDS Patient Service Unavailable"
			i.HTTP.StatusCode = http.StatusRequestTimeout
		}
	} else {
		i.HTTP.ResponseBody = "PDS Patient Service Unavailable"
		i.HTTP.StatusCode = http.StatusRequestTimeout
	}
}

func (i *Trans) newCglReq() {
	header := http.Header{}
	header.Set(ACCEPT, APPLICATION_JSON)
	header.Set(ENV_CGL_X_API_KEY, i.EnvVars.CGL_SERVER_X_API_KEY)
	header.Set(ENV_CGL_X_API_SECRET, i.EnvVars.CGL_SERVER_X_API_SECRET)
	req := HTTPRequest{Method: http.MethodGet, Header: header, URL: i.EnvVars.CGL_SERVER_URL + "=" + i.Query.Pid}

	if i.Error = req.newHTTPRequest(); i.Error == nil {
		if !strings.Contains(string(req.Response), "context deadline exceeded") || req.StatusCode != 204 {
			i.HTTP.ResponseBody = string(req.Response)
			i.HTTP.StatusCode = req.StatusCode
			json.Unmarshal(req.Response, &i.CGLResponse)
			logStruct(i.CGLResponse)
		} else {
			log.Println(req.Response)
			i.HTTP.ResponseBody = "No CGL Patient found"
			i.HTTP.StatusCode = http.StatusOK
		}
	} else {
		i.HTTP.ResponseBody = "CGL Patient Service Unavailable"
		i.HTTP.StatusCode = http.StatusRequestTimeout
	}
}
func (i *Trans) setNHS() {
	i.newPIXmReq()
	if i.Error == nil {
		rsp := PIXmResponse{}
		if i.Error = json.Unmarshal([]byte(i.HTTP.ResponseBody), &rsp); i.Error == nil {
			for _, entry := range rsp.Entry {
				for _, id := range entry.Resource.Identifier {
					if id.System == NHS_OID_DEFAULT {
						i.Query.Nhs = id.Value
					}
				}
			}
		}
	}
}
func (i *Trans) newDsubCancelReq() {
	i.setDsubCancelTemplate()
	header := http.Header{}
	header.Set(HTTP_HEADER_CONTENT_TYPE, "application/soap+xml;charset=UTF-8")
	header.Set(SOAP_ACTION, SOAP_ACTION_UNSUBSCRIBE_REQUEST)
	header.Set(ACCEPT, "text/xml")
	req := HTTPRequest{Method: http.MethodPost, Header: header, URL: i.EnvVars.DSUB_BROKER_URL, Request: []byte(i.HTTP.ResponseBody), Timeout: i.HTTP.Timeout}
	i.Error = req.newHTTPRequest()
	i.HTTP.ResponseBody = string(req.Response)
	i.HTTP.StatusCode = req.StatusCode
}
func (i *Trans) newDsubSubscribeReq() {
	i.setDsubSubscribeTemplate()
	header := http.Header{}
	header.Set(HTTP_HEADER_CONTENT_TYPE, "application/soap+xml;charset=UTF-8")
	header.Set(SOAP_ACTION, SOAP_ACTION_SUBSCRIBE_REQUEST)
	header.Set(ACCEPT, "text/xml")
	req := HTTPRequest{Method: http.MethodPost, Header: header, URL: i.EnvVars.DSUB_BROKER_URL, Request: []byte(i.HTTP.ResponseBody), Timeout: i.HTTP.Timeout}
	i.Error = req.newHTTPRequest()
	i.HTTP.ResponseBody = string(req.Response)
	i.HTTP.StatusCode = req.StatusCode
}
func (i *HTTPRequest) newHTTPRequest() error {
	var err error
	var req *http.Request
	var rsp *http.Response
	if i.Timeout == 0 {
		i.Timeout = 5
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Timeout)*time.Second)
	defer cancel()
	switch i.Method {
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, i.URL, strings.NewReader(string(i.Request)))
		log.Printf("Request Headers \n%s", i.Header)
		log.Printf("Request Body \n%s", string(i.Request))
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, i.URL, nil)
		log.Printf("Request Headers \n%s", i.Header)
	default:
		err = errors.New("unsupported http method")
		i.StatusCode = http.StatusMethodNotAllowed
	}
	if err == nil {
		req.Header = i.Header
		if rsp, err = http.DefaultClient.Do(req.WithContext(ctx)); err == nil {
			defer rsp.Body.Close()
			i.Response, err = io.ReadAll(rsp.Body)
			i.StatusCode = rsp.StatusCode
		}
	}

	log.Printf("Server Response - Status Code %v \n%s", i.StatusCode, string(i.Response))

	return err
}

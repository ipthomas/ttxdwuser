package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func Handle_HTTP_Request(rsp http.ResponseWriter, req *http.Request) {
	i := Trans{Req: req, EnvVars: EnvState, DBVars: DBState, HTTP: HTTP{StatusCode: http.StatusOK}}
	i.newProxyTrans()
	if i.Error != nil {
		i.setError()
	}
	i.handle_HTTP_Response(rsp)
}
func Handle_AWS_Request(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.SetFlags(log.Lshortfile)
	i := Trans{EnvVars: EnvState, DBVars: DBState, Req: req, HTTP: HTTP{StatusCode: http.StatusOK}}
	i.newProxyTrans()
	if i.Error != nil {
		i.setError()
	}
	return i.handle_AWS_Response()
}
func (i *Trans) newProxyTrans() {
	switch request := i.Req.(type) {
	case events.APIGatewayProxyRequest:
		if i.Error == nil {
			i.HTTP.SourceIP = request.RequestContext.Identity.SourceIP
			i.HTTP.IsBase64 = request.IsBase64Encoded
			i.HTTP.Host = request.RequestContext.DomainName
			i.HTTP.Method = request.HTTPMethod
			i.HTTP.Path = request.Path
			i.HTTP.RequestBody = request.Body
			for k, v := range request.Headers {
				if strings.EqualFold(HTTP_HEADER_CONTENT_TYPE, k) {
					i.HTTP.ReqContentType = v
					break
				}
			}
		}
	case *http.Request:
		i.HTTP.SourceIP = request.RemoteAddr
		i.HTTP.Host, _ = os.Hostname()
		i.HTTP.Method = request.Method
		i.HTTP.Path = request.URL.Path
		if i.HTTP.Method == http.MethodPost {
			bodybytes, err := io.ReadAll(request.Body)
			defer request.Body.Close()
			i.HTTP.RequestBody = string(bodybytes)
			i.Error = err
			if i.Error != nil {
				return
			}
		}
		for k, v := range request.Header {
			if strings.EqualFold(HTTP_HEADER_CONTENT_TYPE, k) {
				i.HTTP.ReqContentType = v[0]
				i.HTTP.IsBase64 = i.HTTP.ReqContentType == APPLICATION_BASE64 || i.HTTP.ReqContentType == IMAGE_BASE64
				break
			}
		}
	}
	i.openDBConnection()
	if i.Error == nil {
		switch i.HTTP.Method {
		case http.MethodGet:
			i.newGetHandler()
		case http.MethodPost:
			i.newPostHandler()
		default:
			log.Printf("HTTP Method %s is not supported", i.HTTP.Method)
		}
	}
}
func (i *Trans) handle_HTTP_Response(rsp http.ResponseWriter) {
	httpHeaders := i.responseHeaders()
	for k, v := range httpHeaders {
		rsp.Header().Set(k, v)
	}
	if httpHeaders[CONTENT_TYPE] == "" {
		rsp.Header().Set(CONTENT_TYPE, TEXT_HTML)
	}
	i.logResponse()
	if i.HTTP.StatusCode != 0 {
		rsp.WriteHeader(i.HTTP.StatusCode)
	}
	rsp.Write([]byte(i.HTTP.ResponseBody))
}
func (i *Trans) handle_AWS_Response() (*events.APIGatewayProxyResponse, error) {
	awsHeaders := i.responseHeaders()
	i.logResponse()
	return &events.APIGatewayProxyResponse{
		StatusCode: i.HTTP.StatusCode,
		Headers:    awsHeaders,
		Body:       string(i.HTTP.ResponseBody),
	}, i.Error
}
func (i *Trans) setError() {
	log.Println(i.Error.Error())
	if i.HTTP.StatusCode == http.StatusOK {
		i.HTTP.StatusCode = http.StatusInternalServerError
	}
	if len(i.HTTP.ResponseBody) == 0 {
		i.HTTP.ResponseBody = i.Error.Error()
	}
}
func (i *Trans) responseHeaders() map[string]string {
	// Reference - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html
	headers := map[string]string{
		//HTTP_HEADER_CACHE_CONTROL: SECURITY_HEADER_CACHE_CONTROL,
		//HTTP_HEADER_CLEAR_SITE_DATA: SECURITY_HEADER_CLEAR_SITE_DATA,
		//HTTP_HEADER_CONTENT_SECURITY_POLICY: SECURITY_HEADER_SECURITY_POLICY,
		//HTTP_HEADER_CROSS_ORIGIN_EMBEDDER_POLICY: SECURITY_HEADER_EMBEDDER_POLICY,
		//HTTP_HEADER_CROSS_ORIGIN_OPENER_POLICY:   SECURITY_HEADER_OPENER_POLICY,
		//HTTP_HEADER_CROSS_ORIGIN_RESOURCE_POLICY: SECURITY_HEADER_RESOURCE_POLICY,
		//HTTP_HEADER_PERMISSIONS_POLICY:           SECURITY_HEADER_PERMISSIONS_POLICY,
		//HTTP_HEADER_REFERRER_POLICY:           SECURITY_HEADER_RESOURCE_POLICY,
		//HTTP_HEADER_STRICT_TRANSPORT_SECURITY: SECURITY_HEADER_STRICT_TRANSPORT_SECURITY,
		//HTTP_HEADER_X_CONTENT_TYPE_OPTIONS: SECURITY_HEADER_X_CONTENT_TYPE_OPTIONS,
		//HTTP_HEADER_X_FRAME_OPTIONS:        SECURITY_HEADER_X_FRAME_OPTIONS,
		//HTTP_HEADER_X_XSS_PROTECTION: SECURITY_HEADER_X_XSS_PROTECTION,
		"Access-Control-Allow-Origin": "*",
		HTTP_HEADER_CONTENT_TYPE:      i.HTTP.RspContentType,
		HTTP_HEADER_SERVER:            "Workflow Server",
	}
	if i.EnvVars.SERVER_NAME != "" {
		headers[HTTP_HEADER_SERVER] = i.EnvVars.SERVER_NAME
	}
	return headers
}
func (i *Trans) logRequest() {
	log.Println("\tRequest Query Params")
	logStruct(i.Query)
	if i.EnvVars.DEBUG_MODE || i.Error != nil {
		log.Println("")
		log.Printf("Request Headers")
		log.Println("-------------")
		logStruct(i.HTTP)
		log.Println("-------------")
	}
}
func (i *Trans) logResponse() {
	log.Printf("Response Status Code %v", i.HTTP.StatusCode)
	if i.EnvVars.DEBUG_MODE || i.Error != nil {
		log.Println("")
		log.Println("Response Headers")
		logStruct(i.responseHeaders())
		log.Println("----------------")
		log.Println("Response")
		log.Println("----------------")
		logStruct(i.HTTP)
		log.Println("-------------")
		if i.Error != nil {
			log.Printf("Error %s", i.Error.Error())
			log.Println("-------------")
		}
	}
}

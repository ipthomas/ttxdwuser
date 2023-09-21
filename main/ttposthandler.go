package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func (i *Trans) newPostHandler() {
	log.Printf("Processing new POST %s Request %s %s %s", i.HTTP.Path, i.Query.User, i.Query.Org, i.Query.Role)
	i.parsePost()
	i.logRequest()
	if i.Error == nil {
		switch {
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_PUBLISHER_CODEMAP):
			i.newCodeMap()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_CONSUMER_DEFINITION):
			i.newXDWDefinition()
			i.Query.Act = "Register Definition"
			i.newUploadFileForm()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_PUBLISHER_META):
			i.newXDWMETA()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_CONSUMER_TEMPLATE):
			i.Query.Act = "Register Template"
			i.newTemplate()
			i.newUploadFileForm()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_PUBLISHER_IMAGE):
			i.newImage()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_PUBLISHER_EVENT):
			i.newUserEvent()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_PUBLISHER_DSUB):
			i.newDSUBEvent()
		case strings.HasSuffix(i.HTTP.Path, HTTP_PATH_PUBLISH):
			i.Query.Act = "Publish Content"
			i.newUploadFileForm()
		default:
			log.Println("invalid post path")
		}
	}
}
func (i *Trans) parsePost() {
	switch i.HTTP.ReqContentType {
	case APPLICATION_XML:
		log.Println("Processing DSUB Broker Notify Message")
	case FORM_URL_ENCODED, FORM_URL_ENCODED_CHARSET_UTF_8:
		if i.Req, i.Error = url.ParseQuery(i.HTTP.RequestBody); i.Error == nil {
			i.Query, i.HTTP.ReqContentType = GetQueryVars(i.Req)
			if i.Query.Action == "" {
				i.Query.Action = i.Query.Act
			}
		}
	default:
		if strings.Contains(i.HTTP.ReqContentType, FORM_MULTI_PART) {
			var form *multipart.Form
			var file multipart.File
			var bodybytes []byte
			boundary := strings.Split(i.HTTP.ReqContentType, "boundary=")[1]
			reader := multipart.NewReader(bytes.NewReader([]byte(i.HTTP.RequestBody)), boundary)
			if form, i.Error = reader.ReadForm(32 << 20); i.Error == nil {
				fileHeader := form.File["file"][0]
				if file, i.Error = fileHeader.Open(); i.Error == nil {
					defer file.Close()
					log.Printf("File name: %s\n", fileHeader.Filename)
					if bodybytes, i.Error = io.ReadAll(file); i.Error == nil {
						i.HTTP.RequestBody = string(bodybytes)
						if len(i.HTTP.RequestBody) < 1 {
							i.Error = errors.New("body is empty")
							i.HTTP.StatusCode = http.StatusBadRequest
						} else {
							if strings.Contains(fileHeader.Filename, "_tmplt.") {
								i.Query.Template = strings.Split(fileHeader.Filename, ".")[0]
							} else {
								if strings.Contains(fileHeader.Filename, "_def.") || strings.Contains(fileHeader.Filename, "_meta.") {
									i.Query.Pathway = strings.Split(fileHeader.Filename, ".")[0]
								} else {
									if strings.HasSuffix(fileHeader.Filename, ".zip") {
										log.Printf("Publishing %s to SCR", fileHeader.Filename)
										i.Query.Name = fileHeader.Filename
										i.newS3Object()
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

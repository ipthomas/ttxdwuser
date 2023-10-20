package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var DBState = &DBVars{}
var EnvState = &EnvVars{}
var DEBUG_DB = false
var DEBUG_DB_ERROR = false
var Environ = ""
var LOC *time.Location

func main() {
	Environ = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	SetServiceState()
	LOC, _ = time.LoadLocation("Europe/London")

	if Environ == "" {
		StartServer()
	} else {
		lambda.Start(Handle_AWS_Request)
	}
}

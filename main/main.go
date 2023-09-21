package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var DBState = &DBVars{}
var EnvState = &EnvVars{}
var DEBUG_DB = false
var DEBUG_DB_ERROR = false
var Environ = ""

func main() {
	Environ = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	SetServiceState()
	if Environ == "" {
		StartServer()
	} else {
		lambda.Start(Handle_AWS_Request)
	}
}

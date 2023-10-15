package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func SetServiceState() {
	if Environ != "" {
		log.Println("Setting Environment Variables from AWS Environment Variables")
		for _, awsenvar := range os.Environ() {
			var field reflect.Value
			if strings.HasPrefix(awsenvar, "DB_") {
				field = reflect.ValueOf(DBState).Elem().FieldByName(strings.Split(awsenvar, "=")[0])
			} else {
				field = reflect.ValueOf(EnvState).Elem().FieldByName(strings.Split(awsenvar, "=")[0])
			}
			if field.IsValid() && field.CanSet() {
				if strings.Split(awsenvar, "=")[0] == ENV_DEBUG_MODE || strings.Split(awsenvar, "=")[0] == ENV_PERSIST_TEMPLATES || strings.Split(awsenvar, "=")[0] == ENV_DEBUG_DB || strings.Split(awsenvar, "=")[0] == ENV_DEBUG_DB_ERROR || strings.Split(awsenvar, "=")[0] == ENV_CALENDAR_MODE_DISABLED {
					envBool, _ := strconv.ParseBool(strings.Split(awsenvar, "=")[1])
					field.Set(reflect.ValueOf(envBool))
				} else {
					field.Set(reflect.ValueOf(strings.Split(awsenvar, "=")[1]))
				}
			}
		}
	} else {
		log.Println("Setting Environment Variables from local config file envvars.json")
		if dbvarsmap, envvarsmap, err := loadEnvVarsFile("envvars.json"); err == nil {
			for k, v := range dbvarsmap {
				field := reflect.ValueOf(DBState).Elem().FieldByName(k)
				if field.IsValid() && field.CanSet() {
					field.Set(reflect.ValueOf(v))
				}
			}
			for k, v := range envvarsmap {
				field := reflect.ValueOf(EnvState).Elem().FieldByName(k)
				if field.IsValid() && field.CanSet() {
					if k == ENV_DEBUG_MODE || k == ENV_PERSIST_TEMPLATES || k == ENV_DEBUG_DB || k == ENV_DEBUG_DB_ERROR || k == ENV_PERSIST_DEFINITIONS || k == ENV_CALENDAR_MODE_DISABLED {
						envBool, _ := strconv.ParseBool(v)
						field.Set(reflect.ValueOf(envBool))
					} else {
						field.Set(reflect.ValueOf(v))
					}
				}
			}
		}
	}
	if EnvState.DEBUG_MODE {
		logStruct(EnvState)
	}
}
func loadEnvVarsFile(filename string) (map[string]string, map[string]string, error) {
	if filename == "" {
		filename = "envvars.json"
	}
	fileEnvvars := make(map[string]string)
	dbVars := make(map[string]string)
	envVars := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
		return dbVars, envVars, err
	}
	defer file.Close()
	filebytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
		return dbVars, envVars, err
	}
	if err = json.Unmarshal(filebytes, &fileEnvvars); err != nil {
		log.Println(err.Error())
		return dbVars, envVars, err
	}
	for k, v := range fileEnvvars {
		if v != "" {
			if strings.HasPrefix(k, "DB_") {
				dbVars[k] = v
			} else {
				envVars[k] = v
			}
		}
	}
	if _, ok := envVars[ENV_SERVER_PORT]; !ok {
		envVars[ENV_SERVER_PORT] = ":8080"
	}
	if !strings.HasPrefix(envVars[ENV_SERVER_PORT], ":") {
		envVars[ENV_SERVER_PORT] = ":" + envVars[ENV_SERVER_PORT]
	}
	if _, ok := envVars[ENV_SERVER_URL]; !ok {
		envVars[ENV_SERVER_URL] = "http://localhost" + envVars[ENV_SERVER_PORT] + "/"
	} else {
		envVars[ENV_SERVER_URL] = envVars[ENV_SERVER_URL] + envVars[ENV_SERVER_PORT] + "/"
	}
	log.Printf("Set XDW Server URL %s", envVars[ENV_SERVER_URL])
	return dbVars, envVars, err
}

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn       *sql.DB
	cachedIDMaps = []IdMap{}
	cached       time.Time
)

type DBInterface interface {
	newEvent() error
}

func NewDBEvent(i DBInterface) error {
	return i.newEvent()
}
func (i *Trans) openDBConnection() {
	if DBConn == nil {
		conn := DBConnection{DBUser: i.DBVars.DB_USER, DBPassword: i.DBVars.DB_PASSWORD, DBHost: i.DBVars.DB_HOST, DBPort: i.DBVars.DB_PORT, DBName: i.DBVars.DB_NAME}
		err := conn.newDBEvent()
		if err != nil {
			log.Println(i.Error.Error())
		}
	}
}

// DBConnection

func (i *DBConnection) newDBEvent() error {
	var err error
	i.setDBCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&timeout=%s&readTimeout=%s",
		i.DBUser,
		i.DBPassword,
		i.DBHost+i.DBPort,
		i.DBName,
		i.DBTimeout,
		i.DBReadTimeout)
	log.Printf("Opening DB Connection to mysql instance\nUser: %s\nHost: %s\nPort: %s\nName: %s\nConnect Timeout %s\nRead Timeout %s", i.DBUser, i.DBHost, i.DBPort, i.DBName, i.DBTimeout, i.DBReadTimeout)
	DBConn, err = sql.Open("mysql", dsn)
	if err == nil {
		log.Println("Opened Database")
	}
	return err
}
func (i *DBConnection) setDBCredentials() {
	if i.DBPort == "" {
		i.DBPort = "3306"
	}
	if !strings.HasPrefix(i.DBPort, ":") {
		i.DBPort = ":" + i.DBPort
	}
	if i.DBName == "" {
		i.DBName = "tuk"
	}
	if i.DBTimeout == "" {
		i.DBTimeout = "2"
	}
	if !strings.HasSuffix(i.DBTimeout, "s") {
		i.DBTimeout = i.DBTimeout + "s"
	}
	if i.DBReadTimeout == "" {
		i.DBReadTimeout = "5"
	}
	if !strings.HasSuffix(i.DBReadTimeout, "s") {
		i.DBReadTimeout = i.DBReadTimeout + "s"
	}
}

// Subscriptions
func (i *Subscriptions) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_SUBSCRIPTIONS
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.Subscriptions) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, SUBSCRIPTIONS, reflectStruct(reflect.ValueOf(i.Subscriptions[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		for rows.Next() {
			sub := Subscription{}
			if err := rows.Scan(&sub.Id, &sub.Created, &sub.BrokerRef, &sub.Pathway, &sub.Topic, &sub.Expression, &sub.Email, &sub.NhsId, &sub.User, &sub.Org, &sub.Role); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.Subscriptions = append(i.Subscriptions, sub)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// Events
func GetTaskNotes(pwy string, nhsid string, taskid int, ver int) ([]byte, error) {
	cmnts := Comments{}
	evs := Events{Action: SELECT}
	ev := Event{Pathway: pwy, Nhs: nhsid, Taskid: taskid, Version: ver}
	evs.Events = append(evs.Events, ev)
	err := NewDBEvent(&evs)
	if err == nil && evs.Count > 0 {
		for _, note := range evs.Events {
			if note.Id != 0 && note.Comments != "" && note.Comments != "None" {
				cmnts.Comment = append(cmnts.Comment, Comment{Taskid: note.Taskid, Note: note.Comments})
			}
		}
	}
	notes, _ := json.MarshalIndent(cmnts.Comment, "", "  ")
	return notes, err
}
func (i *Events) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_EVENTS
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.Events) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, EVENTS, reflectStruct(reflect.ValueOf(i.Events[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		for rows.Next() {
			ev := Event{}
			if err := rows.Scan(&ev.Id, &ev.Creationtime, &ev.Eventtype, &ev.Docname, &ev.Classcode, &ev.Confcode, &ev.Formatcode, &ev.Facilitycode, &ev.Practicecode, &ev.Speciality, &ev.Expression, &ev.Authors, &ev.Xdsdocentryuid, &ev.Repositoryuid, &ev.Nhs, &ev.User, &ev.Org, &ev.Role, &ev.Topic, &ev.Pathway, &ev.Comments, &ev.Version, &ev.Taskid); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.Events = append(i.Events, ev)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// Workflows
func (i *Workflows) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_WORKFLOWS
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.Workflows) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, WORKFLOWS, reflectStruct(reflect.ValueOf(i.Workflows[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			workflow := Workflow{}
			if err := rows.Scan(&workflow.Id, &workflow.Pathway, &workflow.NHSId, &workflow.Created, &workflow.XDW_Key, &workflow.XDW_UID, &workflow.XDW_Doc, &workflow.XDW_Def, &workflow.Version, &workflow.Published, &workflow.Status); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.Workflows = append(i.Workflows, workflow)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// XDWs
func getXDW(name string, ismeta bool) string {
	xdws := XDWS{Action: SELECT}
	xdw := XDW{Name: name, IsXDSMeta: ismeta}
	xdws.XDW = append(xdws.XDW, xdw)
	if err := xdws.newEvent(); err != nil {
		log.Println(err.Error())
		return ""
	}
	if xdws.Count == 1 {
		return xdws.XDW[1].XDW
	}
	return ""
}
func getXDWs() XDWS {
	xdws := XDWS{Action: SELECT}
	xdw := XDW{IsXDSMeta: false}
	xdws.XDW = append(xdws.XDW, xdw)
	if err := xdws.newEvent(); err != nil {
		log.Println(err.Error())
	}
	return xdws
}

func (i *XDWS) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_XDWS
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.XDW) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, "xdws", reflectStruct(reflect.ValueOf(i.XDW[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			xdw := XDW{}
			if err := rows.Scan(&xdw.Id, &xdw.Name, &xdw.IsXDSMeta, &xdw.XDW); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.XDW = append(i.XDW, xdw)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// Workflowstates
func (i *WorkflowStates) newEvent() error {
	var err error
	var stmntStr = "SELECT * FROM workflowstate"
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.Workflowstate) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, "workflowstate", reflectStruct(reflect.ValueOf(i.Workflowstate[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			workflow := Workflowstate{}
			if err := rows.Scan(&workflow.WorkflowId, &workflow.Pathway, &workflow.NHSId, &workflow.Version, &workflow.Published, &workflow.Created, &workflow.CreatedBy, &workflow.Status, &workflow.CompleteBy, &workflow.LastUpdate, &workflow.Owner, &workflow.Overdue, &workflow.Escalated, &workflow.TargetMet, &workflow.InProgress, &workflow.Duration, &workflow.TimeRemaining); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.Workflowstate = append(i.Workflowstate, workflow)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// Templates
func PersistTemplates() {
	if tmpltFiles, err := GetFolderFiles("./templates/"); err == nil {
		for _, file := range tmpltFiles {
			if strings.HasSuffix(file.Name(), ".html") || strings.HasSuffix(file.Name(), ".json") || strings.HasSuffix(file.Name(), ".txt") || strings.HasSuffix(file.Name(), ".xml") {
				filebytes := loadFile(file, "./templates/")
				if len(filebytes) > 0 {
					persistTemplate("system", strings.Split(file.Name(), ".")[0], string(filebytes))
					log.Printf("Persisted Template %s", file.Name())
				}
			}
		}
	} else {
		log.Println(err.Error())
	}
}
func persistTemplate(user string, templatename string, templatestr string) error {
	tmplts := Templates{Action: DELETE}
	tmplt := Template{Name: templatename, User: user}
	tmplts.Templates = append(tmplts.Templates, tmplt)
	tmplts.newEvent()
	log.Printf("Persisting Template %s", templatename)
	tmplts = Templates{Action: INSERT}
	tmplt = Template{Name: templatename, User: user, Template: templatestr}
	tmplts.Templates = append(tmplts.Templates, tmplt)
	err := tmplts.newEvent()
	if err == nil {
		log.Printf("Persisted Template %s", templatename)
	}
	return err
}
func (i *Templates) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_TEMPLATES
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.Templates) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, TEMPLATES, reflectStruct(reflect.ValueOf(i.Templates[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			tmplt := Template{}
			if err := rows.Scan(&tmplt.Id, &tmplt.Name, &tmplt.Template, &tmplt.User); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.Templates = append(i.Templates, tmplt)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// Idmaps
func GetMappedValue(user string, localid string) string {
	if user == "" {
		user = "system"
	}
	duration := time.Duration(1) * time.Minute
	expires := cached.Add(duration)
	if len(cachedIDMaps) == 0 || time.Now().After(expires) {
		idmaps := IdMaps{Action: SELECT}
		if err := idmaps.newEvent(); err != nil {
			log.Println(err.Error())
		}
		cachedIDMaps = idmaps.LidMap
		cached = time.Now()
	}
	for _, v := range cachedIDMaps {
		if v.User == user && v.Lid == localid {
			return strings.TrimSpace(v.Mid)
		}
	}
	if user != "system" {
		user = "system"
		for _, v := range cachedIDMaps {
			if v.User == user && v.Lid == localid {
				return strings.TrimSpace(v.Mid)
			}
		}
	}
	return strings.TrimSpace(localid)
}
func GetLocalValue(user string, mid string) string {
	if user == "" {
		user = "system"
	}
	idmaps := IdMaps{Action: SELECT}
	idmap := IdMap{User: user}
	idmaps.LidMap = append(idmaps.LidMap, idmap)
	if err := idmaps.newEvent(); err != nil {
		log.Println(err.Error())
	}
	for _, idmap := range idmaps.LidMap {
		if idmap.Mid == mid && idmap.User == user {
			return idmap.Lid
		}
	}
	return strings.TrimSpace(mid)
}
func (i *IdMaps) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_IDMAPS
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.LidMap) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, ID_MAPS, reflectStruct(reflect.ValueOf(i.LidMap[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			idmap := IdMap{}
			if err := rows.Scan(&idmap.Id, &idmap.Lid, &idmap.Mid, &idmap.User); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.LidMap = append(i.LidMap, idmap)
			i.Cnt = i.Cnt + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

// Statics
func (i *Statics) newEvent() error {
	var err error
	var stmntStr = SQL_DEFAULT_STATICS
	var rows *sql.Rows
	var vals []interface{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	if len(i.Static) > 0 {
		if stmntStr, vals, err = createPreparedStmnt(i.Action, STATICS, reflectStruct(reflect.ValueOf(i.Static[0]))); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	sqlStmnt, err := DBConn.PrepareContext(ctx, stmntStr)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStmnt.Close()

	if i.Action == SELECT {
		rows, err = setRows(ctx, sqlStmnt, vals)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			s := Static{}
			if err := rows.Scan(&s.Id, &s.Name, &s.Content); err != nil {
				switch {
				case err == sql.ErrNoRows:
					return nil
				default:
					log.Println(err.Error())
					return err
				}
			}
			i.Static = append(i.Static, s)
			i.Count = i.Count + 1
		}
	} else {
		i.LastInsertId, err = setLastID(i.Action, ctx, sqlStmnt, vals)
	}
	return err
}

func reflectStruct(i reflect.Value) map[string]interface{} {
	params := make(map[string]interface{})
	structType := i.Type()
	for f := 0; f < i.NumField(); f++ {
		field := structType.Field(f)
		fieldName := field.Name
		fieldType := field.Type
		switch fieldType.Kind() {
		case reflect.Int:
			val := i.Field(f).Interface().(int)
			if val > 0 {
				params[strings.ToLower(fieldName)] = val
			}
		case reflect.Int64:
			val := i.Field(f).Interface().(int64)
			if val > 0 {
				params[strings.ToLower(fieldName)] = val
			}
		case reflect.Bool:
			val := i.Field(f).Interface().(bool)
			params[strings.ToLower(fieldName)] = val
		case reflect.String:
			val := i.Field(f).Interface().(string)
			if val != "" {
				params[strings.ToLower(fieldName)] = strings.TrimSpace(val)
			}
		default:
			log.Printf("Field %s has an unknown type %v\n", fieldName, fieldType.Kind())
		}
	}
	return params
}
func createPreparedStmnt(action string, table string, params map[string]interface{}) (string, []interface{}, error) {
	var vals []interface{}
	stmntStr := "SELECT * FROM " + table
	if len(params) > 0 {
		switch action {
		case SELECT:
			var paramStr string
			stmntStr = stmntStr + " WHERE "
			for param, val := range params {
				paramStr = paramStr + param + "= ? AND "
				vals = append(vals, val)
			}
			paramStr = strings.TrimSuffix(paramStr, " AND ")
			stmntStr = stmntStr + paramStr
		case INSERT:
			var paramStr string
			var qStr string
			stmntStr = "INSERT INTO " + table + " ("
			for param, val := range params {
				paramStr = paramStr + param + ", "
				qStr = qStr + "?, "
				vals = append(vals, val)
			}
			paramStr = strings.TrimSuffix(paramStr, ", ") + ") VALUES ("
			qStr = strings.TrimSuffix(qStr, ", ")
			stmntStr = stmntStr + paramStr + qStr + ")"
		case DEPRECATE:
			switch table {
			case WORKFLOWS:
				stmntStr = "UPDATE workflows SET version = version + 1 WHERE pathway=? AND nhsid=?"
				vals = append(vals, params["pathway"])
				vals = append(vals, params["nhsid"])
				log.Printf("Deprecating %s Workflow for NHS %s", params["pathway"], params["nhsid"])
			case EVENTS:
				stmntStr = "UPDATE events SET version = version + 1 WHERE pathway=? AND nhs=?"
				vals = append(vals, params["pathway"])
				vals = append(vals, params["nhs"])
				log.Printf("Deprecating %s Events for NHS %s", params["pathway"], params["nhs"])
			}
		case ROLLBACK:
			switch table {
			case WORKFLOWS:
				stmntStr = "UPDATE workflows SET version = version -1 where pathway=? AND nhsid=?"
				vals = append(vals, params["pathway"])
				vals = append(vals, params["nhsid"])
			case EVENTS:
				stmntStr = "UPDATE events SET version = version - 1 WHERE pathway=? AND nhs=?"
				vals = append(vals, params["pathway"])
				vals = append(vals, params["nhsid"])
			}
		case UPDATE:
			switch table {
			case WORKFLOWS:
				stmntStr = "UPDATE workflows SET xdw_doc = ?, published = ?, status = ? WHERE pathway = ? AND nhsid = ? AND version = ?"
				vals = append(vals, params["xdw_doc"])
				vals = append(vals, params["published"])
				vals = append(vals, params["status"])
				vals = append(vals, params["pathway"])
				vals = append(vals, params["nhsid"])
				vals = append(vals, params["version"])
			case ID_MAPS:
				stmntStr = "UPDATE idmaps SET "
				var paramStr string
				for param, val := range params {
					if val != "" && param != "id" {
						paramStr = paramStr + param + "= ?, "
						vals = append(vals, val)
					}
				}
				vals = append(vals, params["id"])
				paramStr = strings.TrimSuffix(paramStr, ", ")
				stmntStr = stmntStr + paramStr + " WHERE id = ?"
			}
		case DELETE:
			stmntStr = "DELETE FROM " + table + " WHERE "
			var paramStr string
			for param, val := range params {
				paramStr = paramStr + param + "= ? AND "
				vals = append(vals, val)
			}
			paramStr = strings.TrimSuffix(paramStr, " AND ")
			stmntStr = stmntStr + paramStr
		}
	}
	if DEBUG_DB {
		log.Println(stmntStr)
	}
	if DEBUG_DB_ERROR {
		logStruct(vals)
	}
	return stmntStr, vals, nil
}
func setRows(ctx context.Context, sqlStmnt *sql.Stmt, vals []interface{}) (*sql.Rows, error) {
	if len(vals) > 0 {
		return sqlStmnt.QueryContext(ctx, vals...)
	} else {
		return sqlStmnt.QueryContext(ctx)
	}
}
func setLastID(act string, ctx context.Context, sqlStmnt *sql.Stmt, vals []interface{}) (int, error) {
	if len(vals) > 0 {
		sqlrslt, err := sqlStmnt.ExecContext(ctx, vals...)
		if err != nil {
			log.Println(err.Error())
			return 0, err
		}
		switch act {
		case UPDATE, DEPRECATE, ROLLBACK:
			rows, err := sqlrslt.RowsAffected()
			if err != nil {
				log.Println(err.Error())
				return 0, err
			} else {
				return int(rows), nil
			}
		default:
			id, err := sqlrslt.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return 0, err
			} else {
				return int(id), nil
			}
		}

	}
	return 0, nil
}

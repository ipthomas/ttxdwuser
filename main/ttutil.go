package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

var (
	ServerName, _ = os.Hostname()
	SeedRoot      = "1.2.40.0.13.1.1.3542466645."
	IdSeed        = getIdIncrementSeed(5)
	CodeSystem    = make(map[string]string)
)

func GetQueryVars(req interface{}) (QueryVars, string) {
	var reqFormat = ""
	query := QueryVars{}
	switch request := req.(type) {
	case events.APIGatewayProxyRequest:
		for k, v := range request.QueryStringParameters {
			if format := query.setQueryKeyValue(k, v); format != "" {
				reqFormat = format
			}
		}
	case *http.Request:
		for k, v := range request.URL.Query() {
			if format := query.setQueryKeyValue(k, v[0]); format != "" {
				reqFormat = format
			}
		}
	case url.Values:
		for k, v := range request {
			if format := query.setQueryKeyValue(k, v[0]); format != "" {
				reqFormat = format
			}
		}
	}
	if reqFormat != "" {
		log.Printf("Response Content-Type Requested - %s", reqFormat)
	}
	return query, reqFormat
}
func (i *QueryVars) setQueryKeyValue(k string, v string) string {
	key := capitalizeFirstChar(strings.TrimPrefix(k, "_"))
	if key == "Format" {
		if reflect.ValueOf(i).Elem().FieldByName(key).IsValid() && reflect.ValueOf(i).Elem().FieldByName(key).CanSet() {
			reflect.ValueOf(i).Elem().FieldByName(key).Set(reflect.ValueOf(v))
			return v
		}
	}
	if reflect.ValueOf(i).Elem().FieldByName(key).IsValid() && reflect.ValueOf(i).Elem().FieldByName(key).CanSet() {
		reflect.ValueOf(i).Elem().FieldByName(key).Set(reflect.ValueOf(v))
		return ""
	}
	log.Printf("Invalid query param %s", key)
	return ""
}
func capitalizeFirstChar(s string) string {
	if len(s) > 1 {
		firstChar := strings.ToUpper(string(s[0]))
		return firstChar + s[1:]
	}
	return s
}

// CreateLog checks if the log folder exists and creates it if not. It then checks for a subfolder for the current year i.e. 2022 and creates it if it does not exist. It then checks for a log file with a name equal to the current day and month and extension .log i.e. 0905.log. If it exists log output is appended to the existing file otherwise a new log file is created.
func createLog(log_Folder string) *os.File {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	if _, err := os.Stat(log_Folder); errors.Is(err, fs.ErrNotExist) {
		if e2 := os.Mkdir(log_Folder, 0700); e2 != nil {
			log.Println(err.Error())
			return nil
		}
	}
	if _, err := os.Stat(log_Folder + "/" + Tuk_Year()); errors.Is(err, fs.ErrNotExist) {
		if e2 := os.Mkdir(log_Folder+"/"+Tuk_Year(), 0700); e2 != nil {
			log.Println(err.Error())
			return nil
		}
	}
	if _, err := os.Stat(log_Folder + "/" + Tuk_Year() + "/" + Tuk_Month()); errors.Is(err, fs.ErrNotExist) {
		if e2 := os.Mkdir(log_Folder+"/"+Tuk_Year()+"/"+Tuk_Month(), 0700); e2 != nil {
			log.Println(err.Error())
			return nil
		}
	}
	if _, err := os.Stat(log_Folder + "/" + Tuk_Year() + "/" + Tuk_Month() + "/" + Tuk_Day()); errors.Is(err, fs.ErrNotExist) {
		if e2 := os.Mkdir(log_Folder+"/"+Tuk_Year()+"/"+Tuk_Month()+"/"+Tuk_Day(), 0700); e2 != nil {
			log.Println(err.Error())
			return nil
		}
	}
	logFile, err := os.OpenFile(log_Folder+"/"+Tuk_Year()+"/"+Tuk_Month()+"/"+Tuk_Day()+"/"+Tuk_Hour()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println("Using log file - " + logFile.Name())
	log.SetOutput(logFile)
	log.Println("-----------------------------------------------------------------------------------")
	return logFile
}

// Log takes any struc as input and logs out the struc as a json string
func logStruct(struc interface{}) {
	b, _ := json.MarshalIndent(struc, "", "  ")
	log.Println("\t" + string(b))
}
func GetSummerBankHoliday(year int) time.Time {
	aug31 := time.Date(year, time.August, 31, 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(aug31.Weekday())
	daysToSubtract := (7 - dayOfWeek) % 7
	summerBankHoliday := aug31.AddDate(0, 0, -daysToSubtract)
	return summerBankHoliday
}
func GetSpringBankHoliday(year int) time.Time {
	may31 := time.Date(year, time.May, 31, 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(may31.Weekday())
	daysToSubtract := (7 - dayOfWeek) % 7
	springBankHoliday := may31.AddDate(0, 0, -daysToSubtract)
	return springBankHoliday
}
func GetEarlyMayBankHoliday(year int) time.Time {
	may1 := time.Date(year, time.May, 1, 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(may1.Weekday())
	daysToAdd := (8 - dayOfWeek) % 7
	earlyMayBankHoliday := may1.AddDate(0, 0, daysToAdd)
	return earlyMayBankHoliday
}
func GetEasterDate(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (b - f + 1) % 3
	i := (19*a + b - d - g + 15) / 30
	k := (19*a + b - d - g + 15) % 30
	L := (2*e + 2*h + i - k - c + 4) / 7
	m := (2*e + 2*h + i - k - c + 4) % 7
	month := i - k + L + 114
	day := k + m + 1
	if month > 12 {
		month -= 12
	}
	easterDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return easterDate
}
func isSummerBankHoliday(currentDate time.Time) bool {
	return currentDate == GetSummerBankHoliday(currentDate.Year())
}
func isSpringBankHoliday(currentDate time.Time) bool {
	return currentDate == GetSpringBankHoliday(currentDate.Year())
}
func isEarlyMayBankHoliday(currentDate time.Time) bool {
	return currentDate == GetEarlyMayBankHoliday(currentDate.Year())
}
func isGoodFriday(currentDate time.Time) bool {
	easterDate := GetEasterDate(currentDate.Year())
	return currentDate == easterDate.AddDate(0, 0, -2)
}
func isSaturday(currentDate time.Time) bool {
	return currentDate.Weekday() == time.Saturday
}
func isWeekend(currentDate time.Time) bool {
	return currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday
}
func isXmasDay(currentDate time.Time) bool {
	return currentDate.Month() == time.December && currentDate.Day() == 25
}
func isBoxingDay(currentDate time.Time) bool {
	return currentDate.Month() == time.December && currentDate.Day() == 26
}
func isNewYearsDay(currentDate time.Time) bool {
	return currentDate.Month() == time.January && currentDate.Day() == 1
}
func (i *Trans) CalendarMode(startDate, endDate time.Time, isDuration bool) time.Time {
	debug := i.EnvVars.DEBUG_MODE
	if i.EnvVars.CALENDAR_MODE != "calendardays" {
		adjust := 0
		newEndDate := endDate
		for currentDate := startDate; currentDate.Before(newEndDate) || currentDate.Equal(newEndDate); currentDate = currentDate.AddDate(0, 0, 1) {
			if isXmasDay(currentDate) || isBoxingDay(currentDate) || isNewYearsDay(currentDate) {
				newEndDate = newEndDate.AddDate(0, 0, 1)
				if debug {
					log.Printf("%v is a xmas period holiday. End Date Adjusted 1 Day to %v", currentDate, newEndDate)
				}
				adjust = adjust + 1
			}
			if isWeekend(currentDate) {
				if isSaturday(currentDate) && currentDate == newEndDate {
					adjust = adjust + 2
					newEndDate = newEndDate.AddDate(0, 0, 2)
					if debug {
						log.Printf("%v is a Saturday and the workflow end date. End Date Adjusted 2 Days to %v", currentDate, newEndDate)
					}
					currentDate = newEndDate
				} else {
					adjust = adjust + 1
					newEndDate = newEndDate.AddDate(0, 0, 1)
					if debug {
						log.Printf("%v is a Weekend. End Date Adjusted 1 Day to %v", currentDate, newEndDate)
					}
				}
			} else {
				if isGoodFriday(currentDate) {
					if currentDate == newEndDate {
						adjust = adjust + 4
						newEndDate = newEndDate.AddDate(0, 0, 4)
						if debug {
							log.Printf("%v is Good Friday and the workflow end date. End Date Adjusted 4 Days to %v", currentDate, newEndDate)
						}
						currentDate = newEndDate
					} else {
						adjust = adjust + 2
						newEndDate = newEndDate.AddDate(0, 0, 4)
						if debug {
							log.Printf("%v is Good Friday. End Date Adjusted 2 Days to %v", currentDate, newEndDate)
						}
					}
				}
				if isEarlyMayBankHoliday(currentDate) || isSpringBankHoliday(currentDate) || isSummerBankHoliday(currentDate) {
					adjust = adjust + 1
					if debug {
						log.Printf("%v is a Bank Holiday. Adjust Total = %v", currentDate, adjust)
					}
				}
			}
		}
		if isDuration {
			if debug {
				log.Printf("Workflow Duration adjusted to account for Weekends and UK Public Holidays by %v days ", adjust)
			}
			return startDate.AddDate(0, 0, adjust)
		} else {
			if debug {
				log.Printf("Workflow End Date adjusted to account for Weekends and UK Public Holidays by %v days ", adjust)
			}
			return newEndDate
		}
	} else {
		if isDuration {
			return startDate
		} else {
			return endDate
		}
	}
}

// OHT_FutureDate takes a 'start date' and a period in the future as a string containing an OASIS Human Task api function eg. day(x). It returns x days in the future from the `start date`. Valid periods are min(x),hour(x),day(x),month(x) and year(x)
func (i *Trans) OHT_FutureDate(startdate time.Time, htDate string) time.Time {
	fd := startdate
	if strings.Contains(htDate, "(") && strings.Contains(htDate, ")") {
		periodstr := strings.Split(htDate, "(")[0]
		periodtime := GetIntFromString(strings.Split(strings.Split(htDate, "(")[1], ")")[0])
		switch periodstr {
		case "sec":
			fd = GetFutureDate(startdate, 0, 0, 0, 0, 0, periodtime)
		case "min":
			fd = GetFutureDate(startdate, 0, 0, 0, 0, periodtime, 0)
		case "hour":
			fd = GetFutureDate(startdate, 0, 0, 0, periodtime, 0, 0)
		case "day":
			fd = GetFutureDate(startdate, 0, 0, periodtime, 0, 0, 0)
		case "month":
			fd = GetFutureDate(startdate, 0, periodtime, 0, 0, 0, 0)
		case "year":
			fd = GetFutureDate(startdate, periodtime, 0, 0, 0, 0, 0)
		}
	}
	return i.CalendarMode(startdate, fd, false)
}

// Time_Now returns the current time for location Europe/London.
func Time_Now() string {
	location, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Println(err.Error())
		return time.Now().String()
	}
	return time.Now().In(location).Format(time.RFC3339)
}

func PrettyTime(time string) string {
	if time != "" {
		return GetTimeFromString(time).String()
	}
	return time
}

// TUK_Hour returns the current hour as a 2 digit string
func Tuk_Hour() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Hour())
}

// TUK_Day returns the current day as a 2 digit string
func Tuk_Day() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Day())
}

// TUK_Year returns the current year as a 4 digit string
func Tuk_Year() string {
	return fmt.Sprintf("%d",
		time.Now().Local().Year())
}

// TUK_Month returns the current month as a 2 digit string
func Tuk_Month() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Month())
}

// NewUuid returns a random UUID as a string
func NewUuid() string {
	u := uuid.New()
	return u.String()
}

// GetStringFromInt takes a int input and returns a string of that value.
func GetStringFromInt(i int) string {
	return strconv.Itoa(i)
}

// GetIntFromString takes a string input with an integer value and returns an int of that value. If the input is not numeric, 0 is returned
func GetIntFromString(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

// getDocId
func getDocId(workflowInstanceId string) string {
	return strings.Split(workflowInstanceId, "^^^^")[0]
}

// Substr takes a string input and returns the rune (Substring) defined by the start and length in th start and length input values
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}

// GetXMLNodeList takes an xml message as input and returns the xml node list as a string for the node input value provide
func GetXMLNodeList(message string, node string) string {
	if strings.Contains(message, node) {
		var nodeopen = "<" + node
		var nodeclose = "</" + node + ">"
		log.Println("Searching for XML Element: " + nodeopen + ">")
		var start = strings.Index(message, nodeopen)
		var end = strings.Index(message, nodeclose) + len(nodeclose)
		m := message[start:end]
		log.Println("Extracted XML Element Nodelist")
		return m
	}
	log.Println("Message does not contain Element : " + node)
	return ""
}

// PrettyAuthorInstitution takes a string input (XDS Author.Institution format) and returns a string with just the Institution name
func PrettyAuthorInstitution(institution string) string {
	if strings.Contains(institution, "^") {
		return strings.Split(institution, "^")[0] + ","
	}
	return institution
}

// PrettyAuthorPerson takes a string input (XDS Author.Person format) and returns a string with the person last and first names
func PrettyAuthorPerson(author string) string {
	if strings.Contains(author, "^") {
		authorsplit := strings.Split(author, "^")
		if len(authorsplit) > 2 {
			log.Println("Split Author " + authorsplit[1] + " " + authorsplit[2])
			return authorsplit[1] + " " + authorsplit[2]
		}
		if len(authorsplit) > 1 {
			return authorsplit[1]
		}
	}
	log.Println("Parsed Author " + author)
	return author
}

// GetFolderFiles takes a string input of the complete folder path and returns a fs.DirEntry
func GetFolderFiles(folder string) ([]fs.DirEntry, error) {
	var err error
	var f *os.File
	var fileInfo []fs.DirEntry
	f, err = os.Open(folder)
	if err != nil {
		log.Println(err.Error())
		return fileInfo, err
	}
	fileInfo, err = f.ReadDir(-1)
	f.Close()
	if err != nil {
		log.Println(err.Error())
	}
	return fileInfo, err
}
func loadFile(file fs.DirEntry, folder string) []byte {
	var fileBytes []byte
	var err error
	fileBytes, err = os.ReadFile(folder + file.Name())
	if err != nil {
		log.Println(err.Error())
	}
	return fileBytes
}
func getIdIncrementSeed(len int) int {
	return GetIntFromString(Substr(GetStringFromInt(time.Now().Nanosecond()), 0, len))
}
func GetTimeFromString(timestr string) time.Time {
	timestr = strings.Split(timestr, ".")[0]
	timestr = strings.Split(timestr, " +")[0]
	var err error
	var rsptime time.Time
	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Println(err.Error())
		return rsptime
	}
	if !strings.Contains(timestr, "T") {
		rsptime, err = time.ParseInLocation("2006-01-02 15:04:05", timestr, loc)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		rsptime, err = time.ParseInLocation(time.RFC3339, timestr, loc)
		if err != nil {
			log.Println(err.Error())
			rsptime, err = time.ParseInLocation("2006-01-02T15:04:05", timestr, loc)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
	return rsptime.Local()
	//return parseTime(timestr)

}

func GetFutureDate(startDate time.Time, years int, months int, days int, hours int, mins int, secs int) time.Time {
	fdate := startDate.AddDate(years, months, days)
	fdate = fdate.Add(time.Hour * time.Duration(hours))
	fdate = fdate.Add(time.Minute * time.Duration(mins))
	return fdate.Add(time.Second * time.Duration(secs))
}

// getErrorMessage returns the error message within
// the SOAP response or returns a generic error message
func GetErrorMessage(message string) string {
	if strings.Contains(message, "soap:Reason") {
		var start = strings.Index(message, "<soap:Reason>") + 13
		var end = strings.Index(message, "</soap:Reason>")
		var xmlmessage string = message[start:end]
		start = strings.Index(xmlmessage, ">") + 1
		end = strings.Index(xmlmessage, "</soap")
		return xmlmessage[start:end]
	}

	if strings.Contains(message, "faultstring") {
		var start = strings.Index(message, "<faultstring>") + 13
		var end = strings.Index(message, "</faultstring>")
		return message[start:end]
	}

	return "Soap error reason not found."
}

// containsError checks to see if the supplied
// message contains one of the two error tags
func ContainsError(message string) bool {
	return strings.Contains(message, "<soap:Fault>") || strings.Contains(message, "<faultstring>")
}

// getDocumentReturnList extracts the document
// return list from the SOAP response message
func GetDocumentReturnList(message string) string {
	if strings.Contains(message, "<return>") {
		var start = strings.Index(message, "<return>")
		var end = strings.Index(message, "</return>") + 9
		return message[start:end]
	}
	return message
}
func timeDuration(stime string, etime string) string {
	i := Trans{EnvVars: EnvState}
	return i.TimeDuration(stime, etime)
}
func (i *Trans) TimeDuration(stime string, etime string) string {
	if stime != "" {
		sTime := GetTimeFromString(stime)
		eTime := GetTimeFromString(etime)
		if eTime.After(sTime) {
			sTime = i.CalendarMode(sTime, eTime, true)
			duration := eTime.Sub(sTime)
			return PrettyPrintDuration(duration)
		}
	}
	return ""
}
func PrettyPrintDuration(duration time.Duration) string {
	rsp := ""
	secs := int(duration.Seconds())
	mins := secs / 60
	hrs := mins / 60
	days := hrs / 24
	hrs = hrs % 24
	mins = mins % 60
	if secs < 60 {
		return GetStringFromInt(secs) + " Secs"
	}
	if days == 1 {
		rsp = GetStringFromInt(days) + " Day " + GetStringFromInt(hrs) + " Hours " + GetStringFromInt(mins) + " Mins "
	} else {
		rsp = GetStringFromInt(days) + " Days " + GetStringFromInt(hrs) + " Hours " + GetStringFromInt(mins) + " Mins "
	}

	return strings.TrimPrefix(strings.TrimPrefix(rsp, "0 Days "), "0 Hours ")
}
func DT_Day() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Day())
}
func DT_Hour() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Hour())
}
func DT_Min() string {
	return fmt.Sprintf("%02d", time.Now().Local().Minute())
}
func DT_Sec() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Second())
}
func DT_MilliSec() int {
	return GetMilliseconds()
}
func DT_yyyy_MM_dd() string {
	return DT_Year() + "-" + DT_MM_dd()
}
func DT_yyyy_MM_dd_hh() string {
	return DT_yyyy_MM_dd() + " " + DT_Hour()
}
func DT_yyyy_MM_dd_hh_mm() string {
	return DT_yyyy_MM_dd_hh() + ":" + DT_Min()
}
func DT_yyyy_MM_dd_hh_mm_SS() string {
	return DT_yyyy_MM_dd_hh_mm() + ":" + DT_Sec()
}
func DT_yyyy_MM_dd_hh_mm_SS_sss() string {
	return DT_yyyy_MM_dd_hh_mm_SS() + "." + GetStringFromInt(DT_MilliSec())
}
func DT_yyyyMMddhhmmSSsss() string {
	return DT_Year() + DT_Month() + DT_Hour() + DT_Min() + DT_Sec() + strconv.Itoa(DT_MilliSec())
}
func DT_SQL() string {
	return DT_Year() + "-" + DT_Month() + "-" + DT_Day() + "T" + DT_Hour() + ":" + DT_Min() + ":" + DT_Sec()
}
func DT_SQL_Future_Year() string {
	t := time.Now().Local().AddDate(1, 0, 0)
	return GetStringFromInt(t.Local().Year()) + "-" + GetStringFromInt(int(t.Local().Month())) + "-" + GetStringFromInt(t.Local().Day()) + "T" + GetStringFromInt(t.Local().Hour()) + ":" + GetStringFromInt(t.Local().Minute()) + ":" + GetStringFromInt(t.Local().Second())
}
func DT_Zulu() string {
	return DT_SQL() + "." + GetStringFromInt(DT_MilliSec()) + "Z"
}
func DT_Zulu_Future(future int64) string {
	t := time.Now().Local().Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(future) + time.Second*time.Duration(0))
	return GetStringFromInt(t.Local().Year()) + "-" + GetStringFromInt(int(t.Local().Month())) + "-" + GetStringFromInt(t.Local().Day()) + "T" + GetStringFromInt(t.Local().Hour()) + ":" + GetStringFromInt(t.Local().Minute()) + ":" + GetStringFromInt(t.Local().Second()) + "." + GetStringFromInt(GetMilliseconds()) + "Z"
}
func DT_Zulu_Future_Year() string {
	t := time.Now().Local().AddDate(1, 0, 0)
	return GetStringFromInt(t.Local().Year()) + "-" + GetStringFromInt(int(t.Local().Month())) + "-" + GetStringFromInt(t.Local().Day()) + "T" + GetStringFromInt(t.Local().Hour()) + ":" + GetStringFromInt(t.Local().Minute()) + ":" + GetStringFromInt(t.Local().Second()) + "." + GetStringFromInt(GetMilliseconds()) + "Z"
}
func DT_Kitchen() string {
	return time.Now().Format(time.Kitchen)
}
func DT_Unix() string {
	return time.Now().Format(time.UnixDate)
}
func DT_ANSIC() string {
	return time.Now().Format(time.ANSIC)
}
func DT_Stamp() string {
	return time.Now().Format(time.Stamp)
}
func DT_Date() string {
	return time.Now().Format("Jan 2 2006")
}
func DT_Time() string {
	return time.Now().Format("15:04:05")
}
func DT_EPOCH() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintln(timestamp)
}
func GetMilliseconds() int {
	return GetIntFromString(Substr(GetStringFromInt(time.Now().Nanosecond()), 0, 3))
}

func Newdatetime() string {
	return DT_yyyy_MM_dd_hh_mm_SS()
}
func Newyearfuturezulu() string {
	return DT_Zulu_Future_Year()
}
func Newzulu() string {
	return DT_Zulu()
}
func New30mfutureyearzulu() string {
	return DT_Zulu_Future(30)
}
func DT_MM_dd() string {
	return DT_Month() + "-" + DT_Day()
}

func DT_Year() string {
	return fmt.Sprintf("%d",
		time.Now().Local().Year())
}
func GetIdIncrementSeed(len int) int {
	return GetIntFromString(Substr(GetStringFromInt(time.Now().Nanosecond()), 0, len))
}
func DT_Month() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Month())
}

// returns unique id in format '1.2.40.0.13.1.1.3542466645.20211021090059143.32643'
// idroot constant - 1.2.40.0.13.1.1.3542466645.
// + datetime	   - 20211021090059143.
// + 5 digit seed  - 32643
// The seed is incremented after each call to newid().
func Newid() string {
	id := SeedRoot + DT_yyyyMMddhhmmSSsss() + "." + GetStringFromInt(IdSeed)
	IdSeed = IdSeed + 1
	return id
}

func GetXMLNodeVal(message string, node string) string {
	if strings.Contains(message, node) {
		var nodeopen = "<" + node + ">"
		var nodeclose = "</" + node + ">"
		log.Println("Searching for value in : " + nodeopen + nodeclose)
		var start = strings.Index(message, nodeopen) + len(nodeopen)
		var end = strings.Index(message, nodeclose)
		m := message[start:end]
		log.Println("Returning value : " + m)
		return m
	}
	log.Println("Message does not contain Node : " + node)
	return ""
}
func GetFileBytes(f string) ([]byte, error) {
	file, err := os.Open(f)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	return byteValue, nil
}
func GetXmlReturnNode(message string) string {
	log.Println("Searching for <return> node in response message")
	if strings.Contains(message, "<return>") {
		var start = strings.Index(message, "<return>")
		var end = strings.Index(message, "</return>") + 9
		log.Println("Found Node <return>")
		return message[start:end]
	}
	log.Println("Node <return> Not found. Returning message")
	return message
}
func SplitFhirOid(oid string) string {
	if !strings.Contains(oid, ":") {
		return oid
	}

	splitoid := strings.Split(oid, ":")
	if len(splitoid) > 2 {
		return splitoid[2]
	}
	return oid
}
func SplitExpression(exp string) string {
	if !strings.Contains(exp, "^^") {
		return exp
	}
	str := strings.Split(exp, "^^")[0]
	return str
}

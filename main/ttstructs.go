package main

import (
	"encoding/xml"
	"net/http"
	"time"
)

type Quotes struct {
	Quote []Quote
}
type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
}
type DBConnection struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	DBTimeout     string
	DBReadTimeout string
	DB_URL        string
	DEBUG_MODE    bool
}
type Statics struct {
	Action       string   `json:"action"`
	LastInsertId int      `json:"lastinsertid"`
	Count        int      `json:"count"`
	Static       []Static `json:"static"`
}
type Static struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type Templates struct {
	Action       string     `json:"action"`
	LastInsertId int        `json:"lastinsertid"`
	Count        int        `json:"count"`
	Templates    []Template `json:"templates"`
}
type Template struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Template string `json:"template"`
	User     string `json:"user"`
}
type Subscription struct {
	Id         int    `json:"id"`
	Created    string `json:"created,omitempty"`
	BrokerRef  string `json:"brokerref,omitempty"`
	Pathway    string `json:"pathway,omitempty"`
	Topic      string `json:"topic,omitempty"`
	Expression string `json:"expression,omitempty"`
	Email      string `json:"email,omitempty"`
	NhsId      string `json:"nhsid,omitempty"`
	User       string `json:"user,omitempty"`
	Org        string `json:"org,omitempty"`
	Role       string `json:"role,omitempty"`
}
type Subscriptions struct {
	Action        string         `json:"action"`
	LastInsertId  int            `json:"lastinsertid"`
	Count         int            `json:"count"`
	Subscriptions []Subscription `json:"subscriptions"`
}
type Event struct {
	Id             int    `json:"id"`
	Creationtime   string `json:"creationtime,omitempty"`
	Eventtype      string `json:"eventtype,omitempty"`
	Docname        string `json:"docname,omitempty"`
	Classcode      string `json:"classcode,omitempty"`
	Confcode       string `json:"confcode,omitempty"`
	Formatcode     string `json:"formatcode,omitempty"`
	Facilitycode   string `json:"facilitycode,omitempty"`
	Practicecode   string `json:"practicecode,omitempty"`
	Expression     string `json:"expression,omitempty"`
	Authors        string `json:"authors,omitempty"`
	Xdsdocentryuid string `json:"xdsdocentryuId,omitempty"`
	Repositoryuid  string `json:"repositoryuniqueid,omitempty"`
	Nhs            string `json:"nhs,omitempty"`
	User           string `json:"user,omitempty"`
	Org            string `json:"org,omitempty"`
	Role           string `json:"role,omitempty"`
	Speciality     string `json:"speciality,omitempty"`
	Topic          string `json:"topic,omitempty"`
	Pathway        string `json:"pathway,omitempty"`
	Comments       string `json:"comments,omitempty"`
	Version        int    `json:"ver"`
	Taskid         int    `json:"taskid"`
	Brokerref      string `json:"brokerref,omitempty"`
}
type Events struct {
	Action       string  `json:"action"`
	LastInsertId int     `json:"lastinsertid"`
	Count        int     `json:"count"`
	Events       []Event `json:"events"`
}
type Workflow struct {
	Id        int    `json:"id"`
	Created   string `json:"created,omitempty"`
	Pathway   string `json:"pathway,omitempty"`
	NHSId     string `json:"nhsid,omitempty"`
	XDW_Key   string `json:"xdw_key,omitempty"`
	XDW_UID   string `json:"xdw_uid,omitempty"`
	XDW_Doc   string `json:"xdw_doc,omitempty"`
	XDW_Def   string `json:"xdw_def,omitempty"`
	Version   int    `json:"version"`
	Published bool   `json:"published"`
	Status    string `json:"status,omitempty"`
}
type Workflows struct {
	Action       string     `json:"action"`
	LastInsertId int        `json:"lastinsertid"`
	Count        int        `json:"count"`
	Workflows    []Workflow `json:"workflows"`
}
type WorkflowStates struct {
	Action        string          `json:"action"`
	LastInsertId  int             `json:"lastinsertid"`
	Count         int             `json:"count"`
	Workflowstate []Workflowstate `json:"workflowstate"`
}
type Workflowstate struct {
	WorkflowId      int    `json:"workflowid"`
	Pathway         string `json:"pathway"`
	NHSId           string `json:"nhsid"`
	Version         int    `json:"version"`
	Published       bool   `json:"published"`
	Created         string `json:"created"`
	CreatedBy       string `json:"createdby"`
	Status          string `json:"status"`
	CompleteBy      string `json:"completeby,omitempty"`
	WDCompleteBy    string `json:"wdcompleteby,omitempty"`
	LastUpdate      string `json:"lastupdate"`
	Owner           string `json:"owner,omitempty"`
	Overdue         string `json:"overdue"`
	Escalated       string `json:"escalated"`
	TargetMet       string `json:"targetmet"`
	InProgress      string `json:"inprogress"`
	Duration        string `json:"duration"`
	WDDuration      string `json:"wdduration"`
	TimeRemaining   string `json:"timeremaining"`
	WDTimeRemaining string `json:"wdtimeremaining"`
}
type XDWS struct {
	Action       string `json:"action"`
	LastInsertId int    `json:"lastinsertid"`
	Count        int    `json:"count"`
	XDW          []XDW  `json:"xdws"`
}
type XDW struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsXDSMeta bool   `json:"isxdsmeta"`
	XDW       string `json:"xdw"`
}
type IdMaps struct {
	Action       string
	LastInsertId int
	Where        string
	Value        string
	Cnt          int
	LidMap       []IdMap
}
type IdMap struct {
	Id   int    `json:"id"`
	User string `json:"user"`
	Lid  string `json:"lid"`
	Mid  string `json:"mid"`
}
type Pwys struct {
	Pwy []Pwy
}
type Pwy struct {
	Text  string
	Value string
}
type SMTPEvent struct {
	Body     string
	From     string
	To       string
	Server   string
	Port     string
	Password string
}
type HTTPRequest struct {
	Method     string
	Header     http.Header
	URL        string
	Request    []byte
	Response   []byte
	StatusCode int
	Timeout    int
}
type Patients struct {
	PIX          PIXmResponse
	PDS_Retrieve PDSRetrieveResponse
	PDS_Search   PDSSearchResponse
	CGL          CGLUserResponse
}
type XDWCreator struct {
	User       string
	Org        string
	Role       string
	Pathway    string
	NHS        string
	Comments   string
	Definition WorkflowDefinition
	Meta       WorkflowMeta
	Document   WorkflowDocument
}
type DBVars struct {
	DB_USER     string
	DB_NAME     string
	DB_PORT     string
	DB_HOST     string
	DB_PASSWORD string
}
type Trans struct {
	Req                 any                 `json:"Req"`
	HTTP                HTTP                `json:"HTTP"`
	EnvVars             *EnvVars            `json:"EnvVars"`
	DBVars              *DBVars             `json:"omitempty"`
	Query               QueryVars           `json:"Query"`
	XDWState            XDWState            `json:"State,omitempty"`
	IdMaps              IdMaps              `json:"Idmaps.omitempty"`
	Subscriptions       Subscriptions       `json:"Subscriptions,omitempty"`
	Subscription        Subscription        `json:"Subscription,omitempty"`
	PIXmResponse        PIXmResponse        `json:"PIXmResponse,omitempty"`
	PDSResponse         string              `json:"PDSResponse,omitempty"`
	PDSSearchResponse   PDSSearchResponse   `json:"PDSSearchResponse,omitempty"`
	PDSRetrieveResponse PDSRetrieveResponse `json:"PDSRetrieveResponse,omitempty"`
	CGLResponse         CGLUserResponse     `json:"CGLResponse,omitempty"`
	Error               error               `json:"Error,omitempty"`
}
type EnvVars struct {
	DEBUG_MODE              bool   `json:"DEBUG_MODE"`
	DEBUG_DB                bool   `json:"DEBUG_DB"`
	DEBUG_DB_ERROR          bool   `json:"DEBUG_DB_ERROR"`
	DSUB_BROKER_URL         string `json:"DSUB_BROKER_URL,omitempty"`
	DSUB_CONSUMER_URL       string `json:"DSUB_CONSUMER_URL,omitempty"`
	LOGO_FILE               string `json:"LOGO_FILE,omitempty"`
	SERVER_PORT             string `json:"SERVER_PORT,omitempty"`
	SERVER_URL              string `json:"SERVER_URL,omitempty"`
	SERVER_NAME             string `json:"SERVER_NAME,omitempty"`
	SCR_URL                 string `json:"SCR_URL,omitempty"`
	REG_OID                 string `json:"REG_OID,omitempty"`
	PIXM_SERVER_URL         string `json:"PIXM_SERVER_URL,omitempty"`
	PDS_SERVER_URL          string `json:"PDS_SERVER_URL,omitempty"`
	CGL_SERVER_URL          string `json:"CGL_SERVER_URL,omitempty"`
	CGL_SERVER_X_API_KEY    string `json:"CGL_SERVER_X_API_KEY"`
	CGL_SERVER_X_API_SECRET string `json:"CGL_SERVER_X_API_SECRET"`
	S3_PUBLISH_FILES        string `json:"S3_PUBLISH_FILES"`
	SMTP_SERVER             string `json:"SMTP_SERVER"`
	SMTP_USER               string `json:"SMTP_USER"`
	SMTP_PORT               string `json:"SMTP_PORT"`
	SMTP_SUBJECT            string `json:"SMTP_SUBJECT"`
	SMTP_PASSWORD           string `json:"SMTP_PASSWORD"`
	PERSIST_TEMPLATES       bool   `json:"PERSIST_TEMPLATES"`
	PERSIST_DEFINITIONS     bool   `json:"PERSIST_DEFINITIONS"`
	CALENDAR_MODE           string `json:"CALENDAR_MODE"`
	TIME_LOCATION           string `json:"TIME_LOCATION"`
	TIME_LOCALE             string `json:"TIME_LOCALE"`
	START_OF_DAY_HOUR       string `json:"START_OF_DAY_HOUR"`
	END_OF_DAY_HOUR         string `json:"END_OF_DAY_HOUR"`
}
type QueryVars struct {
	Act            string `json:"Act,omitempty"`
	Action         string `json:"Action,omitempty"`
	Operation      string `json:"Operation,omitempty"`
	User           string `json:"User,omitempty"`
	Org            string `json:"Org,omitempty"`
	Role           string `json:"Role,omitempty"`
	Email          string `json:"Email,omitempty"`
	Nhs            string `json:"Nhs,omitempty"`
	Pid            string `json:"Pid,omitempty"`
	Pidoid         string `json:"Pidoid,omitempty"`
	Pathway        string `json:"Pathway,omitempty"`
	Topic          string `json:"Topic,omitempty"`
	Expression     string `json:"Expression,omitempty"`
	Vers           string `json:"Vers,omitempty"`
	Status         string `json:"Status,omitempty"`
	Id             string `json:"Id,omitempty"`
	Taskid         string `json:"Taskid,omitempty"`
	Format         string `json:"Format,omitempty"`
	Comments       string `json:"Comments,omitempty"`
	Formatcode     string `json:"Formatcode,omitempty"`
	Brokerref      string `json:"Brokerref,omitempty"`
	Eventtype      string `json:"Eventype,omitempty"`
	Template       string `json:"Template,omitempty"`
	Name           string `json:"Name,omitempty"`
	Lid            string `json:"Lid,omitempty"`
	Mid            string `json:"Mid,omitempty"`
	Repositoryuid  string `json:"Repositoryuid,omitempty"`
	Xdsdocentryuid string `json:"Xdsdocentryuid,omitempty"`
}
type HTTP struct {
	Host           string `json:"Host,omitempty"`
	Path           string `json:"Path,omitempty"`
	Auth           string `json:"Auth,omitempty"`
	Token          string `json:"Token,omitempty"`
	Method         string `json:"Method,omitempty"`
	SourceIP       string `json:"SourceIP,omitempty"`
	ReqURL         string `json:"ReqURL,omitempty"`
	ReqContentType string `json:"ReqContentType"`
	RequestBody    string `json:"RequestBody,omitempty"`
	SoapAction     string `json:"SoapAction,omitempty"`
	RspContentType string `json:"RspContentType"`
	ResponseBody   string `json:"ResponseBody"`
	StatusCode     int    `json:"StatusCode"`
	IsBase64       bool   `json:"IsBase64"`
	Timeout        int    `json:"Timeout"`
}
type XDWState struct {
	Workflows          Workflows          `json:"Workflows,omitempty"`
	OpenWorkflows      Workflows          `json:"OpenWorkflows,omitempty"`
	MetWorkflows       Workflows          `json:"SoapAction,omitempty"`
	OverdueWorkflows   Workflows          `json:"OverdueWorkflows,omitempty"`
	EscalatedWorkflows Workflows          `json:"EscalatedWorkflows,omitempty"`
	ClosedWorkflows    Workflows          `json:"ClosedWorkflows,omitempty"`
	Dashboard          Dashboard          `json:"Dashboard,omitempty"`
	Events             Events             `json:"Events,omitempty"`
	Expressions        []string           `json:"Expressions,omitempty"`
	WorkflowStates     []Workflowstate    `json:"WorkflowStates,omitempty"`
	Definition         WorkflowDefinition `json:"Definition,omitempty"`
	Meta               WorkflowMeta       `json:"Meta,omitempty"`
	WorkflowDocument   WorkflowDocument   `xml:"XDW.WorkflowDocument" json:"WorkflowDocument"`
}
type Dashboard struct {
	InProgress   int
	Escalated    int
	TargetMet    int
	TargetMissed int
	Complete     int
	Total        int
}
type WorkflowMeta struct {
	Id                       string `json:"id"`
	Repositoryuniqueid       string `json:"repositoryuniqueid"`
	Docname                  string `json:"docname"`
	ClasscodeValue           string `json:"classcodevalue"`
	PracticesettingcodeValue string `json:"practicesettingcodevalue"`
	ConfcodeValue            string `json:"confcodevalue"`
	FacilitycodeValue        string `json:"facilitycodevalue"`
	FormatcodeValue          string `json:"formatcodevalue"`
}
type WorkflowDefinition struct {
	Ref                 string `json:"ref"`
	Name                string `json:"name"`
	Confidentialitycode string `json:"confidentialitycode"`
	StartByTime         string `json:"startbytime"`
	CompleteByTime      string `json:"completebytime"`
	ExpirationTime      string `json:"expirationtime"`
	PotentialOwners     []struct {
		OrganizationalEntity struct {
			User string `json:"user"`
		} `json:"organizationalEntity"`
	} `json:"potentialOwners,omitempty"`
	CompletionBehavior []struct {
		Completion struct {
			Condition string `json:"condition"`
		} `json:"completion"`
	} `json:"completionBehavior"`
	Tasks []struct {
		ID              string `json:"id"`
		Tasktype        string `json:"tasktype"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		ActualOwner     string `json:"actualowner"`
		ExpirationTime  string `json:"expirationtime,omitempty"`
		StartByTime     string `json:"startbytime,omitempty"`
		CompleteByTime  string `json:"completebytime"`
		IsSkipable      bool   `json:"isskipable,omitempty"`
		PotentialOwners []struct {
			OrganizationalEntity struct {
				User string `json:"user"`
			} `json:"organizationalEntity"`
		} `json:"potentialOwners,omitempty"`
		CompletionBehavior []struct {
			Completion struct {
				Condition string `json:"condition"`
			} `json:"completion"`
		} `json:"completionBehavior"`
		Input []struct {
			Name        string `json:"name"`
			Contenttype string `json:"contenttype"`
			AccessType  string `json:"accesstype"`
		} `json:"input,omitempty"`
		Output []struct {
			Name        string `json:"name"`
			Contenttype string `json:"contenttype"`
			AccessType  string `json:"accesstype"`
		} `json:"output,omitempty"`
	} `json:"tasks"`
}
type WorkflowDocument struct {
	XMLName                        xml.Name              `xml:"XDW.WorkflowDocument" json:"XMLName"`
	Hl7                            string                `xml:"hl7,attr" json:"Hl7"`
	WsHt                           string                `xml:"ws-ht,attr" json:"WsHt"`
	Xdw                            string                `xml:"xdw,attr" json:"Xdw"`
	Xsi                            string                `xml:"xsi,attr" json:"Xsi"`
	SchemaLocation                 string                `xml:"schemaLocation,attr" json:"SchemaLocation"`
	ID                             ID                    `xml:"id" json:"ID"`
	EffectiveTime                  EffectiveTime         `xml:"effectiveTime" json:"EffectiveTime"`
	ConfidentialityCode            ConfidentialityCode   `xml:"confidentialityCode" json:"ConfidentialityCode"`
	Patient                        ID                    `xml:"patient" json:"Patient"`
	Author                         Author                `xml:"author" json:"Author"`
	WorkflowInstanceId             string                `xml:"workflowInstanceId" json:"WorkflowInstanceId"`
	WorkflowDocumentSequenceNumber string                `xml:"workflowDocumentSequenceNumber" json:"WorkflowDocumentSequenceNumber"`
	WorkflowStatus                 string                `xml:"workflowStatus" json:"WorkflowStatus"`
	WorkflowStatusHistory          WorkflowStatusHistory `xml:"workflowStatusHistory" json:"WorkflowStatusHistory"`
	WorkflowDefinitionReference    string                `xml:"workflowDefinitionReference" json:"WorkflowDefinitionReference"`
	TaskList                       TaskList              `xml:"TaskList" json:"TaskList"`
}
type ConfidentialityCode struct {
	Code string `xml:"code,attr" json:"Code"`
}
type EffectiveTime struct {
	Value string `xml:"value,attr" json:"Value"`
}
type Author struct {
	AssignedAuthor AssignedAuthor `xml:"assignedAuthor" json:"AssignedAuthor"`
}
type AssignedAuthor struct {
	ID             ID             `xml:"id" json:"ID"`
	AssignedPerson AssignedPerson `xml:"assignedPerson" json:"AssignedPerson"`
}
type ID struct {
	Root                   string `xml:"root,attr" json:"Root"`
	Extension              string `xml:"extension,attr" json:"Extension"`
	AssigningAuthorityName string `xml:"assigningAuthorityName,attr" json:"AssigningAuthorityName"`
}
type AssignedPerson struct {
	Name Name `xml:"name" json:"Name"`
}
type Name struct {
	Family string `xml:"family" json:"Family"`
	Prefix string `xml:"prefix" json:"Prefix"`
}
type WorkflowStatusHistory struct {
	DocumentEvent []DocumentEvent `xml:"documentEvent" json:"DocumentEvent"`
}
type TaskList struct {
	XDWTask []XDWTask `xml:"XDWTask" json:"Task"`
}
type XDWTask struct {
	TaskData         TaskData         `xml:"taskData" json:"TaskData"`
	TaskEventHistory TaskEventHistory `xml:"taskEventHistory" json:"TaskEventHistory"`
}
type TaskData struct {
	TaskDetails TaskDetails `xml:"taskDetails" json:"TaskDetails"`
	Description string      `xml:"description" json:"Description"`
	Input       []Input     `xml:"input" json:"Input"`
	Output      []Output    `xml:"output" json:"Output"`
}
type TaskDetails struct {
	ID                    string `xml:"id" json:"ID"`
	TaskType              string `xml:"taskType" json:"TaskType"`
	Name                  string `xml:"name" json:"Name"`
	Status                string `xml:"status" json:"Status"`
	ActualOwner           string `xml:"actualOwner" json:"ActualOwner"`
	CreatedTime           string `xml:"createdTime" json:"CreatedTime"`
	CreatedBy             string `xml:"createdBy" json:"CreatedBy"`
	ActivationTime        string `xml:"activationTime" json:"ActivationTime"`
	LastModifiedTime      string `xml:"lastModifiedTime" json:"LastModifiedTime"`
	RenderingMethodExists string `xml:"renderingMethodExists" json:"RenderingMethodExists"`
}
type TaskEventHistory struct {
	TaskEvent []TaskEvent `xml:"taskEvent" json:"TaskEvent"`
}
type AttachmentInfo struct {
	Identifier      string `xml:"identifier" json:"identifier"`
	Name            string `xml:"name" json:"name"`
	AccessType      string `xml:"accessType" json:"accesstype"`
	ContentType     string `xml:"contentType" json:"contenttype"`
	ContentCategory string `xml:"contentCategory" json:"contentcategory"`
	AttachedTime    string `xml:"attachedTime" json:"attachedtime"`
	AttachedBy      string `xml:"attachedBy" json:"attachedby"`
	HomeCommunityId string `xml:"homeCommunityId" json:"homecommunityid"`
}
type Part struct {
	Name           string         `xml:"name,attr" json:"Name"`
	AttachmentInfo AttachmentInfo `xml:"attachmentInfo" json:"AttachmentInfo"`
}
type Output struct {
	Part Part `xml:"part" json:"Part"`
}
type Input struct {
	Part Part `xml:"part" json:"Part"`
}
type DocumentEvent struct {
	EventTime           string `xml:"eventTime" json:"EventTime"`
	EventType           string `xml:"eventType" json:"EventType"`
	TaskEventIdentifier string `xml:"taskEventIdentifier" json:"TaskEventIdentifier"`
	Author              string `xml:"author" json:"Author"`
	PreviousStatus      string `xml:"previousStatus" json:"PreviousStatus"`
	ActualStatus        string `xml:"actualStatus" json:"ActualStatus"`
}
type TaskEvent struct {
	ID         string `xml:"id" json:"ID"`
	EventTime  string `xml:"eventTime" json:"EventTime"`
	Identifier string `xml:"identifier" json:"Identifier"`
	EventType  string `xml:"eventType" json:"EventType"`
	Status     string `xml:"status" json:"Status"`
}
type DSUBSubscribe struct {
	BrokerURL   string
	ConsumerURL string
	Topic       string
	Expression  string
}
type DSUBSubscribeResponse struct {
	XMLName        xml.Name `xml:"Envelope"`
	Text           string   `xml:",chardata"`
	S              string   `xml:"s,attr"`
	A              string   `xml:"a,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Wsnt           string   `xml:"wsnt,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Header         struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text              string `xml:",chardata"`
		SubscribeResponse struct {
			Text                  string `xml:",chardata"`
			SubscriptionReference struct {
				Text    string `xml:",chardata"`
				Address string `xml:"Address"`
			} `xml:"SubscriptionReference"`
		} `xml:"SubscribeResponse"`
	} `xml:"Body"`
}
type DSUBNotifyMessage struct {
	XMLName             xml.Name `xml:"Notify"`
	Text                string   `xml:",chardata"`
	Xmlns               string   `xml:"xmlns,attr"`
	Xsd                 string   `xml:"xsd,attr"`
	Xsi                 string   `xml:"xsi,attr"`
	NotificationMessage struct {
		Text                  string `xml:",chardata"`
		SubscriptionReference struct {
			Text    string `xml:",chardata"`
			Address struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
			} `xml:"Address"`
		} `xml:"SubscriptionReference"`
		Topic struct {
			Text    string `xml:",chardata"`
			Dialect string `xml:"Dialect,attr"`
		} `xml:"Topic"`
		ProducerReference struct {
			Text    string `xml:",chardata"`
			Address struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
			} `xml:"Address"`
		} `xml:"ProducerReference"`
		Message struct {
			Text                 string `xml:",chardata"`
			SubmitObjectsRequest struct {
				Text               string `xml:",chardata"`
				Lcm                string `xml:"lcm,attr"`
				RegistryObjectList struct {
					Text            string `xml:",chardata"`
					Rim             string `xml:"rim,attr"`
					ExtrinsicObject struct {
						Text       string `xml:",chardata"`
						A          string `xml:"a,attr"`
						ID         string `xml:"id,attr"`
						MimeType   string `xml:"mimeType,attr"`
						ObjectType string `xml:"objectType,attr"`
						Slot       []struct {
							Text      string `xml:",chardata"`
							Name      string `xml:"name,attr"`
							ValueList struct {
								Text  string   `xml:",chardata"`
								Value []string `xml:"Value"`
							} `xml:"ValueList"`
						} `xml:"Slot"`
						Name struct {
							Text            string `xml:",chardata"`
							LocalizedString struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value,attr"`
							} `xml:"LocalizedString"`
						} `xml:"Name"`
						Description    string `xml:"Description"`
						Classification []struct {
							Text                 string `xml:",chardata"`
							ClassificationScheme string `xml:"classificationScheme,attr"`
							ClassifiedObject     string `xml:"classifiedObject,attr"`
							ID                   string `xml:"id,attr"`
							NodeRepresentation   string `xml:"nodeRepresentation,attr"`
							ObjectType           string `xml:"objectType,attr"`
							Slot                 []struct {
								Text      string `xml:",chardata"`
								Name      string `xml:"name,attr"`
								ValueList struct {
									Text  string   `xml:",chardata"`
									Value []string `xml:"Value"`
								} `xml:"ValueList"`
							} `xml:"Slot"`
							Name struct {
								Text            string `xml:",chardata"`
								LocalizedString struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"LocalizedString"`
							} `xml:"Name"`
						} `xml:"Classification"`
						ExternalIdentifier []struct {
							Text                 string `xml:",chardata"`
							ID                   string `xml:"id,attr"`
							IdentificationScheme string `xml:"identificationScheme,attr"`
							ObjectType           string `xml:"objectType,attr"`
							RegistryObject       string `xml:"registryObject,attr"`
							Value                string `xml:"value,attr"`
							Name                 struct {
								Text            string `xml:",chardata"`
								LocalizedString struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"LocalizedString"`
							} `xml:"Name"`
						} `xml:"ExternalIdentifier"`
					} `xml:"ExtrinsicObject"`
				} `xml:"RegistryObjectList"`
			} `xml:"SubmitObjectsRequest"`
		} `xml:"Message"`
	} `xml:"NotificationMessage"`
}
type DelphiResponse struct {
	Data struct {
		LocalIdentifier int    `json:"LocalIdentifier,omitempty"`
		Status          string `json:"Status,omitempty"`
		Title           string `json:"Title,omitempty"`
		Forename        string `json:"Forename,omitempty"`
		Surname         string `json:"Surname,omitempty"`
		GenderAtBirth   string `json:"GenderAtBirth,omitempty"`
		DateOfBirth     string `json:"DateOfBirth,omitempty"`
		Address         struct {
			LocalIdentifier int    `json:"LocalIdentifier,omitempty"`
			AddressLine1    string `json:"AddressLine1,omitempty"`
			AddressLine2    string `json:"AddressLine2,omitempty"`
			AddressLine3    string `json:"AddressLine3,omitempty"`
			AddressLine4    string `json:"AddressLine4,omitempty"`
			PostCode1       string `json:"PostCode1,omitempty"`
			PostCode2       string `json:"PostCode2,omitempty"`
		} `json:"Address,omitempty"`
		Keyworker               string `json:"Keyworker,omitempty"`
		LastAttendedAppointment string `json:"LastAttendedAppointment,omitempty"`
		DrugScreening           []any  `json:"DrugScreening,omitempty"`
		Prescriptions           []any  `json:"Prescriptions,omitempty"`
		Risks                   []any  `json:"Risks,omitempty"`
		Careplans               []struct {
			LocalIdentifier          int    `json:"LocalIdentifier,omitempty"`
			AlcohoUse                bool   `json:"AlcohoUse,omitempty"`
			DrugUse                  bool   `json:"DrugUse,omitempty"`
			EffectsOfAlcoholAndDrugs bool   `json:"EffectsOfAlcoholAndDrugs,omitempty"`
			PreventingRelapse        bool   `json:"PreventingRelapse,omitempty"`
			PreventingOverdose       bool   `json:"PreventingOverdose,omitempty"`
			PersonalCare             bool   `json:"PersonalCare,omitempty"`
			FindingThingsIEnjoy      bool   `json:"FindingThingsIEnjoy,omitempty"`
			ManagingMoney            bool   `json:"ManagingMoney,omitempty"`
			SupportForMyChildren     bool   `json:"SupportForMyChildren,omitempty"`
			EducationOrTraining      bool   `json:"EducationOrTraining,omitempty"`
			Other                    bool   `json:"Other,omitempty"`
			AlcoholDrugUse           bool   `json:"AlcoholDrugUse,omitempty"`
			ManagingCravings         bool   `json:"ManagingCravings,omitempty"`
			MentalEmotionalHealth    bool   `json:"MentalEmotionalHealth,omitempty"`
			AccommodationHousing     bool   `json:"AccommodationHousing,omitempty"`
			LegalProblems            bool   `json:"LegalProblems,omitempty"`
			ParentingHelpSupport     bool   `json:"ParentingHelpSupport,omitempty"`
			PhyscialHealth           bool   `json:"PhyscialHealth,omitempty"`
			ImmediateProblem         string `json:"ImmediateProblem,omitempty"`
			LongTermGoal             string `json:"LongTermGoal,omitempty"`
			StepsToAchievingGoal     string `json:"StepsToAchievingGoal,omitempty"`
			HowDidItGo               string `json:"HowDidItGo,omitempty"`
			NextStepForGoal          string `json:"NextStepForGoal,omitempty"`
			CommunityDetox           bool   `json:"CommunityDetox,omitempty"`
			InpatientDetox           bool   `json:"InpatientDetox,omitempty"`
			OverdoseInformation      bool   `json:"OverdoseInformation,omitempty"`
			NutritionalAdvice        bool   `json:"NutritionalAdvice,omitempty"`
			HepCScreening            bool   `json:"HepCScreening,omitempty"`
			HepAAndBVaccination      bool   `json:"HepAAndBVaccination,omitempty"`
			GroupWork                bool   `json:"GroupWork,omitempty"`
			OneToOneSupport          bool   `json:"OneToOneSupport,omitempty"`
			SupportWorker            bool   `json:"SupportWorker,omitempty"`
			PrescribedMedication     bool   `json:"PrescribedMedication,omitempty"`
			Stabilisation            bool   `json:"Stabilisation,omitempty"`
			MedicalReview            bool   `json:"MedicalReview,omitempty"`
			OtherClinical            bool   `json:"OtherClinical,omitempty"`
			CarePlanGivenToClient    string `json:"CarePlanGivenToClient,omitempty"`
			CareplanStartDate        string `json:"CareplanStartDate,omitempty"`
			CarePlanReviewDate       string `json:"CarePlanReviewDate,omitempty"`
		} `json:"Careplans,omitempty"`
		Discharge any `json:"Discharge,omitempty"`
	} `json:"Data,omitempty"`
}
type CGLUserResponse struct {
	Data struct {
		Client struct {
			BasicDetails struct {
				Address struct {
					AddressLine1 string `json:"addressLine1,omitempty"`
					AddressLine2 string `json:"addressLine2,omitempty"`
					AddressLine3 string `json:"addressLine3,omitempty"`
					AddressLine4 string `json:"addressLine4,omitempty"`
					AddressLine5 string `json:"addressLine5,omitempty"`
					PostCode     string `json:"postCode,omitempty"`
				} `json:"address,omitempty"`
				BirthDate                    string `json:"birthDate,omitempty"`
				Disability                   string `json:"disability,omitempty"`
				LastEngagementByCGLDate      string `json:"lastEngagementByCGLDate,omitempty"`
				LastFaceToFaceEngagementDate string `json:"lastFaceToFaceEngagementDate,omitempty"`
				LocalIdentifier              int    `json:"localIdentifier,omitempty"`
				Name                         struct {
					Family string `json:"family,omitempty"`
					Given  string `json:"given,omitempty"`
				} `json:"name,omitempty"`
				NextCGLAppointmentDate string `json:"nextCGLAppointmentDate,omitempty"`
				NhsNumber              string `json:"nhsNumber,omitempty"`
				SexAtBirth             string `json:"sexAtBirth,omitempty"`
			} `json:"basicDetails,omitempty"`
			BbvInformation struct {
				BbvTested        string `json:"bbvTested,omitempty"`
				HepCLastTestDate string `json:"hepCLastTestDate,omitempty"`
				HepCResult       string `json:"hepCResult,omitempty"`
				HivPositive      string `json:"hivPositive,omitempty"`
			} `json:"bbvInformation,omitempty"`
			DrugTestResults struct {
				DrugTestDate          string `json:"drugTestDate,omitempty"`
				DrugTestSample        string `json:"drugTestSample,omitempty"`
				DrugTestStatus        string `json:"drugTestStatus,omitempty"`
				InstantOrConfirmation string `json:"instantOrConfirmation,omitempty"`
				Results               struct {
					Amphetamine     string `json:"amphetamine,omitempty"`
					Benzodiazepine  string `json:"benzodiazepine,omitempty"`
					Buprenorphine   string `json:"buprenorphine,omitempty"`
					Cannabis        string `json:"cannabis,omitempty"`
					Cocaine         string `json:"cocaine,omitempty"`
					Eddp            string `json:"eddp,omitempty"`
					Fentanyl        string `json:"fentanyl,omitempty"`
					Ketamine        string `json:"ketamine,omitempty"`
					Methadone       string `json:"methadone,omitempty"`
					Methamphetamine string `json:"methamphetamine,omitempty"`
					Morphine        string `json:"morphine,omitempty"`
					Opiates         string `json:"opiates,omitempty"`
					SixMam          string `json:"sixMam,omitempty"`
					Tramadol        string `json:"tramadol,omitempty"`
				} `json:"results,omitempty"`
			} `json:"drugTestResults,omitempty"`
			PrescribingInformation []string `json:"prescribingInformation,omitempty"`
			RiskInformation        struct {
				LastSelfReportedDate string `json:"lastSelfReportedDate,omitempty"`
				MentalHealthDomain   struct {
					AttemptedSuicide                            string `json:"attemptedSuicide,omitempty"`
					CurrentOrPreviousSelfHarm                   string `json:"currentOrPreviousSelfHarm,omitempty"`
					DiagnosedMentalHealthCondition              string `json:"diagnosedMentalHealthCondition,omitempty"`
					FrequentLifeThreateningSelfHarm             string `json:"frequentLifeThreateningSelfHarm,omitempty"`
					Hallucinations                              string `json:"hallucinations,omitempty"`
					HospitalAdmissionsForMentalHealth           string `json:"hospitalAdmissionsForMentalHealth,omitempty"`
					NoIdentifiedRisk                            string `json:"noIdentifiedRisk,omitempty"`
					NotEngagingWithSupport                      string `json:"notEngagingWithSupport,omitempty"`
					NotTakingPrescribedMedicationAsInstructed   string `json:"notTakingPrescribedMedicationAsInstructed,omitempty"`
					PsychiatricOrPreviousCrisisTeamIntervention string `json:"psychiatricOrPreviousCrisisTeamIntervention,omitempty"`
					Psychosis                                   string `json:"psychosis,omitempty"`
					SelfReportedMentalHealthConcerns            string `json:"selfReportedMentalHealthConcerns,omitempty"`
					ThoughtsOfSuicideOrSelfHarm                 string `json:"thoughtsOfSuicideOrSelfHarm,omitempty"`
				} `json:"mentalHealthDomain,omitempty"`
				RiskOfHarmToSelfDomain struct {
					AssessedAsNotHavingMentalCapacity  string `json:"assessedAsNotHavingMentalCapacity,omitempty"`
					BeliefTheyAreWorthless             string `json:"beliefTheyAreWorthless,omitempty"`
					Hoarding                           string `json:"hoarding,omitempty"`
					LearningDisability                 string `json:"learningDisability,omitempty"`
					MeetsSafeguardingAdultsThreshold   string `json:"meetsSafeguardingAdultsThreshold,omitempty"`
					NoIdentifiedRisk                   string `json:"noIdentifiedRisk,omitempty"`
					OngoingConcernsRelatingToOwnSafety string `json:"ongoingConcernsRelatingToOwnSafety,omitempty"`
					ProblemsMaintainingPersonalHygiene string `json:"problemsMaintainingPersonalHygiene,omitempty"`
					ProblemsMeetingNutritionalNeeds    string `json:"problemsMeetingNutritionalNeeds,omitempty"`
					RequiresIndependentAdvocacy        string `json:"requiresIndependentAdvocacy,omitempty"`
					SelfNeglect                        string `json:"selfNeglect,omitempty"`
				} `json:"riskOfHarmToSelfDomain,omitempty"`
				SocialDomain struct {
					FinancialProblems         string `json:"financialProblems,omitempty"`
					HomelessRoughSleepingNFA  string `json:"homelessRoughSleepingNFA,omitempty"`
					HousingAtRisk             string `json:"housingAtRisk,omitempty"`
					NoIdentifiedRisk          string `json:"noIdentifiedRisk,omitempty"`
					SociallyIsolatedNoSupport string `json:"sociallyIsolatedNoSupport,omitempty"`
				} `json:"socialDomain,omitempty"`
				SubstanceMisuseDomain struct {
					ConfusionOrDisorientation string `json:"ConfusionOrDisorientation,omitempty"`
					AdmissionToAandE          string `json:"admissionToAandE,omitempty"`
					BlackoutOrSeizures        string `json:"blackoutOrSeizures,omitempty"`
					ConcurrentUse             string `json:"concurrentUse,omitempty"`
					HigherRiskDrinking        string `json:"higherRiskDrinking,omitempty"`
					InjectedByOthers          string `json:"injectedByOthers,omitempty"`
					Injecting                 string `json:"injecting,omitempty"`
					InjectingInNeckOrGroin    string `json:"injectingInNeckOrGroin,omitempty"`
					NoIdentifiedRisk          string `json:"noIdentifiedRisk,omitempty"`
					PolyDrugUse               string `json:"polyDrugUse,omitempty"`
					PreviousOverDose          string `json:"previousOverDose,omitempty"`
					RecentPrisonRelease       string `json:"recentPrisonRelease,omitempty"`
					ReducedTolerance          string `json:"reducedTolerance,omitempty"`
					SharingWorks              string `json:"sharingWorks,omitempty"`
					Speedballing              string `json:"speedballing,omitempty"`
					UsingOnTop                string `json:"usingOnTop,omitempty"`
				} `json:"substanceMisuseDomain,omitempty"`
			} `json:"riskInformation,omitempty"`
			SafeguardingInformation struct {
				LastReviewDate     string `json:"lastReviewDate,omitempty"`
				RiskHarmFromOthers string `json:"riskHarmFromOthers,omitempty"`
				RiskToAdults       string `json:"riskToAdults,omitempty"`
				RiskToChildrenOrYP string `json:"riskToChildrenOrYP,omitempty"`
				RiskToSelf         string `json:"riskToSelf,omitempty"`
			} `json:"safeguardingInformation,omitempty"`
		} `json:"client,omitempty"`
		KeyWorker struct {
			LocalIdentifier int `json:"localIdentifier,omitempty"`
			Name            struct {
				Family string `json:"family,omitempty"`
				Given  string `json:"given,omitempty"`
			} `json:"name"`
			Telecom string `json:"telecom,omitempty"`
		} `json:"keyWorker,omitempty"`
	} `json:"data,omitempty"`
}
type PDQv3Response struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Env     string   `xml:"env,attr"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"Action"`
		MessageID struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"MessageID"`
		RelatesTo struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"RelatesTo"`
		To struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"To"`
	} `xml:"Header"`
	Body struct {
		Text             string `xml:",chardata"`
		PRPAIN201306UV02 struct {
			Text       string `xml:",chardata"`
			Xmlns      string `xml:"xmlns,attr"`
			ITSVersion string `xml:"ITSVersion,attr"`
			ID         struct {
				Text      string `xml:",chardata"`
				Extension string `xml:"extension,attr"`
				Root      string `xml:"root,attr"`
			} `xml:"id"`
			CreationTime struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"creationTime"`
			VersionCode struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"versionCode"`
			InteractionId struct {
				Text      string `xml:",chardata"`
				Extension string `xml:"extension,attr"`
				Root      string `xml:"root,attr"`
			} `xml:"interactionId"`
			ProcessingCode struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"processingCode"`
			ProcessingModeCode struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"processingModeCode"`
			AcceptAckCode struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"acceptAckCode"`
			Receiver struct {
				Text     string `xml:",chardata"`
				TypeCode string `xml:"typeCode,attr"`
				Device   struct {
					Text           string `xml:",chardata"`
					ClassCode      string `xml:"classCode,attr"`
					DeterminerCode string `xml:"determinerCode,attr"`
					ID             struct {
						Text                   string `xml:",chardata"`
						AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
						Root                   string `xml:"root,attr"`
					} `xml:"id"`
					AsAgent struct {
						Text                    string `xml:",chardata"`
						ClassCode               string `xml:"classCode,attr"`
						RepresentedOrganization struct {
							Text           string `xml:",chardata"`
							ClassCode      string `xml:"classCode,attr"`
							DeterminerCode string `xml:"determinerCode,attr"`
							ID             struct {
								Text                   string `xml:",chardata"`
								AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
								Root                   string `xml:"root,attr"`
							} `xml:"id"`
						} `xml:"representedOrganization"`
					} `xml:"asAgent"`
				} `xml:"device"`
			} `xml:"receiver"`
			Sender struct {
				Text     string `xml:",chardata"`
				TypeCode string `xml:"typeCode,attr"`
				Device   struct {
					Text           string `xml:",chardata"`
					ClassCode      string `xml:"classCode,attr"`
					DeterminerCode string `xml:"determinerCode,attr"`
					ID             struct {
						Text string `xml:",chardata"`
						Root string `xml:"root,attr"`
					} `xml:"id"`
					AsAgent struct {
						Text                    string `xml:",chardata"`
						ClassCode               string `xml:"classCode,attr"`
						RepresentedOrganization struct {
							Text           string `xml:",chardata"`
							ClassCode      string `xml:"classCode,attr"`
							DeterminerCode string `xml:"determinerCode,attr"`
							ID             struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
						} `xml:"representedOrganization"`
					} `xml:"asAgent"`
				} `xml:"device"`
			} `xml:"sender"`
			Acknowledgement struct {
				Text     string `xml:",chardata"`
				TypeCode struct {
					Text string `xml:",chardata"`
					Code string `xml:"code,attr"`
				} `xml:"typeCode"`
				TargetMessage struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text      string `xml:",chardata"`
						Extension string `xml:"extension,attr"`
						Root      string `xml:"root,attr"`
					} `xml:"id"`
				} `xml:"targetMessage"`
			} `xml:"acknowledgement"`
			ControlActProcess struct {
				Text      string `xml:",chardata"`
				ClassCode string `xml:"classCode,attr"`
				MoodCode  string `xml:"moodCode,attr"`
				Code      struct {
					Text       string `xml:",chardata"`
					Code       string `xml:"code,attr"`
					CodeSystem string `xml:"codeSystem,attr"`
				} `xml:"code"`
				Subject struct {
					Text                 string `xml:",chardata"`
					ContextConductionInd string `xml:"contextConductionInd,attr"`
					TypeCode             string `xml:"typeCode,attr"`
					RegistrationEvent    struct {
						Text      string `xml:",chardata"`
						ClassCode string `xml:"classCode,attr"`
						MoodCode  string `xml:"moodCode,attr"`
						ID        struct {
							Text       string `xml:",chardata"`
							NullFlavor string `xml:"nullFlavor,attr"`
						} `xml:"id"`
						StatusCode struct {
							Text string `xml:",chardata"`
							Code string `xml:"code,attr"`
						} `xml:"statusCode"`
						Subject1 struct {
							Text     string `xml:",chardata"`
							TypeCode string `xml:"typeCode,attr"`
							Patient  struct {
								Text      string `xml:",chardata"`
								ClassCode string `xml:"classCode,attr"`
								ID        []struct {
									Text                   string `xml:",chardata"`
									AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
									Extension              string `xml:"extension,attr"`
									Root                   string `xml:"root,attr"`
								} `xml:"id"`
								StatusCode struct {
									Text string `xml:",chardata"`
									Code string `xml:"code,attr"`
								} `xml:"statusCode"`
								EffectiveTime struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"effectiveTime"`
								PatientPerson struct {
									Text           string `xml:",chardata"`
									ClassCode      string `xml:"classCode,attr"`
									DeterminerCode string `xml:"determinerCode,attr"`
									Name           struct {
										Text   string `xml:",chardata"`
										Use    string `xml:"use,attr"`
										Given  string `xml:"given"`
										Family string `xml:"family"`
									} `xml:"name"`
									Telecom []struct {
										Text  string `xml:",chardata"`
										Use   string `xml:"use,attr"`
										Value string `xml:"value,attr"`
									} `xml:"telecom"`
									AdministrativeGenderCode struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
									} `xml:"administrativeGenderCode"`
									BirthTime struct {
										Text  string `xml:",chardata"`
										Value string `xml:"value,attr"`
									} `xml:"birthTime"`
									DeceasedInd struct {
										Text  string `xml:",chardata"`
										Value string `xml:"value,attr"`
									} `xml:"deceasedInd"`
									MultipleBirthInd struct {
										Text  string `xml:",chardata"`
										Value string `xml:"value,attr"`
									} `xml:"multipleBirthInd"`
									Addr struct {
										Text              string `xml:",chardata"`
										StreetAddressLine string `xml:"streetAddressLine"`
										City              string `xml:"city"`
										State             string `xml:"state"`
										PostalCode        string `xml:"postalCode"`
										Country           string `xml:"country"`
									} `xml:"addr"`
									MaritalStatusCode struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
									} `xml:"maritalStatusCode"`
									AsCitizen struct {
										Text            string `xml:",chardata"`
										ClassCode       string `xml:"classCode,attr"`
										PoliticalNation struct {
											Text           string `xml:",chardata"`
											ClassCode      string `xml:"classCode,attr"`
											DeterminerCode string `xml:"determinerCode,attr"`
											Code           struct {
												Text string `xml:",chardata"`
												Code string `xml:"code,attr"`
											} `xml:"code"`
										} `xml:"politicalNation"`
									} `xml:"asCitizen"`
									AsMember struct {
										Text      string `xml:",chardata"`
										ClassCode string `xml:"classCode,attr"`
										Group     struct {
											Text           string `xml:",chardata"`
											ClassCode      string `xml:"classCode,attr"`
											DeterminerCode string `xml:"determinerCode,attr"`
											Code           struct {
												Text           string `xml:",chardata"`
												Code           string `xml:"code,attr"`
												CodeSystem     string `xml:"codeSystem,attr"`
												CodeSystemName string `xml:"codeSystemName,attr"`
											} `xml:"code"`
										} `xml:"group"`
									} `xml:"asMember"`
									BirthPlace struct {
										Text string `xml:",chardata"`
										Addr struct {
											Text string `xml:",chardata"`
											City string `xml:"city"`
										} `xml:"addr"`
									} `xml:"birthPlace"`
								} `xml:"patientPerson"`
								ProviderOrganization struct {
									Text           string `xml:",chardata"`
									ClassCode      string `xml:"classCode,attr"`
									DeterminerCode string `xml:"determinerCode,attr"`
									ID             struct {
										Text       string `xml:",chardata"`
										NullFlavor string `xml:"nullFlavor,attr"`
									} `xml:"id"`
									ContactParty struct {
										Text      string `xml:",chardata"`
										ClassCode string `xml:"classCode,attr"`
									} `xml:"contactParty"`
								} `xml:"providerOrganization"`
								SubjectOf1 struct {
									Text                  string `xml:",chardata"`
									TypeCode              string `xml:"typeCode,attr"`
									QueryMatchObservation struct {
										Text      string `xml:",chardata"`
										ClassCode string `xml:"classCode,attr"`
										MoodCode  string `xml:"moodCode,attr"`
										Code      struct {
											Text       string `xml:",chardata"`
											Code       string `xml:"code,attr"`
											CodeSystem string `xml:"codeSystem,attr"`
										} `xml:"code"`
										Value struct {
											Text  string `xml:",chardata"`
											Xsi   string `xml:"xsi,attr"`
											Value string `xml:"value,attr"`
											Type  string `xml:"type,attr"`
										} `xml:"value"`
									} `xml:"queryMatchObservation"`
								} `xml:"subjectOf1"`
							} `xml:"patient"`
						} `xml:"subject1"`
						Custodian struct {
							Text           string `xml:",chardata"`
							TypeCode       string `xml:"typeCode,attr"`
							AssignedEntity struct {
								Text      string `xml:",chardata"`
								ClassCode string `xml:"classCode,attr"`
								ID        struct {
									Text       string `xml:",chardata"`
									NullFlavor string `xml:"nullFlavor,attr"`
								} `xml:"id"`
							} `xml:"assignedEntity"`
						} `xml:"custodian"`
					} `xml:"registrationEvent"`
				} `xml:"subject"`
				QueryAck struct {
					Text    string `xml:",chardata"`
					QueryId struct {
						Text      string `xml:",chardata"`
						Extension string `xml:"extension,attr"`
						Root      string `xml:"root,attr"`
					} `xml:"queryId"`
					StatusCode struct {
						Text string `xml:",chardata"`
						Code string `xml:"code,attr"`
					} `xml:"statusCode"`
					QueryResponseCode struct {
						Text string `xml:",chardata"`
						Code string `xml:"code,attr"`
					} `xml:"queryResponseCode"`
					ResultTotalQuantity struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"resultTotalQuantity"`
					ResultCurrentQuantity struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"resultCurrentQuantity"`
					ResultRemainingQuantity struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"resultRemainingQuantity"`
				} `xml:"queryAck"`
				QueryByParameter struct {
					Text    string `xml:",chardata"`
					QueryId struct {
						Text      string `xml:",chardata"`
						Extension string `xml:"extension,attr"`
						Root      string `xml:"root,attr"`
					} `xml:"queryId"`
					StatusCode struct {
						Text string `xml:",chardata"`
						Code string `xml:"code,attr"`
					} `xml:"statusCode"`
					ResponseModalityCode struct {
						Text string `xml:",chardata"`
						Code string `xml:"code,attr"`
					} `xml:"responseModalityCode"`
					ResponsePriorityCode struct {
						Text string `xml:",chardata"`
						Code string `xml:"code,attr"`
					} `xml:"responsePriorityCode"`
					MatchCriterionList string `xml:"matchCriterionList"`
					ParameterList      struct {
						Text            string `xml:",chardata"`
						LivingSubjectId struct {
							Text  string `xml:",chardata"`
							Value struct {
								Text      string `xml:",chardata"`
								Extension string `xml:"extension,attr"`
							} `xml:"value"`
							SemanticsText string `xml:"semanticsText"`
						} `xml:"livingSubjectId"`
					} `xml:"parameterList"`
				} `xml:"queryByParameter"`
			} `xml:"controlActProcess"`
		} `xml:"PRPA_IN201306UV02"`
	} `xml:"Body"`
}
type PIXv3Response struct {
	XMLName xml.Name `xml:"Envelope"`
	S       string   `xml:"S,attr"`
	Env     string   `xml:"env,attr"`
	Header  struct {
		Action struct {
			Xmlns string `xml:"xmlns,attr"`
			S     string `xml:"S,attr"`
			Env   string `xml:"env,attr"`
		} `xml:"Action"`
		MessageID struct {
			Xmlns string `xml:"xmlns,attr"`
			S     string `xml:"S,attr"`
			Env   string `xml:"env,attr"`
		} `xml:"MessageID"`
		RelatesTo struct {
			Xmlns string `xml:"xmlns,attr"`
			S     string `xml:"S,attr"`
			Env   string `xml:"env,attr"`
		} `xml:"RelatesTo"`
		To struct {
			Xmlns string `xml:"xmlns,attr"`
			S     string `xml:"S,attr"`
			Env   string `xml:"env,attr"`
		} `xml:"To"`
	} `xml:"Header"`
	Body struct {
		PRPAIN201310UV02 struct {
			Xmlns      string `xml:"xmlns,attr"`
			ITSVersion string `xml:"ITSVersion,attr"`
			ID         struct {
				Extension string `xml:"extension,attr"`
				Root      string `xml:"root,attr"`
			} `xml:"id"`
			CreationTime struct {
				Value string `xml:"value,attr"`
			} `xml:"creationTime"`
			VersionCode struct {
				Code string `xml:"code,attr"`
			} `xml:"versionCode"`
			InteractionId struct {
				Extension string `xml:"extension,attr"`
				Root      string `xml:"root,attr"`
			} `xml:"interactionId"`
			ProcessingCode struct {
				Code string `xml:"code,attr"`
			} `xml:"processingCode"`
			ProcessingModeCode struct {
				Code string `xml:"code,attr"`
			} `xml:"processingModeCode"`
			AcceptAckCode struct {
				Code string `xml:"code,attr"`
			} `xml:"acceptAckCode"`
			Receiver struct {
				TypeCode string `xml:"typeCode,attr"`
				Device   struct {
					ClassCode      string `xml:"classCode,attr"`
					DeterminerCode string `xml:"determinerCode,attr"`
					ID             struct {
						AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
						Root                   string `xml:"root,attr"`
					} `xml:"id"`
					AsAgent struct {
						ClassCode               string `xml:"classCode,attr"`
						RepresentedOrganization struct {
							ClassCode      string `xml:"classCode,attr"`
							DeterminerCode string `xml:"determinerCode,attr"`
							ID             struct {
								AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
								Root                   string `xml:"root,attr"`
							} `xml:"id"`
						} `xml:"representedOrganization"`
					} `xml:"asAgent"`
				} `xml:"device"`
			} `xml:"receiver"`
			Sender struct {
				TypeCode string `xml:"typeCode,attr"`
				Device   struct {
					ClassCode      string `xml:"classCode,attr"`
					DeterminerCode string `xml:"determinerCode,attr"`
					ID             struct {
						Root string `xml:"root,attr"`
					} `xml:"id"`
					AsAgent struct {
						ClassCode               string `xml:"classCode,attr"`
						RepresentedOrganization struct {
							ClassCode      string `xml:"classCode,attr"`
							DeterminerCode string `xml:"determinerCode,attr"`
							ID             struct {
								Root string `xml:"root,attr"`
							} `xml:"id"`
						} `xml:"representedOrganization"`
					} `xml:"asAgent"`
				} `xml:"device"`
			} `xml:"sender"`
			Acknowledgement struct {
				TypeCode struct {
					Code string `xml:"code,attr"`
				} `xml:"typeCode"`
				TargetMessage struct {
					ID struct {
						Extension string `xml:"extension,attr"`
						Root      string `xml:"root,attr"`
					} `xml:"id"`
				} `xml:"targetMessage"`
			} `xml:"acknowledgement"`
			ControlActProcess struct {
				ClassCode string `xml:"classCode,attr"`
				MoodCode  string `xml:"moodCode,attr"`
				Code      struct {
					Code       string `xml:"code,attr"`
					CodeSystem string `xml:"codeSystem,attr"`
				} `xml:"code"`
				Subject struct {
					TypeCode          string `xml:"typeCode,attr"`
					RegistrationEvent struct {
						ClassCode string `xml:"classCode,attr"`
						MoodCode  string `xml:"moodCode,attr"`
						ID        struct {
							NullFlavor string `xml:"nullFlavor,attr"`
						} `xml:"id"`
						StatusCode struct {
							Code string `xml:"code,attr"`
						} `xml:"statusCode"`
						Subject1 struct {
							TypeCode string `xml:"typeCode,attr"`
							Patient  struct {
								ClassCode string `xml:"classCode,attr"`
								ID        []struct {
									Extension              string `xml:"extension,attr"`
									Root                   string `xml:"root,attr"`
									AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
								} `xml:"id"`
								StatusCode struct {
									Code string `xml:"code,attr"`
								} `xml:"statusCode"`
								PatientPerson struct {
									ClassCode      string `xml:"classCode,attr"`
									DeterminerCode string `xml:"determinerCode,attr"`
									Name           struct {
										Given  string `xml:"given"`
										Family string `xml:"family"`
									} `xml:"name"`
								} `xml:"patientPerson"`
							} `xml:"patient"`
						} `xml:"subject1"`
						Custodian struct {
							TypeCode       string `xml:"typeCode,attr"`
							AssignedEntity struct {
								ClassCode string `xml:"classCode,attr"`
								ID        struct {
									Root string `xml:"root,attr"`
								} `xml:"id"`
							} `xml:"assignedEntity"`
						} `xml:"custodian"`
					} `xml:"registrationEvent"`
				} `xml:"subject"`
				QueryAck struct {
					QueryId struct {
						Extension string `xml:"extension,attr"`
						Root      string `xml:"root,attr"`
					} `xml:"queryId"`
					StatusCode struct {
						Code string `xml:"code,attr"`
					} `xml:"statusCode"`
					QueryResponseCode struct {
						Code string `xml:"code,attr"`
					} `xml:"queryResponseCode"`
					ResultTotalQuantity struct {
						Value string `xml:"value,attr"`
					} `xml:"resultTotalQuantity"`
					ResultCurrentQuantity struct {
						Value string `xml:"value,attr"`
					} `xml:"resultCurrentQuantity"`
					ResultRemainingQuantity struct {
						Value string `xml:"value,attr"`
					} `xml:"resultRemainingQuantity"`
				} `xml:"queryAck"`
				QueryByParameter struct {
					QueryId struct {
						Extension string `xml:"extension,attr"`
						Root      string `xml:"root,attr"`
					} `xml:"queryId"`
					StatusCode struct {
						Code string `xml:"code,attr"`
					} `xml:"statusCode"`
					ResponsePriorityCode struct {
						Code string `xml:"code,attr"`
					} `xml:"responsePriorityCode"`
					ParameterList struct {
						PatientIdentifier struct {
							Value struct {
								AssigningAuthorityName string `xml:"assigningAuthorityName,attr"`
								Extension              string `xml:"extension,attr"`
								Root                   string `xml:"root,attr"`
							} `xml:"value"`
							SemanticsText string `xml:"semanticsText"`
						} `xml:"patientIdentifier"`
					} `xml:"parameterList"`
				} `xml:"queryByParameter"`
			} `xml:"controlActProcess"`
		} `xml:"PRPA_IN201310UV02"`
	} `xml:"Body"`
}
type PIXmResponse struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Total        int    `json:"total"`
	Link         []struct {
		Relation string `json:"relation"`
		URL      string `json:"url"`
	} `json:"link"`
	Entry []struct {
		FullURL  string `json:"fullUrl"`
		Resource struct {
			ResourceType string `json:"resourceType"`
			ID           string `json:"id"`
			Identifier   []struct {
				Use    string `json:"use,omitempty"`
				System string `json:"system"`
				Value  string `json:"value"`
			} `json:"identifier"`
			Active bool `json:"active"`
			Name   []struct {
				Use    string   `json:"use"`
				Family string   `json:"family"`
				Given  []string `json:"given"`
			} `json:"name"`
			Gender    string `json:"gender"`
			BirthDate string `json:"birthDate"`
			Address   []struct {
				Use        string   `json:"use"`
				Line       []string `json:"line"`
				City       string   `json:"city"`
				PostalCode string   `json:"postalCode"`
				Country    string   `json:"country"`
			} `json:"address"`
		} `json:"resource"`
	} `json:"entry"`
}
type PDSSearchResponse struct {
	ResourceType string    `json:"resourceType,omitempty"`
	Type         string    `json:"type,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	Total        int       `json:"total,omitempty"`
	Entry        []struct {
		FullURL string `json:"fullUrl,omitempty"`
		Search  struct {
			Score int `json:"score,omitempty"`
		} `json:"search,omitempty"`
		Resource struct {
			ResourceType string `json:"resourceType,omitempty"`
			ID           string `json:"id,omitempty"`
			Identifier   []struct {
				System    string `json:"system,omitempty"`
				Value     string `json:"value,omitempty"`
				Extension []struct {
					URL                  string `json:"url,omitempty"`
					ValueCodeableConcept struct {
						Coding []struct {
							System  string `json:"system,omitempty"`
							Version string `json:"version,omitempty"`
							Code    string `json:"code,omitempty"`
							Display string `json:"display,omitempty"`
						} `json:"coding,omitempty"`
					} `json:"valueCodeableConcept,omitempty"`
				} `json:"extension,omitempty"`
			} `json:"identifier,omitempty"`
			Meta struct {
				VersionID string `json:"versionId,omitempty"`
				Security  []struct {
					System  string `json:"system,omitempty"`
					Code    string `json:"code,omitempty"`
					Display string `json:"display,omitempty"`
				} `json:"security,omitempty"`
			} `json:"meta,omitempty"`
			Name []struct {
				ID     string `json:"id,omitempty"`
				Use    string `json:"use,omitempty"`
				Period struct {
					Start string `json:"start,omitempty"`
					End   string `json:"end,omitempty"`
				} `json:"period,omitempty"`
				Given  []string `json:"given,omitempty"`
				Family string   `json:"family,omitempty"`
				Prefix []string `json:"prefix,omitempty"`
				Suffix []string `json:"suffix,omitempty"`
			} `json:"name,omitempty"`
			Gender               string    `json:"gender,omitempty"`
			BirthDate            string    `json:"birthDate,omitempty"`
			MultipleBirthInteger int       `json:"multipleBirthInteger,omitempty"`
			DeceasedDateTime     time.Time `json:"deceasedDateTime,omitempty"`
			Address              []struct {
				ID     string `json:"id,omitempty"`
				Period struct {
					Start string `json:"start,omitempty"`
					End   string `json:"end,omitempty"`
				} `json:"period,omitempty"`
				Use        string   `json:"use,omitempty"`
				Line       []string `json:"line,omitempty"`
				PostalCode string   `json:"postalCode,omitempty"`
				Extension  []struct {
					URL       string `json:"url,omitempty"`
					Extension []struct {
						URL         string `json:"url,omitempty"`
						ValueCoding struct {
							System string `json:"system,omitempty"`
							Code   string `json:"code,omitempty"`
						} `json:"valueCoding,omitempty"`
						ValueString string `json:"valueString,omitempty"`
					} `json:"extension,omitempty"`
				} `json:"extension,omitempty"`
			} `json:"address,omitempty"`
			Telecom []struct {
				ID     string `json:"id,omitempty"`
				Period struct {
					Start string `json:"start,omitempty"`
					End   string `json:"end,omitempty"`
				} `json:"period,omitempty"`
				System    string `json:"system,omitempty"`
				Value     string `json:"value,omitempty"`
				Use       string `json:"use,omitempty"`
				Extension []struct {
					URL         string `json:"url,omitempty"`
					ValueCoding struct {
						System  string `json:"system,omitempty"`
						Code    string `json:"code,omitempty"`
						Display string `json:"display,omitempty"`
					} `json:"valueCoding,omitempty"`
				} `json:"extension,omitempty"`
			} `json:"telecom,omitempty"`
			Contact []struct {
				ID     string `json:"id,omitempty"`
				Period struct {
					Start string `json:"start,omitempty"`
					End   string `json:"end,omitempty"`
				} `json:"period,omitempty"`
				Relationship []struct {
					Coding []struct {
						System  string `json:"system,omitempty"`
						Code    string `json:"code,omitempty"`
						Display string `json:"display,omitempty"`
					} `json:"coding,omitempty"`
				} `json:"relationship,omitempty"`
				Telecom []struct {
					System string `json:"system,omitempty"`
					Value  string `json:"value,omitempty"`
				} `json:"telecom,omitempty"`
			} `json:"contact,omitempty"`
			GeneralPractitioner []struct {
				ID         string `json:"id,omitempty"`
				Type       string `json:"type,omitempty"`
				Identifier struct {
					System string `json:"system,omitempty"`
					Value  string `json:"value,omitempty"`
					Period struct {
						Start string `json:"start,omitempty"`
						End   string `json:"end,omitempty"`
					} `json:"period,omitempty"`
				} `json:"identifier,omitempty"`
			} `json:"generalPractitioner,omitempty"`
			Extension []struct {
				URL       string `json:"url,omitempty"`
				Extension []struct {
					URL                  string `json:"url,omitempty"`
					ValueCodeableConcept struct {
						Coding []struct {
							System  string `json:"system,omitempty"`
							Version string `json:"version,omitempty"`
							Code    string `json:"code,omitempty"`
							Display string `json:"display,omitempty"`
						} `json:"coding,omitempty"`
					} `json:"valueCodeableConcept,omitempty"`
					ValueDateTime time.Time `json:"valueDateTime,omitempty"`
				} `json:"extension,omitempty"`
			} `json:"extension,omitempty"`
		} `json:"resource,omitempty"`
	} `json:"entry,omitempty"`
}
type PDSRetrieveResponse struct {
	ResourceType string `json:"resourceType,omitempty"`
	ID           string `json:"id,omitempty"`
	Identifier   []struct {
		System    string `json:"system,omitempty"`
		Value     string `json:"value,omitempty"`
		Extension []struct {
			URL                  string `json:"url,omitempty"`
			ValueCodeableConcept struct {
				Coding []struct {
					System  string `json:"system,omitempty"`
					Version string `json:"version,omitempty"`
					Code    string `json:"code,omitempty"`
					Display string `json:"display,omitempty"`
				} `json:"coding,omitempty"`
			} `json:"valueCodeableConcept,omitempty"`
		} `json:"extension,omitempty"`
	} `json:"identifier,omitempty"`
	Meta struct {
		VersionID string `json:"versionId,omitempty"`
		Security  []struct {
			System  string `json:"system,omitempty"`
			Code    string `json:"code,omitempty"`
			Display string `json:"display,omitempty"`
		} `json:"security,omitempty"`
	} `json:"meta,omitempty"`
	Name []struct {
		ID     string `json:"id,omitempty"`
		Use    string `json:"use,omitempty"`
		Period struct {
			Start string `json:"start,omitempty"`
			End   string `json:"end,omitempty"`
		} `json:"period,omitempty"`
		Given  []string `json:"given,omitempty"`
		Family string   `json:"family,omitempty"`
		Prefix []string `json:"prefix,omitempty"`
		Suffix []string `json:"suffix,omitempty"`
	} `json:"name,omitempty"`
	Gender               string    `json:"gender,omitempty"`
	BirthDate            string    `json:"birthDate,omitempty"`
	MultipleBirthInteger int       `json:"multipleBirthInteger,omitempty"`
	DeceasedDateTime     time.Time `json:"deceasedDateTime,omitempty"`
	GeneralPractitioner  []struct {
		ID         string `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Identifier struct {
			System string `json:"system,omitempty"`
			Value  string `json:"value,omitempty"`
			Period struct {
				Start string `json:"start,omitempty"`
				End   string `json:"end,omitempty"`
			} `json:"period,omitempty"`
		} `json:"identifier,omitempty"`
	} `json:"generalPractitioner,omitempty"`
	ManagingOrganization struct {
		Type       string `json:"type,omitempty"`
		Identifier struct {
			System string `json:"system,omitempty"`
			Value  string `json:"value,omitempty"`
			Period struct {
				Start string `json:"start,omitempty"`
				End   string `json:"end,omitempty"`
			} `json:"period,omitempty"`
		} `json:"identifier,omitempty"`
	} `json:"managingOrganization,omitempty"`
	Extension []struct {
		URL            string `json:"url,omitempty"`
		ValueReference struct {
			Identifier struct {
				System string `json:"system,omitempty"`
				Value  string `json:"value,omitempty"`
			} `json:"identifier,omitempty"`
		} `json:"valueReference,omitempty"`
		Extension []struct {
			URL                  string `json:"url,omitempty"`
			ValueCodeableConcept struct {
				Coding []struct {
					System  string `json:"system,omitempty"`
					Version string `json:"version,omitempty"`
					Code    string `json:"code,omitempty"`
					Display string `json:"display,omitempty"`
				} `json:"coding,omitempty"`
			} `json:"valueCodeableConcept,omitempty"`
			ValueDateTime string `json:"valueDateTime,omitempty"`
			ValueString   string `json:"valueString,omitempty"`
			ValueBoolean  bool   `json:"valueBoolean,omitempty"`
		} `json:"extension,omitempty"`
		ValueAddress struct {
			City     string `json:"city,omitempty"`
			District string `json:"district,omitempty"`
			Country  string `json:"country,omitempty"`
		} `json:"valueAddress,omitempty"`
	} `json:"extension,omitempty"`
	Telecom []struct {
		ID     string `json:"id,omitempty"`
		Period struct {
			Start string `json:"start,omitempty"`
			End   string `json:"end,omitempty"`
		} `json:"period,omitempty"`
		System    string `json:"system,omitempty"`
		Value     string `json:"value,omitempty"`
		Use       string `json:"use,omitempty"`
		Extension []struct {
			URL         string `json:"url,omitempty"`
			ValueCoding struct {
				System  string `json:"system,omitempty"`
				Code    string `json:"code,omitempty"`
				Display string `json:"display,omitempty"`
			} `json:"valueCoding,omitempty"`
		} `json:"extension,omitempty"`
	} `json:"telecom,omitempty"`
	Contact []struct {
		ID     string `json:"id,omitempty"`
		Period struct {
			Start string `json:"start,omitempty"`
			End   string `json:"end,omitempty"`
		} `json:"period,omitempty"`
		Relationship []struct {
			Coding []struct {
				System  string `json:"system,omitempty"`
				Code    string `json:"code,omitempty"`
				Display string `json:"display,omitempty"`
			} `json:"coding,omitempty"`
		} `json:"relationship,omitempty"`
		Telecom []struct {
			System string `json:"system,omitempty"`
			Value  string `json:"value,omitempty"`
		} `json:"telecom,omitempty"`
	} `json:"contact,omitempty"`
	Address []struct {
		ID     string `json:"id,omitempty"`
		Period struct {
			Start string `json:"start,omitempty"`
			End   string `json:"end,omitempty"`
		} `json:"period,omitempty"`
		Use        string   `json:"use,omitempty"`
		Line       []string `json:"line,omitempty"`
		PostalCode string   `json:"postalCode,omitempty"`
		Extension  []struct {
			URL       string `json:"url,omitempty"`
			Extension []struct {
				URL         string `json:"url,omitempty"`
				ValueCoding struct {
					System string `json:"system,omitempty"`
					Code   string `json:"code,omitempty"`
				} `json:"valueCoding,omitempty"`
				ValueString string `json:"valueString,omitempty"`
			} `json:"extension,omitempty"`
		} `json:"extension,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"address,omitempty"`
}
type NotifyEvent struct {
	Body     string
	From     string
	To       string
	Server   string
	Port     string
	Password string
}

type Pathways struct {
	Pathway []Pathway
}
type Pathway struct {
	Text  string
	Value string
}
type Expressions struct {
	Expression []Expression
}
type Expression struct {
	Text  string
	Value string
}
type Comment struct {
	Taskid int    `json:"taskid"`
	Note   string `json:"note"`
}
type Comments struct {
	Comment []Comment
}
type TaskState struct {
	Taskid               int      `json:"taskid"`
	Name                 string   `json:"name"`
	TargetMet            bool     `json:"targetmet"`
	Escalated            bool     `json:"escalated"`
	Status               string   `json:"status"`
	StartBy              string   `json:"startby"`
	WDStartBy            string   `json:"wdstartby"`
	CompleteBy           string   `json:"completeby"`
	WDCompleteBy         string   `json:"wdcompleteby"`
	EscalateOn           string   `json:"escalateon"`
	WDEscalateOn         string   `json:"wdescalateon"`
	CompletionConditions []string `json:"completionconditions"`
	StartedOn            string   `json:"startedon"`
	Owner                string   `json:"owner"`
	CompletedOn          string   `json:"completedon"`
	Duration             string   `json:"duration"`
	TimeRemaining        string   `json:"timeremaining"`
}
type WorkflowTasksState struct {
	TaskState []TaskState `json:"taskstate"`
}

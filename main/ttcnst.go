/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE', which is part of this source code package.
 */

package main

const (
	ENV_NHS_OID                               = "NHS_OID"
	ENV_REG_OID                               = "REG_OID"
	ENV_IHE_PDQV3_SERVER_URL                  = "IHE_PDQV3_SERVER_URL"
	ENV_IHE_PIXV3_SERVER_URL                  = "IHE_PIXV3_SERVER_URL"
	ENV_IHE_PIXM_SERVER_URL                   = "IHE_PIXM_SERVER_URL"
	ENV_CGL_SERVER_URL                        = "CGL_SERVER_URL"
	ENV_CGL_X_API_KEY                         = "X-API-KEY"
	ENV_CGL_X_API_SECRET                      = "X-API-SECRET"
	ENV_PDQ_SERVER_TYPE                       = "PDQ_SERVER_TYPE"
	ENV_PDQ_SERVER_URL                        = "PDQ_SERVER_URL"
	ENV_PDS_SERVER_TYPE                       = "PDS_SERVER_TYPE"
	ENV_PDS_SERVER_URL                        = "PDS_SERVER_URL"
	ENV_DSUB_BROKER_URL                       = "DSUB_BROKER_URL"
	ENV_DSUB_CONSUMER_URL                     = "DSUB_CONSUMER_URL"
	ENV_SERVER_URL                            = "SERVER_URL"
	ENV_TUK_DB_URL                            = "TUK_DB_URL"
	ENV_DB_HOST                               = "DB_HOST"
	ENV_DB_NAME                               = "DB_NAME"
	ENV_DB_PORT                               = "DB_PORT"
	ENV_DB_USER                               = "DB_USER"
	ENV_DB_PASSWORD                           = "DB_PASSWORD"
	ENV_DEBUG_MODE                            = "DEBUG_MODE"
	ENV_DEBUG_DB                              = "DEBUG_DB"
	ENV_DEBUG_DB_ERROR                        = "DEBUG_DB_ERROR"
	ENV_LOGO_FILE                             = "LOGO_FILE"
	ENV_SERVER_PORT                           = "SERVER_PORT"
	ENV_PERSIST_TEMPLATES                     = "PERSIST_TEMPLATES"
	ENV_PERSIST_DEFINITIONS                   = "PERSIST_DEFINITIONS"
	PDQ_SERVER_TYPE_IHE_PIXM                  = "pixm"
	PDQ_SERVER_TYPE_IHE_PDQV3                 = "pdqv3"
	PDQ_SERVER_TYPE_IHE_PIXV3                 = "pixv3"
	PDQ_SERVER_TYPE_NHS_PDS                   = "pds"
	PDQ_SERVER_TYPE_CGL                       = "cgl"
	PDQ_SERVER_TYPE_ALL                       = "all"
	TASK                                      = "task"
	DASHBOARD                                 = "dashboard"
	SPA                                       = "spa"
	XML                                       = "xml"
	JSON                                      = "json"
	HTML                                      = "html"
	PATHWAY                                   = "pathway"
	DEMO                                      = "demo"
	PULLPOINTS                                = "pullpoints"
	CONFIG                                    = "config"
	AUDITS                                    = "audits"
	EVENT_ACKS                                = "eventacks"
	EVENTS                                    = "events"
	ID_MAPS                                   = "idmaps"
	TEMPLATES                                 = "templates"
	STATICS                                   = "statics"
	SERVICE_STATES                            = "servicestates"
	SUBSCRIBE                                 = "subscribe"
	SUBSCRIBER                                = "subscriber"
	SUBSCRIPTIONS                             = "subscriptions"
	REGISTER                                  = "register"
	WIDGET                                    = "widget"
	ADMIN                                     = "admin"
	TEMPLATE                                  = "template"
	CREATE                                    = "create"
	CREATOR                                   = "creator"
	CANCEL                                    = "cancel"
	LIST                                      = "list"
	ACK                                       = "ack"
	ATTACHMENT                                = "attachment"
	TOPICS                                    = "topics"
	EXPRESSIONS                               = "expressions"
	PIX_QUERY                                 = "pixm"
	STS                                       = "sts"
	STATEID                                   = "stateID"
	PATIENT                                   = "patient"
	SERVICES                                  = "services"
	WORKFLOW                                  = "workflow"
	WORKFLOWS                                 = "workflows"
	ALERT                                     = "alert"
	RST_ISSUE                                 = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/RST/Issue"
	SUBSCRIPTION_REFERENCE_PREFIX             = "https://NotificationBrokerServer/Subscription/"
	SOAP_ACTION_UNSUBSCRIBE_REQUEST           = "http://docs.oasis-open.org/wsn/bw2/SubscriptionManager/UnsubscribeRequest"
	SOAP_ACTION_SUBSCRIBE_REQUEST             = "http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest"
	SOAP_ACTION_PIXV3_Request                 = "urn:hl7-org:v3:PRPA_IN201309UV02"
	SOAP_ACTION_PDQV3_Request                 = "urn:hl7-org:v3:PRPA_IN201305UV02"
	SOAP_ACTION                               = "SOAPAction"
	CONTENT_TYPE                              = "Content-Type"
	TEXT_HTML                                 = "text/html"
	TEXT_PLAIN                                = "text/plain"
	IMAGE_BASE64                              = "image/base64"
	APPLICATION_BASE64                        = "application/base64"
	APPLICATION_JSON                          = "application/json"
	APPLICATION_XML                           = "application/xml"
	SOAP_XML                                  = "application/soap+xml"
	FORM_URL_ENCODED                          = "application/x-www-form-urlencoded"
	FORM_URL_ENCODED_CHARSET_UTF_8            = "application/x-www-form-urlencoded; charset=UTF-8"
	FORM_MULTI_PART                           = "multipart/form-data"
	ACCEPT                                    = "Accept"
	AUTHORIZATION                             = "Authorization"
	ALL                                       = "*/*"
	CONNECTION                                = "Connection"
	KEEP_ALIVE                                = "keep-alive"
	UNSUBSCRIBE_RESPONSE                      = "UnsubscribeResponse"
	PULLPOINT                                 = "pullpoint"
	IN_PROGRESS                               = "IN_PROGRESS"
	TEXT_XML_CHARSET_UTF_8                    = "text/xml; charset=utf-8"
	A_ADDRESS                                 = "a:Address"
	OK                                        = "OK"
	SUBMISSION_TIME                           = "submissionTime"
	CREATION_TIME                             = "creationTime"
	REPOSITORY_UID                            = "repositoryUniqueId"
	SOURCE_PATIENT_ID                         = "sourcePatientId"
	URN_XDS_PID                               = "urn:uuid:58a6f841-87b3-4a3e-92fd-a8ffeff98427"
	URN_XDS_DOCUID                            = "urn:uuid:2e82c1f6-a085-4c72-9da3-8640a32e42ab"
	URN_SUBMISSION_SOURCE_ID                  = "urn:uuid:127d4267-6169-4896-b46c-b34d4d1b6d5d"
	URN_SUBMISSION_UID                        = "urn:uuid:96fdda7c-d067-4183-912e-bf5ee74998a8"
	URN_SUBMISSION_PID                        = "urn:uuid:6b5aea1a-874d-4603-a4bc-96a0a7b38446"
	URN_CLASS_CODE                            = "urn:uuid:41a5887f-8865-4c09-adf7-e362475b143a"
	URN_CONF_CODE                             = "urn:uuid:f4f85eac-e6cb-4883-b524-f2705394840f"
	URN_FORMAT_CODE                           = "urn:uuid:a09d5840-386c-46f2-b5ad-9c3699a4309d"
	URN_FACILITY_CODE                         = "urn:uuid:f33fb8ac-18af-42cc-ae0e-ed0b0bdb91e1"
	URN_PRACTICE_CODE                         = "urn:uuid:cccf5598-8b07-4b77-a05e-ae952c785ead"
	URN_TYPE_CODE                             = "urn:uuid:f0306f51-975f-434e-a61c-c59651d33983"
	URN_AUTHOR                                = "urn:uuid:93606bcf-9494-43ec-9b4e-a7748d1a838d"
	URN_EVENT_LIST                            = "urn:uuid:2c6b8cb7-8b2a-4051-b291-b1ae6a575ef4"
	AUTHOR_PERSON                             = "authorPerson"
	AUTHOR_INSTITUTION                        = "authorInstitution"
	AUTHOR_SPECIALITY                         = "authorSpecialty"
	AUTHOR_ROLE                               = "authorRole"
	XMLNS_XSI                                 = "http://www.w3.org/2001/XMLSchema-instance"
	XMLNS_XSD                                 = "http://www.w3.org/2001/XMLSchema"
	XMLNS                                     = "http://docs.oasis-open.org/wsn/b-2"
	INSERT                                    = "insert"
	SELECT                                    = "select"
	DELETE                                    = "delete"
	DEPRECATE                                 = "deprecate"
	ROLLBACK                                  = "rollback"
	REPLACE                                   = "replace"
	UPDATE                                    = "update"
	ISPUBLISHED                               = "ispublished"
	APPEND                                    = "append"
	XDSDOMAIN                                 = "XDSDOMAIN"
	TEMPLATE_DATA                             = "tmpltdata"
	WorkflowInstanceId                        = "^^^^urn:ihe:iti:xdw:2013:workflowInstanceId"
	XDWNameSpace                              = "urn:ihe:iti:xdw:2011"
	XDWNameLocal                              = "xdw"
	HL7NameSpace                              = "urn:hl7-org:v3"
	HL7NameLocal                              = "hl7"
	WHTNameSpace                              = "http://docs.oasis-open.org/ns/bpel4people/ws-humantask/types/200803"
	WHTNameLocal                              = "ws-ht"
	WorkflowDocumentXsi                       = "http://www.w3.org/2001/XMLSchema-instance"
	WorkflowDocumentSchemaLocation            = "urn:ihe:iti:xdw:2011 XDW-2014-12-23.xsd"
	XDS_REGISTERED                            = "urn:ihe:iti:xdw:2011:XDSregistered"
	MEDIA_TYPES                               = "http://www.iana.org/assignments/media-types"
	ASSERTION_SUBJECT_ID                      = "urn:oasis:names:tc:xspa:1.0:subject:subject-id"
	ASSERTION_ORGANISATION                    = "urn:oasis:names:tc:xspa:1.0:subject:organization"
	ASSERTION_ROLE                            = "urn:oasis:names:tc:xacml:2.0:subject:role"
	ASSERTION_POU                             = "urn:oasis:names:tc:xspa:1.0:subject:purposeofuse"
	RETRIEVE_DOCUMENT_SET                     = "urn:ihe:iti:2007:RetrieveDocumentSet"
	NOTIFICATION_ELEMENT                      = "<wsnt:NotificationMessage>"
	DSUB_NOTIFY_ELEMENT                       = "wsnt:Notify"
	SQL_DEFAULT_IDMAPS                        = "SELECT * FROM idmaps"
	SQL_DEFAULT_SERVICESTATES                 = "SELECT * FROM servicestates"
	SQL_DEFAULT_XDWS                          = "SELECT * FROM xdws"
	SQL_DEFAULT_EVENTS                        = "SELECT * FROM events"
	SQL_DEFAULT_TEMPLATES                     = "SELECT * FROM templates"
	SQL_DEFAULT_STATICS                       = "SELECT * FROM statics"
	SQL_DEFAULT_WORKFLOWS                     = "SELECT * FROM workflows"
	SQL_DEFAULT_SUBSCRIPTIONS                 = "SELECT * FROM subscriptions"
	DSUB_TOPIC_TYPE_CODE                      = "$XDSDocumentEntryTypeCode"
	NHS_OID_DEFAULT                           = "2.16.840.1.113883.2.1.4.1"
	REGION_OID                                = "Region_OID"
	HOME_COMMUNITY_OID                        = "Home_Community_OID"
	FORMAT_JSON_PRETTY                        = "&_format=json&_pretty=true"
	URN_OID_PREFIX                            = "urn:oid:"
	HTTP_HEADER_CACHE_CONTROL                 = "Cache-Control"
	HTTP_HEADER_CLEAR_SITE_DATA               = "Clear-Site-Data"
	HTTP_HEADER_CONTENT_SECURITY_POLICY       = "Content-Security-Policy"
	HTTP_HEADER_CROSS_ORIGIN_EMBEDDER_POLICY  = "Cross-Origin-Embedder-Policy"
	HTTP_HEADER_CROSS_ORIGIN_OPENER_POLICY    = "Cross-Origin-Opener-Policy"
	HTTP_HEADER_CROSS_ORIGIN_RESOURCE_POLICY  = "Cross-Origin-Resource-Policy"
	HTTP_HEADER_PERMISSIONS_POLICY            = "Permissions-Policy"
	HTTP_HEADER_REFERRER_POLICY               = "Referrer-Policy"
	HTTP_HEADER_STRICT_TRANSPORT_SECURITY     = "Strict-Transport-Security"
	HTTP_HEADER_X_CONTENT_TYPE_OPTIONS        = "X-Content-Type-Options"
	HTTP_HEADER_X_FRAME_OPTIONS               = "X-Frame-Options"
	HTTP_HEADER_X_XSS_PROTECTION              = "X-XSS-Protection"
	HTTP_HEADER_CONTENT_TYPE                  = "Content-Type"
	HTTP_HEADER_SERVER                        = "Server"
	HTTP_PATH_PUBLISHER_CREATOR               = "/spa/publisher/creator"
	HTTP_PATH_PUBLISHER_DEFINITION            = "/spa/publisher/definition"
	HTTP_PATH_PUBLISHER_TEMPLATE              = "/spa/publisher/template"
	HTTP_PATH_PUBLISHER_META                  = "/spa/publisher/meta"
	HTTP_PATH_PUBLISHER_CODEMAP               = "/spa/publisher/codemap"
	HTTP_PATH_PUBLISHER_IMAGE                 = "/spa/publisher/image"
	HTTP_PATH_PUBLISHER_EVENT                 = "/spa/publisher/event"
	HTTP_PATH_PUBLISHER_DSUB                  = "/dsub"
	HTTP_PATH_PUBLISHER_UPLOAD                = "/spa/publisher/upload"
	HTTP_PATH_CONSUMER_STATE                  = "/spa/consumer/state"
	HTTP_PATH_CONSUMER_STATIC                 = "/spa/consumer/static"
	HTTP_PATH_CONSUMER_CODEMAP                = "/spa/consumer/codemap"
	HTTP_PATH_CONSUMER_CLEAR_CACHE            = "/spa/consumer/clearcache"
	HTTP_PATH_CONSUMER_PATIENT                = "/spa/consumer/patient"
	HTTP_PATH_CONSUMER_PATHWAYS               = "/spa/consumer/pathways"
	HTTP_PATH_CONSUMER_EXPRESSIONS            = "/spa/consumer/expressions"
	HTTP_PATH_CONSUMER_COMMENTS               = "/spa/consumer/comments"
	HTTP_PATH_CONSUMER_XDWS                   = "/spa/consumer/xdws"
	HTTP_PATH_CONSUMER_XDW                    = "/spa/consumer/xdw"
	HTTP_PATH_CONSUMER_MYSUBS                 = "/spa/consumer/mysubs"
	HTTP_PATH_CONSUMER_NEWSUB                 = "/spa/consumer/newsub"
	HTTP_PATH_CONSUMER_DELSUB                 = "/spa/consumer/delsub"
	HTTP_PATH_CONSUMER_CREATOR                = "/spa/consumer/creator"
	HTTP_PATH_CONSUMER_DEFINITION             = "/spa/consumer/definition"
	HTTP_PATH_CONSUMER_TEMPLATE               = "/spa/consumer/template"
	HTTP_PATH_PUBLISH                         = "/spa/publish"
	HTTP_PATH_CONSUMER_TEMPLATES              = "/spa/consumer/templates"
	HTTP_PATH_CONSUMER_UPLOAD                 = "/spa/consumer/upload"
	HTTP_PATH_CONSUMER_EVENTS                 = "/spa/consumer/events"
	HTTP_PATH_CONSUMER_DEFINITIONS            = "/spa/consumer/definitions"
	HTTP_PATH_CONSUMER_SPA                    = "/spa"
	HTTP_PATH_API_STATE                       = "/api/state"
	HTTP_PATH_API_STATE_PATIENT               = "/api/state/patient"
	HTTP_PATH_API_STATE_DASHBOARD             = "/api/state/dashboard"
	HTTP_PATH_API_STATE_EVENTS                = "/api/state/events"
	HTTP_PATH_API_STATE_TERMINOLOGY           = "/api/state/terminology"
	HTTP_PATH_API_STATE_WORKFLOWS             = "/api/state/workflows"
	HTTP_PATH_API_STATE_WORKFLOWS_COUNT       = "/api/state/workflows/count"
	HTTP_PATH_API_STATE_WORKFLOW              = "/api/state/workflow"
	HTTP_PATH_API_STATE_SUBSCRIPTIONS         = "/api/state/subscriptions"
	HTTP_PATH_API_STATE_PATHWAYS              = "/api/state/pathways"
	HTTP_PATH_API_STATE_TASK_STATUS           = "/api/state/task/status"
	HTTP_PATH_API_STATE_TASKS_STATUS          = "/api/state/tasks/status"
	HTTP_PATH_API_STATE_EXPRESSIONS           = "/api/state/expressions"
	HTTP_PATH_API_STATE_COMMENTS              = "/api/state/comments"
	HTTP_PATH_API_STATE_DEFINITION            = "/api/state/definition"
	HTTP_PATH_API_STATE_META                  = "/api/state/meta"
	HTTP_PATH_EVENT_DSUB                      = "/spa/event/dsub"
	HTTP_PATH_EVENT_USER                      = "/spa/event/user"
	HTTP_PATH_ROOT                            = ""
	XDW_TASKEVENTTYPE_WORKFLOW_COMPLETED      = "WORKFLOW_COMPLETED"
	XDW_TASKEVENTTYPE_CLAIM                   = "claim"
	XDW_TASKEVENTTYPE_START                   = "start"
	XDW_TASKEVENTTYPE_COMPLETE                = "complete"
	XDW_TASKEVENTTYPE_ATTACHMENT              = "ATTACHMENT"
	XDW_TASKEVENTTYPE_COMMENT                 = "COMMENT"
	XDW_TASKEVENTTYPE_ESCALATED               = "ESCALATED"
	XDW_OPERATION_WORKFLOW                    = "WORKFLOW"
	XDW_OPERATION_CREATE_WORKFLOW             = "CREATE_WORKFLOW"
	XDW_OPERATION_START_WORKFLOW              = "START_WORKFLOW"
	XDW_OPERATION_CLOSE_WORKFLOW              = "CLOSE_WORKFLOW"
	XDW_OPERATION_CREATE_TASK                 = "CREATE_TASK"
	XDW_OPERATION_GET_NHS                     = "getnhs"
	XDW_OPERATION_GET_EXPRESSIONS             = "getExpressions"
	XDW_OPERATION_GET_TASKS                   = "getTasks"
	XDW_OPERATION_GET_COMMENT                 = "getComments"
	XDW_OPERATION_GET_PATHWAYS                = "getPathways"
	XDW_OPERATION_ADD_ATTACHMENT              = "addAttachment"
	XDW_OPERATION_ADD_COMMENT                 = "addComment"
	XDW_OPERATION_CLAIM                       = "claim"
	XDW_OPERATION_COMPLETE                    = "complete"
	XDW_OPERATION_DELEGATE                    = "delegate"
	XDW_OPERATION_GET_ATTACHMENT              = "getAttachment"
	XDW_OPERATION_GET_COMMENTS                = "getComments"
	XDW_OPERATION_GET_DESCRIPTION             = "getTaskDescription"
	XDW_OPERATION_GET_DETAILS                 = "getTaskDetails"
	XDW_OPERATION_GET_HISTORY                 = "getTaskHistory"
	XDW_OPERATION_EDIT_CODEMAP                = "edit"
	XDW_OPERATION_INSERT_CODEMAP              = "insert"
	XDW_OPERATION_DELETE_CODEMAP              = "delete"
	XDW_OPERATION_TASK_START                  = "start"
	XDW_OPERATION_TASK_QUERY                  = "query"
	XDW_OPERATION_INIT_TEMPLATES              = "inittmplts"
	XDW_OPERATION_REGISTER_TEMPLATE           = "registetmplt"
	QUERY_PARAM_ID                            = "id"
	QUERY_PARAM_ACT                           = "act"
	QUERY_PARAM_ACTION                        = "action"
	QUERY_PARAM_TASK                          = "task"
	QUERY_PARAM_OP                            = "op"
	QUERY_PARAM_VERSION                       = "vers"
	QUERY_PARAM_STATUS                        = "status"
	QUERY_PARAM_TASK_ID                       = "taskid"
	QUERY_PARAM_ROLE                          = "role"
	QUERY_PARAM_ORG                           = "org"
	QUERY_PARAM_USER                          = "user"
	QUERY_PARAM_PATHWAY                       = "pathway"
	QUERY_PARAM_TOPIC                         = "topic"
	QUERY_PARAM_EXPRESSION                    = "expression"
	QUERY_PARAM_COMMENTS                      = "comments"
	QUERY_PARAM_NHS                           = "nhs"
	QUERY_PARAM_PID                           = "pid"
	QUERY_PARAM_PID_OID                       = "pidoid"
	QUERY_PARAM_CONFIG                        = "config"
	QUERY_PARAM_DOCREF                        = "docref"
	QUERY_PARAM_FORMAT                        = "_format"
	QUERY_PARAM_TEMPLATE                      = "template"
	QUERY_PARAM_EMAIL                         = "email"
	QUERY_PARAM_PUBLISHED                     = "published"
	QUERY_PARAM_ATTACHMENT_TYPE               = "attachmenttype"
	QUERY_PARAM_EVENT_TYPE                    = "eventtype"
	QUERY_PARAM_LID                           = "lid"
	QUERY_PARAM_MID                           = "mid"
	QUERY_PARAM_CONF_CODE                     = "confcode"
	QUERY_PARAM_METHOD                        = "method"
	QUERY_PARAM_BROKER_REF                    = "brokerref"
	STATUS_MET                                = "MET"
	STATUS_MISSED                             = "MISSED"
	STATUS_ESCALATED                          = "ESCALATED"
	STATUS_OPEN                               = "OPEN"
	STATUS_CLOSED                             = "CLOSED"
	STATUS_COMPLETE                           = "COMPLETE"
	STATUS_CREATED                            = "CREATED"
	SECURITY_HEADER_CACHE_CONTROL             = "private, max-age=0, no-cache"
	SECURITY_HEADER_CLEAR_SITE_DATA           = "\"*\""
	SECURITY_HEADER_SECURITY_POLICY           = "default-src 'self' data:; img-src 'self' localhost:8080; font-src 'self'; style-src 'self' 'unsafe-inline'; script-src 'self' 'unsafe-inline'; frame-ancestors 'self' tiani-spirit.co.uk *.nhs.net *.lpres.co.uk;"
	SECURITY_HEADER_EMBEDDER_POLICY           = "require-corp"
	SECURITY_HEADER_OPENER_POLICY             = "same-origin-allow-popups"
	SECURITY_HEADER_RESOURCE_POLICY           = "same-origin"
	SECURITY_HEADER_PERMISSIONS_POLICY        = "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), cross-origin-isolated=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), screen-wake-lock=(), sync-xhr=(), usb=(), web-share=(), xr-spatial-tracking=()"
	SECURITY_HEADER_STRICT_TRANSPORT_SECURITY = "max-age=63072000"
	SECURITY_HEADER_X_CONTENT_TYPE_OPTIONS    = "nosniff"
	SECURITY_HEADER_X_FRAME_OPTIONS           = "DENY"
	SECURITY_HEADER_X_XSS_PROTECTION          = "0"
)

# ttxdwuser

This package combines various IHE Profiles a UK NHS SPINE PDS Consumer and a UK CGL Consumer to provide Cross Community Workflow and Event Services

The Backend utilises the following IHE Profiles and Actors
    
    Profile     Actor
    XDS         Repositories
    XDS         Registry
    DSUB        Broker
    XDW         Content Creator
    XDW         Content Consumer
    XDW         Content Updator
    PIXm        Consumer

DSUB Implementation

    It also implements DSUB 'pullpoint' functionality by providing a DSUB 'proxy' Broker service. 
    This service parses workflow definitions and registers subscriptions with the SCR DSUB Broker for applicable workflow tasks. 
    Users/Epr's can use this proxy DSUB Broker service to create their own subscriptions to specifc workflows and events and receive email notifications when those workflow events occure.

Patient Services

    In addition to the PIXm service, patient information is also retrieved from the UK NHS SPINE PDS service and from the Change Grow Live (CGL) interface

Deployment

    The Event Service utilises environment variables for configuration of the service. 
    If the service is deployed in AWS as a lambda service, the lambda configuration/environment variables are used. 
    Deployment on a local server utilises the envvars.json file in the root folder of the application. 
    When the application is started or triggered it automattically loads the environment vars from the relevant configuration (file or AWS vars).

    Example envvars.json 

        {
            "DB_USER": "your db admin username",
            "DB_PASSWORD": "your db admin password",
            "DB_HOST": "your db host",
            "DB_PORT": "your db host port",
            "DB_NAME": "your db name",
            "DEBUG_MODE": "true or false",
            "DEBUG_DB": "true or false",
            "DEBUG_DB_ERROR": "true or false",
            "DSUB_BROKER_URL": "HIE DSUB Broker URL",
            "DSUB_CONSUMER_URL": "the event service URL",
            "LOGO_FILE": "logo file to use in html responses",
            "SERVER_PORT": "the event service port",
            "SERVER_URL": "the event service URL",
            "SERVER_NAME": "the event service name",
            "SCR_URL": "your HIE MHD endpoint URL",
            "REG_OID": "your HIE Reg/Regional OID",
            "PIXM_SERVER_URL": "your HIE PIXm endpoint URL",
            "PDS_SERVER_URL": "your hIE PDS Fhir endpoint URL",
            "CGL_SERVER_URL": "your CGL endpoint URL",
            "CGL_SERVER_X_API_KEY": "your CGL API key",
            "CGL_SERVER_X_API_SECRET": "your CGL API secret",
            "S3_PUBLISH_FILES": "Not Implemented",
            "SMTP_USER": "your smtp user",
            "SMTP_SERVER": "your smtp server",
            "SMTP_PORT": "your smtp port",
            "SMTP_PASSWORD": "your smtp password",
            "SMTP_SUBJECT": "the subject text for email notifications",
            "PERSIST_TEMPLATES": "true or false",
            "PERSIST_DEFINITIONS": "true or false",
            "CALENDAR_MODE": "calendardays or workingdays",
            "TIME_LOCATION" : "Europe/London",
            "TIME_LOCALE" : "en_GB"
        }
    
    Database

        Create a database called tuk
        'use tuk'
        The build folder contains a database folder with tuk.sql that can be used to create the various event service 'tuk' tables. 


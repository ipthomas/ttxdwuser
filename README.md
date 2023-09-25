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

    It also implements DSUB 'pullpoint' functionality by providing a DSUB 'proxy' Broker service. This service parses workflow definitions and registers subscriptions with the SCR Broker for applicable workflow tasks. Users/Epr's can use this proxy DSUB Broker service to create their subscriptions to specifc workflows and events and receive email notifications when those workflow events occure.

Patient Services

    In addition to the PIXm service, patient information is also retrieved from the UK NHS SPINE PDS service and from the Change Grow Live (CGL) interface

Deployment



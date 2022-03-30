# go-vesca

This repository contains a PoC of Vertically Sliced + Clean Architecture in Go

## Motivation

- As a **Software Architect**
- I want to **create a distributed application, using Go**
- For **fun and study**

## Purpose

In this project, we will create a distributed application that aims to illustrate a number of concepts using Golang:

- Clean Architecture
- Feature Slicing
- Event Sourcing/CQRS
- Naive implementation of the Actor Model

## Requirements

A system is needed that collects information from temperature/humidity sensors in a greenhouse and displays it in a real-time dashboard.

The dashboard must have the capability to check the situation at any moment in time.

A VAC-actuator in the greenhouse will act when the temperature or humidity exceeds certain boundaries.

## Reference Materials

-  


## Environment Variables

### TH-Sensor
```dotenv
## Setup
GO_VESCA_GREENHOUSE_ID      # The UUID for the Greenhouse 
GO_VESCA_GREENHOUSE_NAME    # default: "Greenhouse 0"
GO_VESCA_TH_SENSOR_ID       # The UUID for the TH-Sensor
GO_VESCA_TH_SENSOR_NAME     # default: "TH-Sensor 0"
## NATS
GO_VESCA_NATS_URL           # The url of the NATS infrastructure (default: "http://127.0.0.1:4200")
GO_VESCA_NATS_USER          # The NATS user name (default: "a", for dev only!)
GO_VESCA_NATS_PWD           # The NATS password (default: "a", for dev only!)
```

### VAC-Actuator
```dotenv
## Setup
GO_VESCA_GREENHOUSE_ID          # The UUID for the Greenhouse 
GO_VESCA_GREENHOUSE_NAME        # default: "Greenhouse 0"
GO_VESCA_VAC_ACTUATOR_ID        # The UUID for the VAC-Actuator
GO_VESCA_VAC_ACTUATOR_NAME      # default "VAC-Actuator 0"
## NATS
GO_VESCA_NATS_URL               # The url of the NATS infrastructure (default: "http://127.0.0.1:4200")
GO_VESCA_NATS_USER              # The NATS user name (default: "a", for dev only!)
GO_VESCA_NATS_PWD               # The NATS password (default: "a", for dev only!)
```


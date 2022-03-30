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

## Environment Variables

### TH-Sensor
```dotenv
GO_VESCA_GREENHOUSE_ID    #The UUID for the Greenhouse 
GO_VESCA_TH_SENSOR_ID     #The UUID for the TH-Sensor
GO_VESCA_NATS_URL         #The url of the NATS infrastructure
```

VACC-Actuator

- GREENHOUSE_ID |> The UUID for the Greenhouse
- VAC_ACTUATOR_ID |> The UUID for the VAC-Actuator

## Requirements

A system is needed that collects information from temperature/humidity sensors in a greenhouse and displays it in a real-time dashboard.

The dashboard must have the capability to check the situation at any moment in time.

A VAC-actuator in the greenhouse will act when the temperature or humidity exceeds certain boundaries.

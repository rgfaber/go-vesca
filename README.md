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

- AERO_DATA_BOX_API_KEY |> The RapidAPI API Key for AeroDataBox

## Requirements

A system is needed that collects information from temperature/humidity sensors in a greenhouse and displays it in a real-time dashboard. The dashboard must have the capability to check the situation at any moment in time. An actuator in the greenhouse will be triggered to respond to when temperature/humidity exceeds certain boundaries.

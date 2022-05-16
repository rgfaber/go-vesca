package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize"
)

func main() {
	sup := NewSupervisor(initialize.Feature)
	go sup.Supervise()
	select {}
}

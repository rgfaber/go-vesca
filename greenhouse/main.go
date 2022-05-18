package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/initialize"
)

func main() {
	sup := SingletonSupervisor(initialize.Feature)
	sup.Supervise()
	select {}
}

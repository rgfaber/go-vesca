package main

import (
	"fmt"
	"os"
)

func main() {
	if TheSupervisor == nil {
		err := fmt.Errorf("Could not create Supervisor.")
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	go TheSupervisor.Supervise()
	select {}
}

package main

import (
	"gobridge/bridge"
	"log"
)

func main() {
	br := bridge.NewBridge()

	ports := []string{"eth3", "eth4", "eth5", "eth6"}
	for _, name := range ports {
		err := br.AddPort(name)
		if err != nil {
			log.Fatalf("Cannot add port %s: %v", name, err)
		}
	}

	br.Run()
}

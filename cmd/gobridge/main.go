package main

import (
	"github.com/marokiki/gobridge/internal/bridge"
)

func main() {
	br := bridge.NewBridge()

	// ports := []string{"eth3", "eth4", "eth5", "eth6"}
	ports := []string{"eth1"}
	for _, name := range ports {
		br.AddPort(name)
	}

	br.Run()
}

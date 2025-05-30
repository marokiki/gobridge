package bridge

import (
	"fmt"

	"github.com/google/gopacket/layers"
)

type MacTable struct {
	macAddr string
	port    *Port
}

type Bridge struct {
	ports    []*Port
	macTable []MacTable
}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) AddPort(name string) {
	port := &Port{
		name:   name,
		bridge: b,
	}
	b.ports = append(b.ports, port)
}

func (b *Bridge) Run() {
	for _, port := range b.ports {
		go port.CapturePacket()
	}

	select {}
}

func (b *Bridge) LearnMac(packet *layers.Ethernet, port *Port) bool {
	for _, macTable := range b.macTable {
		if macTable.macAddr == packet.SrcMAC.String() {
			return false
		}
	}

	b.macTable = append(b.macTable, MacTable{
		macAddr: packet.SrcMAC.String(),
		port:    port,
	})
	fmt.Println(b.macTable)
	return true
}

func (b *Bridge) SendPacket(packet *layers.Ethernet) *Port {
	for _, macTable := range b.macTable {
		if macTable.macAddr == packet.DstMAC.String() {
			return macTable.port
		}
	}
	return nil
}

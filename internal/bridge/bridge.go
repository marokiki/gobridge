package bridge

import "github.com/google/gopacket/layers"

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
	port := &Port{name: name}
	b.ports = append(b.ports, port)
}

func (b *Bridge) Run() {
	for _, port := range b.ports {
		go port.CapturePacket()
	}

	select {}
}

func (b *Bridge) SendPacket(packet *layers.Ethernet) *Port {
	for _, macTable := range b.macTable {
		if macTable.macAddr == packet.DstMAC.String() {
			return macTable.port
		}
	}
	return nil
}

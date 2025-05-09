package bridge

import (
	"github.com/marokiki/gobridge/internal/packet"
)

type Port struct {
	name string
}

func (p *Port) CapturePacket() {
	packet.CapturePacket(p.name)
}

// func LearnMac(ethernetPacket *layers.Ethernet) {

// }

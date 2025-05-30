package bridge

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/marokiki/gobridge/internal/packet"
)

type Port struct {
	name   string
	bridge *Bridge
}

func (p *Port) CapturePacket() {
	handle := packet.OpenPort(p.name)
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if ethernetLayer := packet.Layer(layers.LayerTypeEthernet); ethernetLayer != nil {
			ethernetPacket, ok := ethernetLayer.(*layers.Ethernet)
			if !ok {
				continue
			}
			fmt.Println(ethernetPacket.SrcMAC.String())
			fmt.Println(ethernetPacket.DstMAC.String())
			fmt.Println(ethernetPacket.EthernetType.String())
			if p.bridge.LearnMac(ethernetPacket, p) {
				// flood packet
			}
		}
	}
}

package packet

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/marokiki/gobridge/internal/bridge"
)

func CapturePacket(port string) *layers.Ethernet {
	handle, err := pcap.OpenLive("eth1", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if ethernetLayer := packet.Layer(layers.LayerTypeEthernet); ethernetLayer != nil {
			ethernetPacket, ok := ethernetLayer.(*layers.Ethernet)
			if !ok {
				log.Fatal("Failed to cast to Ethernet")
			}
			fmt.Println(ethernetPacket)
		}
	}
}

// func DecodePacket(packet *layers.Ethernet) (*Ethernet, error) {
// 	ethernet := &packet
// }

package packet

import (
	"log"

	"github.com/google/gopacket/pcap"
)

func OpenPort(port string) *pcap.Handle {
	handle, err := pcap.OpenLive(port, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	return handle
}

// func DecodePacket(packet *layers.Ethernet) (*Ethernet, error) {
// 	ethernet := &packet
// }

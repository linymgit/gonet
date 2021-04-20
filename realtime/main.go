package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

var (
	device            = "\\Device\\NPF_{8E6EF140-3A55-473F-AA23-6B7879689E08}"
	snapshotLen int32 = 1024
	timeout           = 30 * time.Second
	handle      pcap.Handle
)

func main() {
	// Open device
	livehandle, err := pcap.OpenLive(device, snapshotLen, false, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(livehandle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}

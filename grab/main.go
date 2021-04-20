package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"os"
	"time"
)

var (
	deviceName  string = "\\Device\\NPF_{8E6EF140-3A55-473F-AA23-6B7879689E08}"
	snapshotLen int32  = 1024
	timeout            = -1 * time.Second
	packetCount        = 0
)

func main() {
	// Open output pcap file and write header
	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, snapshotLen, false, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		err := w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		if err != nil {
			log.Print(err)
		}
		packetCount++
		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}

}

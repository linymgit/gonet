package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

// Use tcpdump to create a test file
// tcpdump -w test.pcap
// or use the example above for writing pcap files
func main() {
	// Open file instead of device
	handle, err := pcap.OpenOffline("E:\\awesomeProject\\gonet\\test.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}

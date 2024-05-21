package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"testing"
)

func TestName(t *testing.T) {
	ifs, _ := pcap.FindAllDevs()
	for _, i := range ifs {
		fmt.Printf("Device Name: %+v\n", i)
	}
}

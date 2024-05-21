package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
)

func main() {
	LoadConfig("config.json")

	devs, _ := pcap.FindAllDevs()

	// deviceName有可能处理不准确，可通过实际情况手动选择或者修改以下代码来获取自身电脑正确的deviceName
	deviceName := ""
	for _, dev := range devs {
		if dev.Flags == 22 && len(dev.Addresses) > 0 {
			deviceName = dev.Name
		}
	}
	//////////////////////////////////////////////////////////////////

	handle, err := pcap.OpenLive(deviceName, 65536, true, pcap.BlockForever)

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	go listenARP(handle, true)

	time.Sleep(1 * time.Second) // Give the listener some time to start before sending

	err = sendARP(handle, cfg.SrcIp, cfg.SrcMac, "192.168.3.9", "ff:ff:ff:ff:ff:ff")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ARP packet sent")
	time.Sleep(2 * time.Second) // Wait for ARP response
	for {

	}
}
func listenARP(handle *pcap.Handle, isFilter bool) {
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		output := fmt.Sprintf("\n 时间 = [%v]", packet.Metadata().Timestamp)

		// 链路层
		if packet.LinkLayer() != nil {
			linkLayer := packet.LinkLayer()
			switch linkLayer.(type) {
			case *layers.Ethernet:
				macLayer := linkLayer.(*layers.Ethernet)

				if isFilter {
					if macLayer.SrcMAC.String() == cfg.SrcMac || macLayer.DstMAC.String() == cfg.SrcMac {
						output += fmt.Sprintf(" 链路层 = [SrcMAC=%v, DstMAC=%v]", macLayer.SrcMAC, macLayer.DstMAC)
					} else {
						continue
					}
				} else {
					output += fmt.Sprintf(" 链路层 = [SrcMAC=%v, DstMAC=%v]", macLayer.SrcMAC, macLayer.DstMAC)
				}

			default:
				output += fmt.Sprintf(" 链路层 = [packet.NetworkLayer=%+v]", linkLayer)
			}
		}

		// 网络层
		if packet.NetworkLayer() != nil {
			netWorkLayer := packet.NetworkLayer()
			switch netWorkLayer.(type) {
			case *layers.IPv4:
				netLayer := netWorkLayer.(*layers.IPv4)
				output += fmt.Sprintf(" 网络层 = [SrcIP=%v, DstIP=%v]", netLayer.SrcIP, netLayer.DstIP)
			case *layers.IPv6:
				netLayer := netWorkLayer.(*layers.IPv6)
				output += fmt.Sprintf(" 网络层 = [SrcIP=%v, DstIP=%v]", netLayer.SrcIP, netLayer.DstIP)
			default:
				output += fmt.Sprintf(" 网络层 = [packet.NetworkLayer=%+v]", netWorkLayer)
			}
		}

		// 传输层
		if packet.TransportLayer() != nil {
			transPortLayer := packet.TransportLayer()
			switch transPortLayer.(type) {
			case *layers.TCP:
				tcpLayer := transPortLayer.(*layers.TCP)
				output += fmt.Sprintf(" 传输层 = [SrcPort=%v, DstPort=%v]", tcpLayer.SrcPort, tcpLayer.DstPort)
			case *layers.UDP:
				udpLayer := transPortLayer.(*layers.UDP)
				output += fmt.Sprintf(" 传输层 = [SrcPort=%v, DstPort=%v]", udpLayer.SrcPort, udpLayer.DstPort)
			default:
				output += fmt.Sprintf(" 传输层 = [packet.TransportLayer=%+v]", transPortLayer)
			}
		}

		// 应用层
		if packet.ApplicationLayer() != nil {
			//output += fmt.Sprintf(" 应用层 = [packet.ApplicationLayer=%s]", packet.ApplicationLayer())
			applicationLayer := packet.ApplicationLayer()
			switch applicationLayer.(type) {
			case *gopacket.Fragment:
				appLayer := applicationLayer.(*gopacket.Fragment)
				output += fmt.Sprintf(" 应用层- = [%s]", string(appLayer.Payload()))
			case *gopacket.Payload:
				appLayer := applicationLayer.(*gopacket.Payload)

				//layerClass := appLayer.CanDecode()
				//for _, val := range layerClass.LayerTypes() {
				//
				//}
				//output += fmt.Sprintf(" 应用层+ = [%+v]", layerClass)

				//nf := gopacket.DecodingLayerParser{}
				//err := appLayer.DecodeFromBytes(appLayer.LayerContents(), &nf)
				//if err != nil {
				//	output += fmt.Sprintf(" 应用层+ = [DecodeFromBytes err: %s]", err)
				//} else {
				//	output += fmt.Sprintf(" 应用层+ = [%+v]", nf)
				//}
				//output += fmt.Sprintf(" 应用层+ = [%s]", appLayer.CanDecode().LayerTypes())
				output += fmt.Sprintf(" 应用层++ = [%#v, ====, %v]", appLayer.LayerContents(), appLayer.LayerContents())

			default:
				output += fmt.Sprintf(" 应用层 = [packet.TransportLayer=%+v]", applicationLayer)
			}
		}

		fmt.Println(output)

	}
}

func listenARPBak(handle *pcap.Handle) {
	//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	//for packet := range packetSource.Packets() {
	//	fmt.Println("\n==============================================================================================================")
	//	//fmt.Printf("packet.Metadata=%+v\n", packet.Metadata())
	//	/**
	//	packet.Metadata=&{CaptureInfo:{Timestamp:2024-05-18 18:19:58.727321 +0800 CST CaptureLength:60 Length:60 InterfaceIndex:0 AncillaryData:[]} Truncated:false}
	//	*/
	//
	//	/**
	//	fmt.Printf("packet.dump=%s\n", packet.Dump())
	//
	//	packet.dump=-- FULL PACKET DATA (60 bytes) ------------------------------------
	//	00000000  00 0c 29 01 51 22 a4 3b  0e 83 49 1f 08 00 45 00  |..).Q".;..I...E.|
	//	00000010  00 28 00 00 40 00 34 06  18 fc 27 9c 42 0e c0 a8  |.(..@.4...'.B...|
	//	00000020  03 82 00 50 d4 c9 e4 af  30 00 00 00 00 00 50 04  |...P....0.....P.|
	//	00000030  00 00 98 42 00 00 00 00  00 00 00 00              |...B........|
	//	--- Layer 1 ---
	//	Ethernet        {Contents=[..14..] Payload=[..46..] SrcMAC=a4:3b:0e:83:49:1f DstMAC=00:0c:29:01:51:22 EthernetType=IPv4 Length=0}
	//	00000000  00 0c 29 01 51 22 a4 3b  0e 83 49 1f 08 00        |..).Q".;..I...|
	//	--- Layer 2 ---
	//	IPv4    {Contents=[..20..] Payload=[..20..] Version=4 IHL=5 TOS=0 Length=40 Id=0 Flags=DF FragOffset=0 TTL=52 Protocol=TCP Checksum=6396 SrcIP=39.156.66.14 DstIP=192.168.3.130 Options=[] Padding=[]}
	//	00000000  45 00 00 28 00 00 40 00  34 06 18 fc 27 9c 42 0e  |E..(..@.4...'.B.|
	//	00000010  c0 a8 03 82                                       |....|
	//	--- Layer 3 ---
	//	TCP     {Contents=[..20..] Payload=[] SrcPort=80(http) DstPort=54473 Seq=3836686336 Ack=0 DataOffset=5 FIN=false SYN=false RST=true PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=0 Checksum=38978 Urgent=0 Options=[] Padding=[]}
	//	00000000  00 50 d4 c9 e4 af 30 00  00 00 00 00 50 04 00 00  |.P....0.....P...|
	//	00000010  98 42 00 00                                       |.B..|
	//
	//	*/
	//
	//	if packet.LinkLayer() != nil {
	//		if macLayer := packet.LinkLayer().(*layers.Ethernet); macLayer != nil {
	//			// 链路层 = layerType=Ethernet, EthernetType=IPv4, SrcMAC=c8:5e:a9:c4:54:dc, DstMAC=a4:3b:0e:83:49:1f, NextLayerType=IPv4
	//			fmt.Println(fmt.Sprintf("链路层 = layerType=%v, EthernetType=%s, SrcMAC=%v, DstMAC=%v, NextLayerType=%v", macLayer.LayerType(), macLayer.EthernetType, macLayer.SrcMAC, macLayer.DstMAC, macLayer.NextLayerType()))
	//		}
	//	}
	//
	//	if packet.NetworkLayer() != nil {
	//		if netLayer := packet.NetworkLayer().(*layers.IPv4); netLayer != nil {
	//			// 网络层 = layerType=IPv4, Protocol=TCP version=4, SrcMAC=192.168.3.9, DstMAC=59.110.73.151, NextLayerType=TCP
	//			fmt.Println(fmt.Sprintf("网络层 = layerType=%v, Protocol=%v version=%d, SrcMAC=%v, DstMAC=%v, NextLayerType=%v", netLayer.LayerType(), netLayer.Protocol, netLayer.Version, netLayer.SrcIP, netLayer.DstIP, netLayer.NextLayerType()))
	//		}
	//	}
	//
	//	if packet.TransportLayer() != nil {
	//		if tcpLayer := packet.TransportLayer().(*layers.TCP); tcpLayer != nil {
	//			// 传输层 = layerType=TCP,  SrcMAC=52966, DstMAC=80(http), NextLayerType=Payload, tcpLayer=&{BaseLay......
	//			fmt.Println(fmt.Sprintf("传输层 = layerType=%v,  SrcMAC=%v, DstMAC=%v, NextLayerType=%v, tcpLayer=%+v", tcpLayer.LayerType(), tcpLayer.SrcPort, tcpLayer.DstPort, tcpLayer.NextLayerType(), tcpLayer))
	//		}
	//	}
	//
	//	fmt.Printf("\npacket.ApplicationLayer=%+v\n", packet.ApplicationLayer())
	//
	//	//for _, layer := range packet.Layers() {
	//	//	if layer.LayerType() == layers.LayerTypeARP {
	//	//		fmt.Println(fmt.Sprintf("arptyp = layerType=%v, palyLoad=%s, contents=%+v", layer.LayerType(), layer.LayerPayload(), layer.LayerContents()))
	//	//	} else if layer.LayerType() == layers.LayerTypeEthernet {
	//	//
	//	//		fmt.Println(fmt.Sprintf("ethernet = layerType=%v, palyLoad=%s, contents=%+v", layer.LayerType(), layer.LayerPayload(), layer.LayerContents()))
	//	//	}
	//	//}
	//	//if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
	//	//	arpPacket, _ := arpLayer.(*layers.ARP)
	//	//	if arpPacket.Operation == layers.ARPReply {
	//	//		fmt.Println("Received ARP response:")
	//	//		fmt.Printf("Sender MAC: %s\n", arpPacket.SourceHwAddress)
	//	//		fmt.Printf("Sender IP: %s\n", arpPacket.SourceProtAddress)
	//	//	}
	//	//}
	//}
}

func sendARP(handle *pcap.Handle, srcIP, srcMAC, targetIP, targetMAC string) error {
	srcHardwareMac, _ := net.ParseMAC(srcMAC)
	dstHardwareMac, _ := net.ParseMAC(targetMAC)
	arpRequest := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   srcHardwareMac,
		SourceProtAddress: net.ParseIP(srcIP).To4(),
		DstHwAddress:      dstHardwareMac,
		DstProtAddress:    net.ParseIP(targetIP).To4(),
		//DstHwAddress:      []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		//DstProtAddress:    []byte{0, 0, 0, 0, 0, 0, 0, 0}, // 目标IP设为广播地址，以获取所有设备回应
	}
	ethernetPacket := layers.Ethernet{
		SrcMAC:       srcHardwareMac,
		DstMAC:       dstHardwareMac, // Broadcast MAC
		EthernetType: layers.EthernetTypeARP,
	}
	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, gopacket.SerializeOptions{},
		&ethernetPacket,
		&arpRequest,
	)
	outgoingPacket := buffer.Bytes()
	err := handle.WritePacketData(outgoingPacket)
	if err != nil {
		return err
	}
	return nil
}

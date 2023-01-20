package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func init() {

}

var (
	dataInterface       = "enx00e04c01137b"
	snapLen       int32 = 262144
)

func startDeviceListener() {
	getLogger().Info("Listening for devices on interface...")
	handle, err := pcap.OpenLive(dataInterface, snapLen, true,
		pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		handlePacket(pkt)
	}

}

func handlePacket(packet gopacket.Packet) {

	// Do nothing if packet is not VLAN tagged
	if packet.Layer(layers.LayerTypeDot1Q) == nil {
		return
	}

	// Packet belongs to a device under test
	taggedPacketEthernetLayer := packet.Layer(layers.LayerTypeEthernet)
	ethernetPacket := taggedPacketEthernetLayer.(*layers.Ethernet)

	taggedPacketVlanLayer := packet.Layer(layers.LayerTypeDot1Q)
	vlanPacket, _ := taggedPacketVlanLayer.(*layers.Dot1Q)
	vlanId := vlanPacket.VLANIdentifier

	// TODO: Check if a device has already been registered with this vlan
	handleDeviceLearn(vlanId, ethernetPacket.SrcMAC.String())

}

func handleDeviceLearn(vlanId uint16, macAddr string) {
	getLogger().Info("Device ", macAddr, " found on port ", vlanId-100)
}

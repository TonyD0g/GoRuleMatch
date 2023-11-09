package Core

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"os"
	"strconv"
)

func GetHttpData(port int) {
	if port <= 0 || port > 65536 {
		fmt.Println("[-] 端口超出范围")
		os.Exit(1)
	}

	// 打开网络接口
	handle, err := pcap.OpenLive("\\Device\\NPF_{F99942B2-D22D-4DF0-9A28-17F3BBA78D3F}", 65536, true, pcap.BlockForever)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer handle.Close()

	// 设置过滤器，只捕获HTTP数据包
	httpFilter := "tcp and port " + strconv.Itoa(port)
	if err := handle.SetBPFFilter(httpFilter); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 捕获并处理数据包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println(packet.Dump())
	}
}

// PrintDevice 调用 FindAllDevs() 函数获取系统上的网络适配器列表
func PrintDevice() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Printf("Error finding devices: %v\n", err)
		return
	}

	// 遍历设备列表并打印每个设备的信息
	for _, device := range devices {
		fmt.Printf("Name: %s\n", device.Name)
		fmt.Printf("Description: %s\n", device.Description)
		fmt.Printf("Devices addresses: %s\n", device.Addresses)
		fmt.Println("--------------------------")
	}
}

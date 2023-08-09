package server

import (
	"fmt"
	"github.com/jaypipes/ghw"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func sendAgentRegister() {
	workdir, _ := os.Getwd()     //工作目录
	hostname, _ := os.Hostname() //主机名
	// 处理器
	var cpus []string
	cpu, _ := ghw.CPU()
	for _, processor := range cpu.Processors {
		cpus = append(cpus, processor.Model)
	}
	//内存
	memory, _ := ghw.Memory()
	//显卡
	var gpus []string
	gpu, _ := ghw.GPU()
	for _, card := range gpu.GraphicsCards {
		if strings.HasPrefix(card.Address, "PCI") {
			gpus = append(gpus, card.DeviceInfo.Product.VendorID)
		}
	}
	//用户
	currentUser, _ := user.Current()
	username := currentUser.Username
	//网络ip
	var ips []string
	network, _ := ghw.Network()
	var NetCardNames []string
	for _, c := range network.NICs {
		if !c.IsVirtual {
			NetCardNames = append(NetCardNames, strings.Split(c.Name, " - ")[0])
		}
	}
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if ContainsString(NetCardNames, iface.Name) {
			if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
				addrs, _ := iface.Addrs()
				for _, addr := range addrs {
					ip, _, _ := net.ParseCIDR(addr.String())
					if !ip.IsLinkLocalUnicast() {
						ips = append(ips, ip.String())
					}
				}
			}
		}
	}
	agentInfo := AgentInfo{
		Workdir:    workdir,
		Hostname:   hostname,
		Memory:     memory.String(),
		Cpus:       cpus,
		Ips:        ips,
		OsArch:     runtime.GOARCH,
		OsType:     runtime.GOOS,
		SystemUser: username,
	}
	fmt.Printf("%v\n", &agentInfo)
}

func ContainsString(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

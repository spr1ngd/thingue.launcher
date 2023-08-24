package instance

import (
	"github.com/jaypipes/ghw"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/util"
)

func GetDeviceInfo() *domain.DeviceInfo {
	workdir, _ := os.Getwd()     //工作目录
	hostname, _ := os.Hostname() //主机名
	// 处理器
	var cpus []string
	cpu, _ := ghw.CPU()
	if cpu != nil {
		for _, processor := range cpu.Processors {
			cpus = append(cpus, processor.Model)
		}
	}
	//内存
	var memoryMsg string
	memory, _ := ghw.Memory()
	if memory != nil {
		memoryMsg = memory.String()
	}
	//显卡
	var gpus []string
	gpu, _ := ghw.GPU()
	if gpu != nil {
		for _, card := range gpu.GraphicsCards {
			if strings.HasPrefix(card.DeviceInfo.Address, "PCI") {
				gpus = append(gpus, card.DeviceInfo.Product.Name)
			}
		}
	}
	//系统用户
	currentUser, _ := user.Current()
	username := currentUser.Username
	//网络ip
	var ips []string
	network, _ := ghw.Network()
	var ifaceNames []string
	if network != nil {
		for _, c := range network.NICs {
			if !c.IsVirtual {
				ifaceNames = append(ifaceNames, strings.Split(c.Name, " - ")[0])
			}
		}
	}
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if util.ContainsString(ifaceNames, iface.Name) {
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
	return &domain.DeviceInfo{
		Version:    global.APP_VERSION,
		Workdir:    workdir,
		Hostname:   hostname,
		Memory:     memoryMsg,
		Cpus:       cpus,
		Gpus:       gpus,
		Ips:        ips,
		OsArch:     runtime.GOARCH,
		OsType:     runtime.GOOS,
		SystemUser: username,
	}
}

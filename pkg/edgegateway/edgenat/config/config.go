package config

import (
	"fmt"
	"k8s.io/klog"
	"net"
	"sync"
)

var Config Configure
var once sync.Once

const (
	Interface     = "docker0"
	ListenPort    = 4001
	DefaultSubnet = "192.168.0.0/16"
)

type Configure struct {
	Listener        *net.TCPListener
	ListenInterface string
	SubNet          string
}

func InitConfigure() {
	once.Do(func() {
		ListenIP, err := GetInterfaceIP(Interface)
		if err != nil {
			klog.Errorf("failed to get listen ip from interface %s, error: %v", Interface, err)
		}
		// get listener
		listenAddr := &net.TCPAddr{
			IP:   ListenIP,
			Port: ListenPort,
		}

		ln, err := net.ListenTCP("tcp", listenAddr)
		if err != nil {
			klog.Errorf("failed to get listener, error: %v", err)
		}
		Config.Listener = ln
		Config.ListenInterface = Interface
		Config.SubNet = DefaultSubnet
	})
}

func GetInterfaceIP(name string) (net.IP, error) {
	ifi, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}
	addrs, _ := ifi.Addrs()
	for _, addr := range addrs {
		if ip, ipn, _ := net.ParseCIDR(addr.String()); len(ipn.Mask) == 4 {
			return ip, nil
		}
	}
	return nil, fmt.Errorf("no ip of version 4 found for interface %s", name)
}

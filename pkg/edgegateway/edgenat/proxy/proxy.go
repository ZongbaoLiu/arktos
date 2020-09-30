package proxy

import (
	"github.com/vishvananda/netlink"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/edgegateway/edgenat/config"
	utildbus "k8s.io/kubernetes/pkg/util/dbus"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
	utilexec "k8s.io/utils/exec"
	"strings"
)

// iptables rules
type Proxier struct {
	iptables     utiliptables.Interface
	inboundRule  string
	outboundRule string
	dNatRule     string
	sNatRule     string
}

const (
	meshChain  = "EDGE-MESH"
	hostResolv = "/etc/resolv.conf"
)

var (
	proxier *Proxier
	route   netlink.Route
)

func init() {
	protocol := utiliptables.ProtocolIpv4
	exec := utilexec.New()
	dbus := utildbus.New()
	iptInterface := utiliptables.New(exec, dbus, protocol)
	proxier = &Proxier{
		iptables:     iptInterface,
		inboundRule:  "-p tcp -d " + config.Config.SubNet + " -i " + config.Config.ListenInterface + " -j " + meshChain,
		outboundRule: "-p tcp -d " + config.Config.SubNet + " -o " + config.Config.ListenInterface + " -j " + meshChain,
		dNatRule:     "-p tcp -j DNAT --to-destination " + config.Config.Listener.Addr().String(),
		sNatRule:     "iptables -t nat -A POSTROUTING -s 10.0.1.0/24 ! -d 10.0.1.0/24 -j SNAT --to-source 172.16.1.6:1111",
	}

	inboundRule := strings.Split(proxier.inboundRule, " ")
	outboundRule := strings.Split(proxier.outboundRule, " ")
	dNatRule := strings.Split(proxier.dNatRule, " ")
	sNatRule := strings.Split(proxier.sNatRule, " ")


	exist, err := proxier.iptables.EnsureRule(utiliptables.Append, utiliptables.TableNAT, utiliptables.ChainPrerouting, inboundRule...)
	if err != nil {
		klog.Errorf("")
	}
	if !exist {

	}
	exist, err = proxier.iptables.EnsureRule(utiliptables.Append, utiliptables.TableNAT, utiliptables.ChainOutput, outboundRule...)
	if err != nil {

	}
	if !exist {

	}
	exist, err = proxier.iptables.EnsureRule(utiliptables.Append, utiliptables.TableNAT, meshChain, dNatRule...)
	if err != nil {

	}
	if !exist {

	}
	exist, err = proxier.iptables.EnsureRule(utiliptables.Append, utiliptables.TableNAT, meshChain, sNatRule...)
	if err != nil {

	}
	if !exist {

	}
}

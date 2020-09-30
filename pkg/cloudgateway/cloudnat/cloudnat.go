package cloudnat

import (
	"fmt"
	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/viaduct/pkg/api"
	"github.com/kubeedge/viaduct/pkg/conn"
	"github.com/kubeedge/viaduct/pkg/server"
	"k8s.io/klog"
)

type cloudNat struct {
	//listener *net.TCPListener
	enable bool
}

func newCloudNat(enable bool) *cloudNat {
	return &cloudNat{
		enable: enable,
	}
}

func Register(enable bool) {
	//config.InitConfigure()
	core.Register(newCloudNat(enable))
}

func (n *cloudNat) Name() string {
	//return modules.CloudNatModuleName
	return "cloudNat"
}

func (n *cloudNat) Group() string {
	//return modules.CloudNatGroup
	return "cloudNat"
}

func (n *cloudNat) Enable() bool {
	return n.enable
}

func (n *cloudNat) Start() {
	svc := server.Server{
		Type:       api.ProtocolTypeWS,
		AutoRoute:  true,
		ConnNotify: process,
		Addr:       fmt.Sprintf("%s:%d", "192.168.10.240", 1111),
		ExOpts:     api.WSServerOption{Path: "/"},
	}
	klog.Infof("Starting cloudNat %s server", api.ProtocolTypeWS)
	klog.Fatal(svc.ListenAndServeTLS("", ""))
}

func process(connection conn.Connection) {
	//tlsConfig := createTLSConfig(hubconfig.Config.Ca, hubconfig.Config.Cert, hubconfig.Config.Key)
	buf := make([]byte, 2048)
	num, err := connection.Read(buf)
	if err != nil {
		return
	}
	rawData := buf[:num]
	klog.Infof("get message stream from cloudNat: %s", string(rawData))
}

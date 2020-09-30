package edgenat

import (
	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/viaduct/pkg/api"
	wsclient "github.com/kubeedge/viaduct/pkg/client"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/edgegateway/edgenat/config"
	"k8s.io/kubernetes/pkg/edgegateway/edgenat/listener"
	"net"
	"strings"
)

type edgeNat struct {
	listener *net.TCPListener
	enable   bool
}

func newEdgeNat(enable bool) *edgeNat {
	return &edgeNat{
		enable: enable,
	}
}

func Register(enable bool) {
	config.InitConfigure()
	core.Register(newEdgeNat(enable))
}

func (n *edgeNat) Name() string {
	//return modules.EdgeNatModuleName
	return "edgenat"
}

func (n *edgeNat) Group() string {
	//return modules.EdgeNatGroup
	return "edgenat"
}

func (n *edgeNat) Enable() bool {
	return n.enable
}

func (n *edgeNat) Start() {
	//var listener *net.TCPListener
	//for {
	//	listener, err := net.Listen("tcp", "127.0.0.1:4001")
	//	if err == nil {
	//		break
	//	}
	//
	//}
	//
	//conn, err := listener.Accept()
	//if err != nil {
	//
	//}
	//ip, port, err := realServerAddress(&conn)
	//go proxy.init()
	//go listener.Start()

	WebSocketURL := strings.Join([]string{"wss:/", "192.168.10.240:1111", "eh.SiteID", "events"}, "/")
	option := wsclient.Options{
		Type:      api.ProtocolTypeWS,
		Addr:      WebSocketURL,
		AutoRoute: false,
		ConnUse:   api.UseTypeStream,
	}
	client := wsclient.Client{Options: option}
	connection, err := client.Connect()
	if err != nil {
		klog.Errorf("failed to connect websocket, err: %v", err)
	}

	conn, err := config.Config.Listener.Accept()
	if err != nil {
		return
	}

	ip, port, err := listener.RealServerAddress(&conn)
	if err != nil {
		klog.Errorf("failed to get real server address, err: %v", err)
		return
	}
	klog.Infof("get real ip is %s, port is %d", ip, port)

	buf := make([]byte, 2048)
	num, err := conn.Read(buf)
	if err != nil {
		return
	}

	rawData := buf[:num]

	connection.Write(rawData)
	//connection.Read(buf[:num])


	//for {
	//	select {
	//	case <-beehiveContext.Done():
	//		klog.Warning("edgeNat stop")
	//		return
	//	default:
	//	}
	//	msg, err := beehiveContext.Receive(modules.EdgeNatModuleName)
	//	if err != nil {
	//		klog.Warningf("%s receive message error: %v", modules.EdgeNatModuleName, err)
	//		continue
	//	}
	//	// TODO(liuzongbao): route msg stream to edge service
	//	println(msg)
	//}
}

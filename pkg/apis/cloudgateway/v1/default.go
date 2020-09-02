package v1

import (
	"k8s.io/kubernetes/pkg/cloudgateway/common/constants"
	"path"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/kubernetes/pkg/apis/cloudgateway"
)

// NewCloudGatewayConfig returns a full CloudGatewayConfig object
func NewCloudGatewayConfig() *CloudGatewayConfig {
	advertiseAddress, _ := utilnet.ChooseHostInterface()

	c := &CloudGatewayConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       cloudgateway.Kind,
			APIVersion: path.Join(cloudgateway.GroupName, cloudgateway.Version),
		},
		Modules: &Modules{
			CloudHub: &CloudHub{
				Enable:                  true,
				KeepaliveInterval:       30,
				NodeLimit:               1000,
				TLSCAFile:               constants.DefaultCAFile,
				TLSCAKeyFile:            constants.DefaultCAKeyFile,
				TLSCertFile:             constants.DefaultCertFile,
				TLSPrivateKeyFile:       constants.DefaultKeyFile,
				WriteTimeout:            30,
				AdvertiseAddress:        []string{advertiseAddress.String()},
				EdgeCertSigningDuration: 365,
				Quic: &CloudHubQUIC{
					Enable:             false,
					Address:            "0.0.0.0",
					Port:               10001,
					MaxIncomingStreams: 10000,
				},
				WebSocket: &CloudHubWebSocket{
					Enable:  true,
					Port:    10000,
					Address: "0.0.0.0",
				},
				HTTPS: &CloudHubHTTPS{
					Enable:  true,
					Port:    10002,
					Address: "0.0.0.0",
				},
			},
		},
	}
	return c
}

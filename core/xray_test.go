package core_test

import (
	"testing"

	"github.com/decardlabs/uuhappyserver/app/dispatcher"
	"github.com/decardlabs/uuhappyserver/app/proxyman"
	"github.com/decardlabs/uuhappyserver/common"
	"github.com/decardlabs/uuhappyserver/common/net"
	"github.com/decardlabs/uuhappyserver/common/protocol"
	"github.com/decardlabs/uuhappyserver/common/serial"
	"github.com/decardlabs/uuhappyserver/common/uuid"
	. "github.com/decardlabs/uuhappyserver/core"
	"github.com/decardlabs/uuhappyserver/features/dns"
	"github.com/decardlabs/uuhappyserver/features/dns/localdns"
	_ "github.com/decardlabs/uuhappyserver/main/distro/all"
	"github.com/decardlabs/uuhappyserver/proxy/dokodemo"
	"github.com/decardlabs/uuhappyserver/proxy/vmess"
	"github.com/decardlabs/uuhappyserver/proxy/vmess/outbound"
	"github.com/decardlabs/uuhappyserver/testing/servers/tcp"
	"google.golang.org/protobuf/proto"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	}, false)
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(net.LocalHostIP),
					Port:     uint32(0),
					Networks: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: &protocol.ServerEndpoint{
						Address: net.NewIPOrDomain(net.LocalHostIP),
						Port:    uint32(0),
						User:  &protocol.User{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: userID.String(),
							}),
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}

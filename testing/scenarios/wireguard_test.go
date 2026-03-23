package scenarios

import (
	"testing"
	//"time"

	"github.com/decardlabs/uuhappyserver/app/log"
	"github.com/decardlabs/uuhappyserver/app/proxyman"
	"github.com/decardlabs/uuhappyserver/common"
	clog "github.com/decardlabs/uuhappyserver/common/log"
	"github.com/decardlabs/uuhappyserver/common/net"
	"github.com/decardlabs/uuhappyserver/common/serial"
	core "github.com/decardlabs/uuhappyserver/core"
	"github.com/decardlabs/uuhappyserver/infra/conf"
	"github.com/decardlabs/uuhappyserver/proxy/dokodemo"
	"github.com/decardlabs/uuhappyserver/proxy/freedom"
	"github.com/decardlabs/uuhappyserver/proxy/wireguard"
	"github.com/decardlabs/uuhappyserver/testing/servers/tcp"
	"github.com/decardlabs/uuhappyserver/testing/servers/udp"
	//"golang.org/x/sync/errgroup"
)

func TestWireguard(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	serverPrivate, _ := conf.ParseWireGuardKey("EGs4lTSJPmgELx6YiJAmPR2meWi6bY+e9rTdCipSj10=")
	serverPublic, _ := conf.ParseWireGuardKey("osAMIyil18HeZXGGBDC9KpZoM+L2iGyXWVSYivuM9B0=")
	clientPrivate, _ := conf.ParseWireGuardKey("CPQSpgxgdQRZa5SUbT3HLv+mmDVHLW5YR/rQlzum/2I=")
	clientPublic, _ := conf.ParseWireGuardKey("MmLJ5iHFVVBp7VsB0hxfpQ0wEzAbT2KQnpQpj0+RtBw=")

	serverPort := udp.PickPort()
	serverConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(serverPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&wireguard.DeviceConfig{
					IsClient:    false,
					NoKernelTun: false,
					Endpoint:    []string{"10.0.0.1"},
					Mtu:         1420,
					SecretKey:   serverPrivate,
					Peers: []*wireguard.PeerConfig{{
						PublicKey:  serverPublic,
						AllowedIps: []string{"0.0.0.0/0", "::0/0"},
					}},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := tcp.PickPort()
	clientConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(clientPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(dest.Address),
					Port:     uint32(dest.Port),
					Networks: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&wireguard.DeviceConfig{
					IsClient:    true,
					NoKernelTun: false,
					Endpoint:    []string{"10.0.0.2"},
					Mtu:         1420,
					SecretKey:   clientPrivate,
					Peers: []*wireguard.PeerConfig{{
						Endpoint:   "127.0.0.1:" + serverPort.String(),
						PublicKey:  clientPublic,
						AllowedIps: []string{"0.0.0.0/0", "::0/0"},
					}},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	// FIXME: for some reason wg server does not receive

	// var errg errgroup.Group
	// for i := 0; i < 1; i++ {
	// 	errg.Go(testTCPConn(clientPort, 1024, time.Second*2))
	// }
	// if err := errg.Wait(); err != nil {
	// 	t.Error(err)
	// }
}

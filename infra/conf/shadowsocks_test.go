package conf_test

import (
	"testing"

	"github.com/decardlabs/uuhappyserver/common/net"
	"github.com/decardlabs/uuhappyserver/common/protocol"
	"github.com/decardlabs/uuhappyserver/common/serial"
	. "github.com/decardlabs/uuhappyserver/infra/conf"
	"github.com/decardlabs/uuhappyserver/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "xray-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				Users: []*protocol.User{{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "xray-password",
					}),
				}},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}

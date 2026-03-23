package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/decardlabs/uuhappyserver/app/dispatcher"
	_ "github.com/decardlabs/uuhappyserver/app/proxyman/inbound"
	_ "github.com/decardlabs/uuhappyserver/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/decardlabs/uuhappyserver/app/commander"
	_ "github.com/decardlabs/uuhappyserver/app/log/command"
	_ "github.com/decardlabs/uuhappyserver/app/proxyman/command"
	_ "github.com/decardlabs/uuhappyserver/app/stats/command"

	// Developer preview services
	_ "github.com/decardlabs/uuhappyserver/app/observatory/command"

	// Other optional features.
	_ "github.com/decardlabs/uuhappyserver/app/dns"
	_ "github.com/decardlabs/uuhappyserver/app/dns/fakedns"
	_ "github.com/decardlabs/uuhappyserver/app/log"
	_ "github.com/decardlabs/uuhappyserver/app/metrics"
	_ "github.com/decardlabs/uuhappyserver/app/policy"
	_ "github.com/decardlabs/uuhappyserver/app/reverse"
	_ "github.com/decardlabs/uuhappyserver/app/router"
	_ "github.com/decardlabs/uuhappyserver/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/decardlabs/uuhappyserver/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/decardlabs/uuhappyserver/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/decardlabs/uuhappyserver/proxy/blackhole"
	_ "github.com/decardlabs/uuhappyserver/proxy/dns"
	_ "github.com/decardlabs/uuhappyserver/proxy/dokodemo"
	_ "github.com/decardlabs/uuhappyserver/proxy/freedom"
	_ "github.com/decardlabs/uuhappyserver/proxy/http"
	_ "github.com/decardlabs/uuhappyserver/proxy/loopback"
	_ "github.com/decardlabs/uuhappyserver/proxy/shadowsocks"
	_ "github.com/decardlabs/uuhappyserver/proxy/socks"
	_ "github.com/decardlabs/uuhappyserver/proxy/trojan"
	_ "github.com/decardlabs/uuhappyserver/proxy/vless/inbound"
	_ "github.com/decardlabs/uuhappyserver/proxy/vless/outbound"
	_ "github.com/decardlabs/uuhappyserver/proxy/vmess/inbound"
	_ "github.com/decardlabs/uuhappyserver/proxy/vmess/outbound"
	_ "github.com/decardlabs/uuhappyserver/proxy/wireguard"

	// Transports
	_ "github.com/decardlabs/uuhappyserver/transport/internet/grpc"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/httpupgrade"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/kcp"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/reality"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/splithttp"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/tcp"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/tls"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/udp"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/websocket"

	// Transport headers
	_ "github.com/decardlabs/uuhappyserver/transport/internet/headers/http"
	_ "github.com/decardlabs/uuhappyserver/transport/internet/headers/noop"

	// JSON & TOML & YAML
	_ "github.com/decardlabs/uuhappyserver/main/json"
	_ "github.com/decardlabs/uuhappyserver/main/toml"
	_ "github.com/decardlabs/uuhappyserver/main/yaml"

	// Load config from file or http(s)
	_ "github.com/decardlabs/uuhappyserver/main/confloader/external"

	// Commands
	_ "github.com/decardlabs/uuhappyserver/main/commands/all"
)

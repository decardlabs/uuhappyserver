package udp

import (
	"github.com/decardlabs/uuhappyserver/common/buf"
	"github.com/decardlabs/uuhappyserver/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}

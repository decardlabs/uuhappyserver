package outbound

import (

	"github.com/decardlabs/uuhappyserver/common/net"
	"github.com/decardlabs/uuhappyserver/common/protocol"
)

// As a stub command consumer.
func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch cmd.(type) {
	default:
	}
}

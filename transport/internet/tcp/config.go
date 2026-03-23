package tcp

import (
	"github.com/decardlabs/uuhappyserver/common"
	"github.com/decardlabs/uuhappyserver/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}

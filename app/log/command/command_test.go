package command_test

import (
	"context"
	"testing"

	"github.com/decardlabs/uuhappyserver/app/dispatcher"
	"github.com/decardlabs/uuhappyserver/app/log"
	. "github.com/decardlabs/uuhappyserver/app/log/command"
	"github.com/decardlabs/uuhappyserver/app/proxyman"
	_ "github.com/decardlabs/uuhappyserver/app/proxyman/inbound"
	_ "github.com/decardlabs/uuhappyserver/app/proxyman/outbound"
	"github.com/decardlabs/uuhappyserver/common"
	"github.com/decardlabs/uuhappyserver/common/serial"
	"github.com/decardlabs/uuhappyserver/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}

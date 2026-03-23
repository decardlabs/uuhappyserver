package all

import (
	"github.com/decardlabs/uuhappyserver/main/commands/all/api"
	"github.com/decardlabs/uuhappyserver/main/commands/all/convert"
	"github.com/decardlabs/uuhappyserver/main/commands/all/tls"
	"github.com/decardlabs/uuhappyserver/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
		cmdMLDSA65,
		cmdMLKEM768,
		cmdVLESSEnc,
		cmdBuildMphCache,
	)
}

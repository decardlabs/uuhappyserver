package protocol // import "github.com/decardlabs/uuhappyserver/common/protocol"

import (
	"errors"
)

var ErrProtoNeedMoreData = errors.New("protocol matches, but need more data to complete sniffing")

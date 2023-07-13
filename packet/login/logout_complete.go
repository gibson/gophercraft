package login

import (
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type LogoutComplete struct {
}

func (l *LogoutComplete) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LOGOUT_COMPLETE
	return nil
}

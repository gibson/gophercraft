package login

import (
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/tempest"
	"github.com/gibson/gophercraft/vsn"
)

type VerifyWorld struct {
	MapID    uint32
	Position tempest.C4Vector
}

func (vw *VerifyWorld) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LOGIN_VERIFY_WORLD

	out.WriteUint32(vw.MapID)
	return vw.Position.Encode(out)
}

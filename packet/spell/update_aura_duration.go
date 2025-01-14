package spell

import (
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type AuraDurationUpdate struct {
	Slot     uint8
	Duration uint32
}

func (adu *AuraDurationUpdate) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_UPDATE_AURA_DURATION
	out.WriteByte(adu.Slot)
	out.WriteUint32(adu.Duration)
	return nil
}

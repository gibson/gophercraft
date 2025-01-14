package teleport

import (
	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/packet/update"
	"github.com/gibson/gophercraft/vsn"
)

type Ack struct {
	GUID guid.GUID

	Info *update.MovementInfo
}

func (ta *Ack) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.MSG_MOVE_TELEPORT_ACK
	ta.GUID.EncodePacked(build, out)
	out.WriteUint32(0)
	if err := update.EncodeMovementInfo(build, out.Buffer, ta.Info); err != nil {
		return err
	}
	return nil
}

func (ta *Ack) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	ta.GUID, err = guid.DecodePacked(build, in)
	if err != nil {
		return
	}
	in.ReadUint32()
	if ta.Info, err = update.DecodeMovementInfo(build, in); err != nil {
		return err
	}
	return nil
}

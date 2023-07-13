package update

import (
	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type SetActiveMover struct {
	ID guid.GUID
}

func (sam *SetActiveMover) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_SET_ACTIVE_MOVER
	err = sam.ID.EncodeUnpacked(build, out)
	return
}

func (sam *SetActiveMover) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	sam.ID, err = guid.DecodeUnpacked(build, in)
	return
}

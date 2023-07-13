package login

import (
	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type Player struct {
	Character guid.GUID
}

func (j *Player) Decode(build vsn.Build, in *packet.WorldPacket) error {
	var err error
	j.Character, err = guid.DecodeUnpacked(build, in)
	return err
}

func (j *Player) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_PLAYER_LOGIN
	j.Character.EncodeUnpacked(build, out)
	return nil
}

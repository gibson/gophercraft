package spell

import (
	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type PlayVisual struct {
	ID             guid.GUID
	SpellVisualKit uint32
}

func (p *PlayVisual) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_PLAY_SPELL_VISUAL
	if err := p.ID.EncodeUnpacked(build, out); err != nil {
		return err
	}
	out.WriteUint32(p.SpellVisualKit)
	return nil
}

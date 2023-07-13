package realm

import (
	"github.com/gibson/gophercraft/packet"
)

func (s *Session) SendAllAcheivementData() {
	p := packet.NewWorldPacket(packet.SMSG_ALL_ACHIEVEMENT_DATA)
	p.WriteUint32(0)
	s.SendPacket(p)
}

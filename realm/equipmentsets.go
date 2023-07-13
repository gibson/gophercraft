package realm

import "github.com/gibson/gophercraft/packet"

func (s *Session) SendEquipmentSetList() {
	p := packet.NewWorldPacket(packet.SMSG_EQUIPMENT_SET_LIST)
	p.WriteUint32(0)
	s.SendPacket(p)
}

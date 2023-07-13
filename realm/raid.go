package realm

import (
	"github.com/gibson/gophercraft/packet/raid"
)

func (s *Session) HandleRequestRaidInfo() {
	var info raid.InstanceInfo
	s.Send(&info)
}

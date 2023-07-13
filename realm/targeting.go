package realm

import (
	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet/update"
)

func (s *Session) GetTarget() guid.GUID {
	return s.Get("Target").GUID()
}

func (s *Session) Target(id guid.GUID) {
	s.SetGUID("Target", id)
	s.UpdatePlayer()
}

func (s *Session) HandleTarget(target *update.Target) {
	s.Target(target.ID)
}

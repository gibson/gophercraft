package commands

import "github.com/gibson/gophercraft/realm"

func cmdTeach(s *realm.Session, spellID uint32) {
	if err := s.LearnAbility(spellID); err != nil {
		s.Warnf("%s", err)
	}
}

package commands

import "github.com/gibson/gophercraft/realm"

func cmdModLevel(s *realm.Session, level int) {
	s.LevelUp(uint32(level))
}

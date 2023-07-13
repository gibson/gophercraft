package commands

import (
	"github.com/gibson/gophercraft/realm"
	"github.com/gibson/gophercraft/realm/wdb/models"
)

func cmdSetGamemode(s *realm.Session, gm models.GameMode) {
	s.SetGameMode(gm)
}

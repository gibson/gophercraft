package commands

import (
	"github.com/gibson/gophercraft/realm"
	"github.com/gibson/gophercraft/realm/wdb/models"
)

func cmdMoney(s *realm.Session, add models.Money) {
	s.AddMoney(add)
	s.Warnf("Added %s: new balance: %s", add, s.Char.Coinage)
}

package commands

import (
	"log"

	"github.com/gibson/gophercraft/realm"
)

func cmdMorph(s *realm.Session, displayID uint32) {
	log.Println("Morphing to ", displayID)

	s.SetUint32("DisplayID", displayID)
	s.UpdatePlayer()
}

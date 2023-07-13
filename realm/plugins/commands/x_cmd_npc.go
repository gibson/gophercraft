package commands

import (
	"github.com/gibson/gophercraft/realm"
	"github.com/gibson/gophercraft/realm/wdb"
	"github.com/gibson/gophercraft/realm/wdb/models"
)

func cmdAddNPC(s *realm.Session, npcID string) {
	var cr *models.CreatureTemplate
	s.DB().Lookup(wdb.BucketKeyStringID, npcID, &cr)
	if cr == nil {
		s.Warnf("No CreatureTemplate could be found with the ID %s", npcID)
		return
	}

	creature := s.Server.NewCreature(cr, s.Position())
	s.Map().AddObject(creature)

	s.Warnf("Object spawned successfully: %s", creature.GUID())
}

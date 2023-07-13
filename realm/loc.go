package realm

import (
	"github.com/gibson/gophercraft/realm/wdb"
	"github.com/gibson/gophercraft/realm/wdb/models"
)

func (s *Session) GetLoc(str string) string {
	var loc *models.LocString
	s.DB().Lookup(wdb.BucketKeyStringID, str, &loc)
	if loc == nil {
		return str
	}

	return loc.Text.GetLocalized(s.Locale)
}

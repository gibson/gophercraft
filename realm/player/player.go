package player

import (
	"github.com/gibson/gophercraft/packet/update"
	"github.com/gibson/gophercraft/realm/wdb/models"
)

// A Character can be a flesh-and-blood human, a bot, or a Goober.
type Character struct {
	*update.ValuesBlock
	Model        models.Character
	MovementInfo *update.MovementInfo
}

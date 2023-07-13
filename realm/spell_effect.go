package realm

import (
	"github.com/gibson/gophercraft/format/dbc/dbdefs"
	"github.com/gibson/gophercraft/packet/spell"
)

type SpellEffectData struct {
	Map         *Map
	Spell       *dbdefs.Ent_Spell
	EffectIndex int
	Caster      Unit
	Target      Unit
	Aura        *Aura
}

type SpellEffect func(*SpellEffectData)

func (s *Server) initSpellEffects() {
	s.SpellEffects = make([]SpellEffect, spell.NumEffects)
	s.SpellEffects[spell.EffectApplyAura] = EffectApplyAura
}

package commands

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/realm"
)

func cmdAuraList(s *realm.Session) {
	m := s.Map()
	if m == nil {
		return
	}

	target := s.GetTarget()

	if target == guid.Nil {
		target = s.GUID()
	}

	targetObject := m.GetObject(target)
	if targetObject == nil {
		return
	}

	targetUnit, ok := targetObject.(realm.Unit)
	if !ok {
		s.Warnf("%s is not a Unit", target)
		return
	}

	as := targetUnit.GetAuras()

	s.Warnf("%s", spew.Sdump(as))
}

func cmdAuraAdd(s *realm.Session, au uint32) {
	m := s.Map()
	if m == nil {
		return
	}

	target := s.GetTarget()

	if target == guid.Nil {
		target = s.GUID()
	}

	targetObject := m.GetObject(target)
	if targetObject == nil {
		return
	}

	targetUnit, ok := targetObject.(realm.Unit)
	if !ok {
		s.Warnf("%s is not a Unit", target)
		return
	}

	var aa realm.AuraApplication
	aa.ID = au
	aa.Level = 255
	aa.Caster = targetUnit

	if err := m.ApplyAura(targetUnit, &aa); err != nil {
		s.Warnf("%s", err)
		return
	}
}

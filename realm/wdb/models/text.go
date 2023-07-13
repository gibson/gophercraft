package models

import "github.com/gibson/gophercraft/i18n"

type NPCTextOptionEmote struct {
	Delay uint32
	ID    uint32
}

type NPCTextOption struct {
	Text  i18n.Text
	Lang  uint32
	Prob  float32
	Emote []NPCTextOptionEmote
}

type NPCText struct {
	ID    string
	Entry uint32
	Opts  []NPCTextOption
}

type LocString struct {
	ID   string
	Text i18n.Text
}

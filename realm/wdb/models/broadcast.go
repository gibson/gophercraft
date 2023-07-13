package models

import "github.com/gibson/gophercraft/i18n"

type Broadcast struct {
	Message i18n.Text
	Weight  int
}

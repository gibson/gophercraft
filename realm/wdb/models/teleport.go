package models

import "github.com/gibson/gophercraft/tempest"

type PortLocation struct {
	ID       string `xorm:"'port_id' pk" csv:"name"`
	Location tempest.C4Vector
	Map      uint32 `xorm:"'map'" csv:"mapID"`
}

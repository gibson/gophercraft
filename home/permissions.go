package home

import (
	"github.com/gibson/gophercraft/home/models"
	"github.com/gibson/gophercraft/home/rpcnet"
)

func (s *Server) CanEnlistRealm(acc *models.Account) bool {
	return acc.Tier >= rpcnet.Tier_Admin
}

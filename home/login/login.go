package login

import (
	"strings"

	"github.com/gibson/gophercraft/crypto/srp"
	"github.com/gibson/gophercraft/home/models"
	"github.com/gibson/gophercraft/home/rpcnet"
	"golang.org/x/crypto/bcrypt"
)

func Credentials(username, password string) []byte {
	I := strings.ToUpper(username)
	P := strings.ToUpper(password)
	return []byte(I + ":" + P)
}

// Returns true if matches hash
func Verify(username, password string, againstHash []byte) bool {
	webLogin := Credentials(username, password)
	ok := bcrypt.CompareHashAndPassword(againstHash, webLogin) == nil
	return ok
}

func SetAccount(acc *models.Account, username, password string, tier rpcnet.Tier) error {
	if err := PasswordValidate(password); err != nil {
		return err
	}

	if err := UsernameValidate(password); err != nil {
		return err
	}

	if password == "" {
		return ErrPasswordCannotBeEmpty
	}

	acc.Username = username
	acc.Tier = tier
	acc.IdentityHash = srp.HashCredentials(
		username, password,
	)
	var err error
	acc.WebLoginHash, err = bcrypt.GenerateFromPassword(
		Credentials(username, password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	return nil
}

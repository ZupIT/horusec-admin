package authz

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

const TokenLength = 20

type Authz struct {
	token     string
	createdAt time.Time
}

func NewAuthz() *Authz {
	authz := &Authz{}
	_ = authz.generateToken()

	return authz
}

func (a *Authz) generateToken() error {
	b := make([]byte, TokenLength)
	if _, err := rand.Read(b); err != nil {
		return err
	}

	a.token = hex.EncodeToString(b)
	a.createdAt = time.Now()

	return nil
}

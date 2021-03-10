package authz

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

const TokenLength = 20
const RefreshTokenInterval = 2 * time.Hour

type Authz struct {
	token                   string
	createdAt               time.Time
	quitRefreshTokenCycleCn chan struct{}
}

func NewAuthz() *Authz {
	authz := &Authz{}

	authz.generateToken()
	authz.startRefreshTokenCycle()

	return authz
}

func (a *Authz) startRefreshTokenCycle() {
	ticker := time.NewTicker(RefreshTokenInterval)
	a.quitRefreshTokenCycleCn = make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				a.generateToken()
				a.PrintToken()
			case <-a.quitRefreshTokenCycleCn:
				ticker.Stop()
				return
			}
		}
	}()
}

func (a *Authz) generateToken() {
	token, _ := a.getRandTokenString()

	a.token = token
	a.createdAt = time.Now()
}

func (a *Authz) getRandTokenString() (string, error) {
	b := make([]byte, TokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func (a *Authz) PrintToken() {
	print("Token:", a.token)
	print("Valid until:", a.GetTokenExpiresTime().String())
}

func (a *Authz) Stop() {
	close(a.quitRefreshTokenCycleCn)
}

func (a *Authz) GetToken() string {
	return a.token
}

func (a *Authz) GetTokenExpiresTime() time.Time {
	return a.createdAt.Add(RefreshTokenInterval)
}

package authz

import "time"

type Authz struct {
	token     string
	createdAt time.Time
}

func NewAuthz() *Authz {
	authz := &Authz{}
	authz.generateToken()

	return authz
}

func (a *Authz) generateToken() {
	a.token = "somthing"
	a.createdAt = time.Now()
}

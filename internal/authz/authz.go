// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authz

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/ZupIT/horusec-admin/internal/logger"
)

const TokenLength = 20
const RefreshTokenInterval = 2 * time.Hour

type Authz struct {
	token                   string
	createdAt               time.Time
	quitRefreshTokenCycleCn chan struct{}
}

func New() *Authz {
	authz := &Authz{}

	authz.generateNewTokenAndPrint()
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
				a.generateNewTokenAndPrint()
			case <-a.quitRefreshTokenCycleCn:
				ticker.Stop()
				return
			}
		}
	}()
}

func (a *Authz) generateNewTokenAndPrint() {
	token, _ := a.getRandTokenString()

	a.token = token
	a.createdAt = time.Now()

	a.PrintToken()
}

func (a *Authz) getRandTokenString() (string, error) {
	b := make([]byte, TokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func (a *Authz) PrintToken() {
	ctx := context.TODO()
	l := logger.WithPrefix(ctx, "authz")

	l.Info("Token:", a.token)
	l.Info("Valid until:", a.GetTokenExpiresTime().String())
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

func (a *Authz) IsValid(token string) bool {
	return token == a.token
}

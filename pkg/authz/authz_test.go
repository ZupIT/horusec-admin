package authz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	t.Run("should successfully create a new authz instance", func(t *testing.T) {
		authz := NewAuthz()
		assert.NotNil(t, authz)
	})

	t.Run("token and createAt values should be setted", func(t *testing.T) {
		authz := NewAuthz()
		assert.NotEmpty(t, authz.token)
		assert.NotEmpty(t, authz.createdAt)
	})
}

func TestGetToken(t *testing.T) {
	t.Run("should return the token value", func(t *testing.T) {
		authz := NewAuthz()
		assert.Equal(t, authz.token, authz.GetToken())
	})
}

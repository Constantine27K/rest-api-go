package teststore

import (
	"github.com/Constantine27K/rest-api-go.git/internal/app/model"
	"github.com/Constantine27K/rest-api-go.git/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := New()
	u := model.TestUser(t)
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u = model.TestUser(t)
	u.Email = "user@exmaple.com"
	s.User().Create(u)
	u, err = s.User().FindByEmail("user@exmaple.com")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

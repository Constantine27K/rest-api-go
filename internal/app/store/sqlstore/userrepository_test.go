package sqlstore_test

import (
	"github.com/Constantine27K/rest-api-go.git/internal/app/model"
	"github.com/Constantine27K/rest-api-go.git/internal/app/store"
	"github.com/Constantine27K/rest-api-go.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}
func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	s.User().Create(u)
	u, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
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

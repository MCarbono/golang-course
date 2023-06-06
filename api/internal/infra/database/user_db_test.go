package database

import (
	"api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, err := entity.NewUser("John", "j@j.com", "123456")
	if err != nil {
		t.Error(err)
	}
	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)
	insertedUser, err := userDB.FindByEmail("j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, insertedUser.Email, user.Email)
	assert.Equal(t, insertedUser.Name, user.Name)
	assert.Equal(t, insertedUser.ID, user.ID)
	assert.NotNil(t, insertedUser.Password)
}

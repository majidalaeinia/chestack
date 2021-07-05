package hooks

import (
	"github.com/MajidAlaeinia/chestack/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type user struct{}

func newUser() user {
	return user{}
}

func (u user) AfterInsert(exec boil.Executor, m *models.User) error {
	log.Println("triggered from user insert hook")
	return nil
}

func (u user) AfterUpdate(exec boil.Executor, m *models.User) error {
	log.Println("triggered from user update hook")
	return nil
}

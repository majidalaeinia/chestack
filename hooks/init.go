package hooks

import (
	"github.com/MajidAlaeinia/chestack/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Register() {
	u := newUser()
	models.AddUserHook(boil.AfterInsertHook, u.AfterInsert)
	models.AddUserHook(boil.AfterUpdateHook, u.AfterUpdate)
}

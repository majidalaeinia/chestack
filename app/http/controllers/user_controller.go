package controllers

import (
	"context"
	"fmt"
	"github.com/MajidAlaeinia/chestack/app/http/resources"
	"github.com/MajidAlaeinia/chestack/models"
	"github.com/MajidAlaeinia/chestack/modext"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"net/http"
)

type UserController struct{}

func (ctrl UserController) Show(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	var user resources.User
	err := queries.RawG(fmt.Sprintf("SELECT %s FROM `user` WHERE id = %d",
		modext.VisibleFieldsForRawSelect(), id)).BindG(context.Background(), &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"status":  false,
				"message": "No records was found.", //TODO: localization.
			})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

type UserCreateReqBody struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

func (ctrl UserController) Create(c *gin.Context) {
	var user models.User
	var reqBody UserCreateReqBody //TODO: Validation
	err := c.BindJSON(&reqBody)
	if err != nil {
		log.Println(err)
	}
	user.Name = reqBody.Name
	user.Email = reqBody.Email
	user.Mobile = reqBody.Mobile
	err = user.InsertG(boil.Infer())
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, user)
}

type UserUpdateReqBody struct {
	Name string `json:"name"`
}

func (ctrl UserController) Update(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	var reqBody UserUpdateReqBody //TODO: Validation
	err := c.BindJSON(&reqBody)
	user, err := models.Users(qm.Where("id = ?", id)).OneG()
	if err != nil {
		log.Println(err)
	}
	user.Name = reqBody.Name
	user.UpdateG(boil.Infer())

	c.JSON(http.StatusOK, user) //TODO: Hide redundant fields.
}

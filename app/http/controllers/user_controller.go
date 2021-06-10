package controllers

import (
	"github.com/MajidAlaeinia/chestack/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
)

type UserController struct{}

func (ctrl UserController) Show(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	usr, err := models.Users(qm.Where("id = ?", id)).OneG()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"status":  false,
				"message": "No records found.",
			})
		return
	}
	c.JSON(http.StatusOK, usr)
	return
}

package controllers

import (
	"context"
	"fmt"
	"github.com/MajidAlaeinia/chestack/app/http/resources"
	"github.com/MajidAlaeinia/chestack/modext"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries"
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

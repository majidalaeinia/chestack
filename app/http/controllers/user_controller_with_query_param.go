package controllers

import (
	"github.com/MajidAlaeinia/chestack/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl UserController) ChuckNorris(c *gin.Context) {
	fv := utils.Q(c)

	c.JSON(http.StatusOK, gin.H{"msg": fv.Q})
}

package routes

import (
	"github.com/MajidAlaeinia/chestack/app/http/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(r *gin.Engine) {
	r.GET("/handshake", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.GET("/users/:id", new(controllers.UserController).Show)
}

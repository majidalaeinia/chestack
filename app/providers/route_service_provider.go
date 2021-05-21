package providers

import (
	"github.com/MajidAlaeinia/chestack/routes"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	routes.SetRoutes(r)

	return r
}

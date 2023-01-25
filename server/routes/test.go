package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTestRoutes(rg *gin.RouterGroup) {
	test := rg.Group("/test")

	test.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "test")
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func basicController(r *gin.RouterGroup) {
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
}

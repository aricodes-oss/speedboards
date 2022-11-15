package controllers

import (
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"

	"backend/db"
)

type ControllerInjector func(*gin.RouterGroup)
type ControllerMapping map[string]ControllerInjector

var Router = gin.Default()
var apiGroup = Router.Group("/api")
var pathMappings = ControllerMapping{
	"/basic": basicController,
}

func mountControllers() {
	for path, injector := range pathMappings {
		group := apiGroup.Group(path)
		injector(group)
	}
}

func configureSessions() {
	store := gormsessions.NewStore(db.Conn, true, []byte("secret"))
	Router.Use(sessions.Sessions("session", store))
}

func Start() {
	mountControllers()
	Router.Run(":3000")
}

package main

import (
	"studier/controllers"

	"github.com/gin-gonic/gin"
)

var Routes = gin.Default();

func (a *App) InitRoutes() {
	Routes.GET("/verses", controllers.GetVerses)
}

func (a *App) Run(addr string) {
	Routes.Run(addr)
}


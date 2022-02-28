package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func  GetVerses(context *gin.Context){ 
	context.IndentedJSON(http.StatusOK, context.Request.URL.Query())
}
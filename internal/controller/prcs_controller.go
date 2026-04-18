package controller

import (
	l "cortex/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {

	l.Log.Info("ping endpoint hit")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pong",
	})
}

func AnalyzeHandler(c *gin.Context) {

}

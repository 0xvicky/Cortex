package controller

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	service "cortex/internal/service/analyser"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnalyserController struct {
	Service *service.AnalyserService
}

func NewAnalyserController(s *service.AnalyserService) *AnalyserController {
	return &AnalyserController{
		Service: s,
	}
}

func (h *AnalyserController) Pong(c *gin.Context) {

	l.Log.Info("ping endpoint hit")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pong",
	})
}

func (h *AnalyserController) AnalyzeRepoHandler(c *gin.Context) {
	var req model.AnalyseRepoRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	if analyRepoErr := h.Service.AnalyseRepo(c, req); analyRepoErr != nil {
		l.Log.Error("repo analyser error",
			zap.Error(analyRepoErr),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   analyRepoErr.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

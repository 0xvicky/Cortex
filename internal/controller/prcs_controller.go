package controller

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Pong(c *gin.Context) {

	l.Log.Info("ping endpoint hit")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pong",
	})
}

func AnalyzeRepoHandler(c *gin.Context) {
	var req model.AnalyseRepoRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	l.Log.Info("repo link received",
		zap.String("repo", req.RepoLink),
	)
	//Repo cloning into tmp
	//0. generate a uuid
	uuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		l.Log.Error("uuid generation error",
			zap.Error(uuidErr),
		)
	}

	l.Log.Info(uuid.String())
	//1. create a unique folder in a tmp folder => /tmp/{uuid}
	path := "Z:/Code/Golang/Projects/cortex/internal/tmp/repo/" + uuid.String()
	if pathErr := os.MkdirAll(path, os.ModePerm); pathErr != nil {
		l.Log.Error("error while creating repo path",
			zap.Error(pathErr),
		)
	}
	//2. Cloning the project in path
	cmd := exec.Command("git", "clone", req.RepoLink, path)
	if cmdErr := cmd.Run(); cmdErr != nil {
		l.Log.Error("error while cloning repo",
			zap.Error(cmdErr),
		)

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "clone success",
		"path":    path,
	})

}

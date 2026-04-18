package service

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AnalyserService struct{}

func NewAnalyserService() *AnalyserService {
	return &AnalyserService{}
}

func (s *AnalyserService) AnalyseRepo(c *gin.Context, repo model.AnalyseRepoRequest) error {

	l.Log.Info("repo link received",
		zap.String("repo", repo.RepoLink),
	)
	//Repo cloning into tmp
	//0. generate a uuid
	id, uuidErr := uuid.NewRandom()
	if uuidErr != nil {

		return fmt.Errorf("uuid generation error:%w", uuidErr)
	}
	l.Log.Info(id.String())
	//1. create a unique folder in a tmp folder => /tmp/{uuid}
	path := "Z:/Code/Golang/Projects/cortex/internal/tmp/repo/" + id.String()
	if pathErr := os.MkdirAll(path, os.ModePerm); pathErr != nil {
		return fmt.Errorf("path creation error:%w", pathErr)
	}
	//2. Cloning the project in path
	cmd := exec.Command("git", "clone", repo.RepoLink, path)
	if cmdErr := cmd.Run(); cmdErr != nil {
		return fmt.Errorf("cloning error:%w", cmdErr)
	}

	l.Log.Info(path)

	return nil

}

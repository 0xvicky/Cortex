package service

import (
	"cortex/internal/model"
	"cortex/internal/utils"
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AnalyserService struct{}

func NewAnalyserService() *AnalyserService {
	return &AnalyserService{}
}

func (s *AnalyserService) Analyse(c *gin.Context, repo model.AnalyseRequest) (int, error) {
	//Repo cloning into tmp
	// generate a uuid
	id, uuidErr := uuid.NewRandom()
	if uuidErr != nil {

		return 0, fmt.Errorf("uuid generation error:%w", uuidErr)
	}
	// create a unique folder in a tmp folder => /tmp/{uuid}
	path := "Z:/Code/Golang/Projects/cortex/internal/tmp/repo/" + id.String()
	if pathErr := os.MkdirAll(path, os.ModePerm); pathErr != nil {
		return 0, fmt.Errorf("path creation error:%w", pathErr)
	}
	//Cloning the project in path
	cmd := exec.Command("git", "clone", repo.RepoLink, path)
	if cmdErr := cmd.Run(); cmdErr != nil {
		return 0, fmt.Errorf("cloning error:%w", cmdErr)
	}

	//codebase scanner
	fileCount, scanErr := utils.DirScanner(path)
	if scanErr != nil {
		return 0, fmt.Errorf("error occured while traversing file path:%w", scanErr)
	}

	return fileCount, nil

}

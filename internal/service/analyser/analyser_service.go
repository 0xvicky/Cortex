package service

import (
	"cortex/internal/model"
	"cortex/internal/service/processor"
	"fmt"
	"os"
	"os/exec"
	"sync"

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
	//call the workers
	var wg sync.WaitGroup
	var queue chan model.ChannelData = make(chan model.ChannelData, 100)
	const nWorkers int = 10
	processor.WorkersInit(nWorkers, &wg, queue)

	//codebase scanner
	go func() {
		processor.DirScanner(path, queue)
		close(queue)
	}()
	wg.Wait()
	// if scanErr != nil {
	// 	return 0, fmt.Errorf("error occured while traversing file path:%w", scanErr)
	// }

	return 0, nil

}

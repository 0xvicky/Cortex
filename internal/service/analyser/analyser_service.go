package service

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"cortex/internal/utils"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AnalyserService struct{}

func NewAnalyserService() *AnalyserService {
	return &AnalyserService{}
}

func workersInit(nWorkers int, wg *sync.WaitGroup, queue <-chan model.ChannelData) {
	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for d := range queue {
				l.Log.Info("worker pick a new file",
					zap.Int("worker", workerId),
					zap.Int("file_no", d.FileNo),
					zap.String("file_path", d.FilePath),
				)
			}
		}(i)
	}
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
	var queue chan model.ChannelData = make(chan model.ChannelData)
	const nWorkers int = 10
	workersInit(nWorkers, &wg, queue)

	//codebase scanner
	go func() {
		utils.DirScanner(path, queue)
		close(queue)
	}()
	wg.Wait()
	// if scanErr != nil {
	// 	return 0, fmt.Errorf("error occured while traversing file path:%w", scanErr)
	// }

	return 0, nil

}

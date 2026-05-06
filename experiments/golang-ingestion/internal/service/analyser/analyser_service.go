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

// HStorage := processor.RepoStorage{Files: make(map[string][]model.AggregationResponse)}

func (s *AnalyserService) Analyse(c *gin.Context, repo model.AnalyseRequest) (processor.RepoStorage, error) {
	//Repo cloning into tmp
	// generate a uuid
	id, uuidErr := uuid.NewRandom()
	if uuidErr != nil {

		return processor.RepoStorage{}, fmt.Errorf("uuid generation error:%w", uuidErr)
	}

	// create a unique folder in a tmp folder => /tmp/{uuid}
	path := "Z:/Code/Golang/Projects/cortex/clone/repo/" + id.String()
	if pathErr := os.MkdirAll(path, os.ModePerm); pathErr != nil {
		return processor.RepoStorage{}, fmt.Errorf("path creation error:%w", pathErr)
	}
	//Cloning the project in path
	fmt.Println("[LOG] clone started")
	cmd := exec.Command("git", "clone", repo.RepoLink, path)
	if cmdErr := cmd.Run(); cmdErr != nil {
		return processor.RepoStorage{}, fmt.Errorf("cloning error:%w", cmdErr)
	}

	fmt.Println("[LOG] clone completed")
	//call the workers
	var wg sync.WaitGroup
	var aggrWg sync.WaitGroup
	aggrWg.Add(1)
	h := processor.RepoStorage{Files: make(map[string][]model.AggregationResponse)}
	var queue chan model.ChannelData = make(chan model.ChannelData, 100)
	var aggrQueue chan model.AggregationResponse = make(chan model.AggregationResponse, 100)
	const nWorkers int = 10

	// go func() {
	// }()
	go processor.AggregateResult(aggrQueue, &h, &aggrWg)

	processor.WorkersInit(nWorkers, &wg, queue, aggrQueue)

	//codebase scanner
	go func() {
		processor.DirScanner(path, queue)
		close(queue)
	}()

	go func() {
		wg.Wait()
		close(aggrQueue)
	}()
	aggrWg.Wait()

	// if scanErr != nil {
	// 	return 0, fmt.Errorf("error occured while traversing file path:%w", scanErr)
	// }

	return h, nil

}

package service

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"cortex/internal/utils"
	"errors"
	"fmt"
	"io"
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

// dummy ai response
func processChunk(d model.ChannelData, aggrQueue chan<- model.AggregationData) error {
	//from file path start reading the content
	file, _ := os.Open(d.FilePath)
	stat, _ := file.Stat()
	fileSize := stat.Size()
	var chunkSize = 2000
	var buffer = make([]byte, chunkSize)
	var chunkCount = 1
	//create a fix amount of chunk of characters or words
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			continue
		}
		var payload = model.AIChunkPayload{
			FileNo:   d.FileNo,
			FilePath: d.FilePath,
			ChunkId:  chunkCount,
			Payload:  buffer,
		}
		if n > 0 {
			//send to ai layer(dummy rn) and push the response to aggregation channel
			aiRes := utils.AIDummy(payload)
			aggrQueue <- aiRes
		}

		//overlapping logic
		if chunkCount*chunkSize < int(fileSize) {
			_, seekErr := file.Seek(-200, io.SeekCurrent)
			if seekErr != nil {
				return seekErr
			}
		}
		chunkCount++
	}

	if fileClErr := file.Close(); fileClErr != nil {
		return fileClErr
	}

	//send the next chunk and reapeat the process until file is done reading
	return nil
	//error handling as well
}

func workersInit(nWorkers int, wg *sync.WaitGroup, queue <-chan model.ChannelData) {
	var aggrQueue = make(chan model.AggregationData)
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
				processChunk(d, aggrQueue)
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
	var queue chan model.ChannelData = make(chan model.ChannelData, 100)
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

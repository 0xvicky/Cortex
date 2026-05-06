package processor

import (
	"cortex/internal/model"
	"fmt"
	"sync"
)

func AggregateResult(aggrQueue <-chan model.AggregationResponse, hs *RepoStorage, aggrWg *sync.WaitGroup) {
	//mapping[file_name]=>[chunk1, chunk2, chunk3]
	defer func() {
		fmt.Println("[LOG] processing completed")
		aggrWg.Done()
	}()
	for response := range aggrQueue {
		hs.AddChunk(response)
	}
}

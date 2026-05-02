package processor

import (
	"cortex/internal/model"
	"fmt"
	"sync"
)

func WorkersInit(nWorkers int, wg *sync.WaitGroup, queue <-chan model.ChannelData, aggrQueue chan<- model.AggregationResponse) {

	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for d := range queue {

				processErr := ProcessChunk(d, aggrQueue)
				if processErr != nil {
					fmt.Println("[ERROR]:%w", processErr)
					continue
				}
			}
		}(i)
	}

}

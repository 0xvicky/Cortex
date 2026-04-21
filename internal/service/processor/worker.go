package processor

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"sync"

	"go.uber.org/zap"
)

func WorkersInit(nWorkers int, wg *sync.WaitGroup, queue <-chan model.ChannelData, aggrQueue chan<- model.AggregationResponse) {

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
				ProcessChunk(d, aggrQueue)
			}
		}(i)
	}

}

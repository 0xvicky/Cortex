package processor

import (
	"cortex/internal/model"
	"cortex/internal/utils"
	"errors"
	"io"
	"os"
)

func ProcessChunk(d model.ChannelData, aggrQueue chan<- model.AggregationResponse) error {
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
			Payload:  buffer[:n],
		}
		if n > 0 {
			//send to ai layer(dummy rn) and push the response to aggregation channel
			aiRes := utils.AIDummy(payload)
			aggrQueue <- aiRes
		}

		//overlapping logic
		fileOffset, _ := file.Seek(0, io.SeekCurrent)
		if fileOffset < fileSize {
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

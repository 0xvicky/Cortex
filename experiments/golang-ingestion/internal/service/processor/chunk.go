package processor

import (
	"cortex/internal/model"
	"cortex/internal/utils"
	"errors"
	"io"
	"math"
	"os"
)

func ProcessChunk(d model.ChannelData, aggrQueue chan<- model.AggregationResponse) error {
	file, err := os.Open(d.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := stat.Size()
	fileName := stat.Name()

	const chunkSize = 2000
	const overlap = 200
	const stride = chunkSize - overlap

	totalChunks := int(math.Ceil(float64(fileSize) / float64(stride)))
	buffer := make([]byte, chunkSize)
	chunkCount := 0

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err // ← surface real errors
		}

		if n == 0 {
			continue
		}

		payload := model.AIChunkPayload{
			FileNo:      d.FileNo,
			FilePath:    d.FilePath,
			FileName:    fileName,
			ChunkId:     chunkCount,
			TotalChunks: totalChunks,
			Payload:     buffer[:n],
		}

		aiRes := utils.AIDummy(payload)
		aggrQueue <- aiRes
		chunkCount++

		// overlap — seek back so next read has context from previous chunk
		fileOffset, err := file.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		if fileOffset < fileSize {
			_, err = file.Seek(-int64(overlap), io.SeekCurrent)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

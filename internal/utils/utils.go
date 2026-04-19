package utils

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"io/fs"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DirScanner(repoPath string, queue chan<- model.ChannelData) (int, error) {
	l.Log.Info("In dir scanner",
		zap.String("string", repoPath),
	)
	var nFiles int = 0
	filepath.WalkDir(repoPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			nFiles++
			fileData := model.ChannelData{
				FilePath: path,
				FileNo:   nFiles,
			}
			queue <- fileData
			// l.Log.Info("file processed",
			// 	zap.Int("file_no", nFiles),
			// 	zap.String("file_path", path),
			// )
		}

		return nil
	})
	return nFiles, nil
}

func AIDummy(payload model.AIChunkPayload) model.AggregationResponse {
	var response = model.AggregationResponse{
		FileNo:   payload.FileNo,
		FilePath: payload.FilePath,
		ChunkId:  payload.ChunkId,
		AIResponse: gin.H{
			"insights": []string{
				"Good separation of concerns",
				"Follows RESTful design",
			},

			"risks": []string{
				"No input validation in some handlers",
				"Lack of error handling in DB layer",
			},

			"recommendations": []string{
				"Add middleware for validation",
				"Improve error handling",
			},
		},
	}

	return response
}

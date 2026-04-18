package utils

import (
	l "cortex/internal/logger"
	"cortex/internal/model"
	"io/fs"
	"path/filepath"

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

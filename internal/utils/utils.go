package utils

import (
	l "cortex/internal/logger"
	"io/fs"
	"path/filepath"

	"go.uber.org/zap"
)

func DirScanner(repoPath string) (int, error) {
	l.Log.Info("In dir scanner",
		zap.String("string", repoPath),
	)
	var nFiles int
	filepath.WalkDir(repoPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		nFiles++
		l.Log.Info("file processed",
			zap.Int("file_no", nFiles),
			zap.String("file_path", path),
		)

		return nil
	})
	return nFiles, nil
}

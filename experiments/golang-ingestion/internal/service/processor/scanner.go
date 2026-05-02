package processor

import (
	"cortex/internal/constants"
	l "cortex/internal/logger"
	"cortex/internal/model"
	"cortex/internal/utils"
	"io/fs"
	"path/filepath"

	"go.uber.org/zap"
)

func DirScanner(repoPath string, queue chan<- model.ChannelData) (int, error) {
	l.Log.Info("In dir scanner",
		zap.String("string", repoPath),
	)
	var nFiles int = 0
	walkErr := filepath.WalkDir(repoPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// skip directories early
		if d.IsDir() {
			if constants.SkipDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil // don't process dirs, just walk into them
		}

		if utils.ShouldProcess(path, d) {
			nFiles++
			fileData := model.ChannelData{
				FilePath: path,
				FileNo:   nFiles,
			}
			queue <- fileData
		}

		return nil
	})

	if walkErr != nil {
		return 0, walkErr
	}
	return nFiles, nil
}

package utils

import (
	"cortex/internal/constants"
	"cortex/internal/model"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func AIDummy(payload model.AIChunkPayload) model.AggregationResponse {
	var response = model.AggregationResponse{
		FileNo:      payload.FileNo,
		FilePath:    payload.FilePath,
		FileName:    payload.FileName,
		TotalChunks: payload.TotalChunks,
		ChunkId:     payload.ChunkId,
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

func isBinaryFile(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return true
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return true
	}

	// 1. Use Go's built-in MIME detection
	contentType := http.DetectContentType(buf[:n])
	if !strings.HasPrefix(contentType, "text/") {
		return true
	}

	// 2. Fallback null byte check for edge cases
	if slices.Contains(buf[:n], byte(0)) {
		return true
	}
	return false

}
func ShouldProcess(path string, d fs.DirEntry) bool {
	//if dir, we'll process separately
	var name = d.Name()
	var ext = filepath.Ext(name)

	//skip extensions
	if constants.SkipExtensions[ext] {
		return false
	}

	//skip filenames
	if constants.SkipFilenames[name] {
		return false
	}

	// 4. Skip by filename pattern
	for _, pattern := range constants.SkipFilenamePatterns {
		if matched, _ := filepath.Match(pattern, name); matched {
			return false
		}
	}

	return !isBinaryFile(path)

}

package processor

import (
	"cortex/internal/model"

	"github.com/samber/lo"
)

type RepoStorage struct {
	Files map[string][]model.AggregationResponse
}

func NewRepoStorage() *RepoStorage {
	return &RepoStorage{
		Files: make(map[string][]model.AggregationResponse),
	}
}
func (hs *RepoStorage) AddChunk(chunk model.AggregationResponse) {
	if !lo.HasKey(hs.Files, chunk.FilePath) {
		// hs.[chunk.FilePath]
		var collection = make([]model.AggregationResponse, chunk.TotalChunks)
		hs.Files[chunk.FilePath] = collection
	}
	hs.Files[chunk.FilePath][chunk.ChunkId] = chunk
}

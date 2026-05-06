package model

type AnalyseRequest struct {
	RepoLink string `json:"repoUrl" binding:"required"`
}

type ChannelData struct {
	FileNo   int
	FilePath string
}

type AIChunkPayload struct {
	FileNo      int
	FilePath    string
	FileName    string
	ChunkId     int
	TotalChunks int
	Payload     []byte
}

type AggregationResponse struct {
	FileNo      int
	FilePath    string
	FileName    string
	ChunkId     int
	TotalChunks int
	AIResponse  any // will change later

}

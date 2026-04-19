package model

type AnalyseRequest struct {
	RepoLink string `json:"repoUrl" binding:"required"`
}

type ChannelData struct {
	FileNo   int
	FilePath string
}

type AIChunkPayload struct {
	FileNo   int
	FilePath string
	ChunkId  int
	Payload  []byte
}

type AggregationResponse struct {
	FileNo     int
	FilePath   string
	ChunkId    int
	AIResponse any // will change later
}

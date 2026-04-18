package model

type AnalyseRequest struct {
	RepoLink string `json:"repoUrl" binding:"required"`
}

type ChannelData struct {
	FileNo   int
	FilePath string
}

package model

type AnalyseRequest struct {
	RepoLink string `json:"repoUrl" binding:"required"`
}

package model

type AnalyseRepoRequest struct {
	RepoLink string `json:"repoUrl" binding:"required"`
}

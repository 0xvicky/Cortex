package utils

import (
	"cortex/internal/model"

	"github.com/gin-gonic/gin"
)

func AIDummy(payload model.AIChunkPayload) model.AggregationResponse {
	var response = model.AggregationResponse{
		FileNo:   payload.FileNo,
		FilePath: payload.FilePath,
		ChunkId:  payload.ChunkId,
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

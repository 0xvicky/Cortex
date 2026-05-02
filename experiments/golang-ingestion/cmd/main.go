package main

import (
	"cortex/internal/controller"
	l "cortex/internal/logger"
	service "cortex/internal/service/analyser"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	//init logger
	l.InitLogger()
	if l.Log == nil {
		panic("logger not initialized")
	}
	defer l.Log.Sync()
	// fmt.Println("Cortex:AI codebase intelligence engine")
	l.Log.Info("Server Code Runnig",
		zap.String("app", "Cortex:AI codebase intelligence engine"),
	)

	//server
	w := gin.Default()
	//services and controllers instances
	svc := service.NewAnalyserService()
	h := controller.NewAnalyserController(svc)

	w.GET("/ping", h.Pong)
	w.POST("/analyze", h.AnalyzeHandler)

	l.Log.Info("listening at 6969")
	if err := w.Run(":6969"); err != nil {
		// fmt.Println("[ERROR] error running server:%w", err)
		l.Log.Fatal("error while starting server",
			zap.Any("error", err),
		)
	}

}

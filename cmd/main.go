package main

import (
	"cortex/internal/controller"
	l "cortex/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	//init logger
	l.InitLogger()
	defer l.Log.Sync()
	// fmt.Println("Cortex:AI codebase intelligence engine")
	l.Log.Info("Server Code Runnig",
		zap.String("app", "Cortex:AI codebase intelligence engine"),
	)

	//server
	w := gin.Default()
	w.GET("/ping", controller.Pong)

	l.Log.Info("listening at 6969")
	if err := w.Run(":6969"); err != nil {
		// fmt.Println("[ERROR] error running server:%w", err)
		l.Log.Fatal("error while starting server",
			zap.Any("error", err),
		)
	}

}

//setup for a basic webserver using gin✅
//add air✅
//basic unit testing framework

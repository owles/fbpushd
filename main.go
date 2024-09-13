package main

import (
	_ "fbpushd/docs"
	"fbpushd/internal/api"
	"fbpushd/internal/push"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

// @title Firebase Push Notification API
// @version 1.0
// @description A simple API to send push notifications using Firebase
// @host 0.0.0.0:8080
// @BasePath /
// @schemes http
func main() {
	app, _, err := push.SetupFirebase()
	if err != nil {
		slog.Error("Failed to initialize Firebase", "message", err.Error())
	}

	if os.Getenv("MODE") != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	web := api.NewAPI(app)
	web.RegisterRoutes(r, app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("Listening on " + port)

	err = r.Run(":" + port)
	if err != nil {
		slog.Error("Failed to start server", "message", err.Error())
	}
}

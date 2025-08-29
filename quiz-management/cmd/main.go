package main

import (
	"log"

	"github.com/Sasikumar00/QuizForgeAI/quiz-management/config"
	"github.com/Sasikumar00/QuizForgeAI/quiz-management/internal/database"
	"github.com/Sasikumar00/QuizForgeAI/quiz-management/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	client, err := database.ConnectMongo(cfg.MongoURL)
	if err != nil {
		log.Fatal("MongoDB connection failed", err)
	}
	db := client.Database(cfg.DBName)

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
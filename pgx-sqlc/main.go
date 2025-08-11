package main

import (
	"log"
	"study-gin/internal/db"
	"study-gin/internal/handlers"
	"study-gin/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env not found")
	}

	if err := db.InitDB(); err != nil {
		log.Fatal("unable to connect to db")
	}

	log.Println(db.DB)

	r := gin.Default()

	userRepository := repository.NewSQLUserRepository(db.DB)
	userHandler := handlers.NewUserHandler(userRepository)
	r.GET("/api/v1/users/:id", userHandler.GetUserByUuid)
	r.POST("/api/v1/users", userHandler.CreateUser)

	r.Run(":8080")
}
package main

import (
	"app/config"
	"app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	r := gin.Default()

	r.Static("/uploads", "./uploads")
	uploadDirs := [...]string{"users", "product"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 755)
	}

	routes.Serve(r)
	r.Run(":" + os.Getenv("PORT"))
}

package main

import (
	"log"
	"myapp/internal/database"
	"myapp/internal/routes"
)

func main() {

	database.Connect("mongodb+srv://webashraf2:6s9snA3cDA6Qk6yE@cluster0.uxmgd.mongodb.net/go-first-project?retryWrites=true&w=majority&appName=Cluster0")


	r := routes.SetupRouter()
	log.Println("🚀 Server running on port 8080")
	r.Run(":8080") 
}

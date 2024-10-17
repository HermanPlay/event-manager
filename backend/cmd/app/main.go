package main

import (
	"log"
	"os"
	"strconv"

	"github.com/HermanPlay/web-app-backend/internal/api/http"
	"github.com/HermanPlay/web-app-backend/internal/api/http/server"
	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	if _, ok := os.LookupEnv("DEBUG"); ok {
		err := godotenv.Load("././.env")
		if err != nil {
			panic(err)
		}
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("No config!")
	}

	init := http.Init(cfg)
	app := server.Init(init)

	log.Println("Server is running on port:", cfg.App.Port)
	app.Run(":" + strconv.Itoa(cfg.App.Port))
}

package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/OybekAbduvosiqov/book/api"
	"github.com/OybekAbduvosiqov/book/config"

	"context"

	"github.com/OybekAbduvosiqov/book/storage/postgres"
)

func main() {

	cfg := config.Load()

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer storage.CloseDB()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.NewApi(r, storage)

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"blog-v2/src"
	"blog-v2/src/adapters/httpserver"
	"embed"
	"log"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
)

//go:embed posts/*

var posts embed.FS

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	config, err := loadAppConfig()
	if err != nil {
		log.Fatalf("failed to load config - %v", err)
	}

	app := src.NewApp(ctx, posts)
	//if err != nil {
	//	log.Fatal("failed to create app")
	//}

	//serverMiddlewares, err := httpserver.NewMiddlewares()
	//if err != nil {
	//	log.Fatal("failed to create middlewares")
	//}

	router := httpserver.NewRouter(
		app.BlogService,
	)
	server := gracefulshutdown.NewServer(httpserver.NewWebServer(config.ServerConfig, router))

	log.Printf("Listening on port %s\n", config.ServerConfig.Port)
	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatal("http server listen failed", err)
	}
}

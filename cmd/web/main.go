package main

import (
	"blog-v2/src"
	"blog-v2/src/adapters/httpserver"
	"log"
	"os"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	config, err := loadAppConfig()
	if err != nil {
		log.Fatalf("failed to load config - %v", err)
	}

	app := src.NewApp(ctx, os.DirFS("./cmd/web/posts"))

	router := httpserver.NewRouter(
		app.BlogService,
	)
	server := gracefulshutdown.NewServer(httpserver.NewWebServer(config.ServerConfig, router))

	log.Printf("Listening on port %s\n", config.ServerConfig.Port)
	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatal("http server listen failed", err)
	}
}

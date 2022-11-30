package main

import (
	"blog-v2/src/adapters/httpserver"
	"log"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	config, err := loadAppConfig()
	if err != nil {
		log.Fatalf("failed to load config - %v", err)
	}

	app := newApp(ctx)
	//if err != nil {
	//	log.Fatal("failed to create app")
	//}

	//serverMiddlewares, err := httpserver.NewMiddlewares()
	//if err != nil {
	//	log.Fatal("failed to create middlewares")
	//}

	router := httpserver.NewRouter(
		app.Greeter,
		app.BlogService,
	)
	server := gracefulshutdown.NewServer(httpserver.NewWebServer(config.ServerConfig, router))

	log.Printf("Listening on port %s\n", config.ServerConfig.Port)
	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatal("http server listen failed", err)
	}
}

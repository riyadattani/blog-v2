package main

import (
	"blog-v2/src/domain/greet"
	"context"

	greethandler2 "blog-v2/src/adapters/httpserver/greethandler"
)

// App holds and creates dependencies for your application.
type App struct {
	Greeter greethandler2.GreeterService
}

func newApp(applicationContext context.Context) *App {
	go handleInterrupts(applicationContext)

	//appMetrics, err := metrics.NewClient()
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create app metrics client - %w", err)
	//}

	return &App{
		Greeter: greethandler2.GreeterServiceFunc(greet.HelloGreeter),
	}
}

// this is just an example of how the services within an app could listen to the
// cancellation signal and stop their work gracefully. So it's probably a decent
// idea to pass the application context to services if you want to do some
// cleanup before finishing.
func handleInterrupts(ctx context.Context) {
	<-ctx.Done()
	// logger.Info(ctx, "shutting down")
}

package src

import (
	"blog-v2/src/adapters/httpserver/bloghandler"
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/domain/blog"
	"context"
	"io/fs"
)

// App holds and creates dependencies for your application.
type App struct {
	BlogService bloghandler.BlogService
}

func NewApp(applicationContext context.Context, posts fs.FS) *App {
	go handleInterrupts(applicationContext)

	//appMetrics, err := metrics.NewClient()
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create app metrics client - %w", err)
	//}

	return &App{
		BlogService: blog.NewService(inmem.NewRepository(posts)),
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

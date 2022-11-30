package httpserver

import (
	"github.com/gorilla/mux"
)

type Middlewares struct {
	O11y    mux.MiddlewareFunc
	General []mux.MiddlewareFunc
}

func NewMiddlewares() (Middlewares, error) {
	//o11yMiddleware, err := middleware.New(tki, middleware.WithBodyLogging(shouldLogBodies))
	//if err != nil {
	//	return Middlewares{}, fmt.Errorf("failed to create o11y middleware: %w", err)
	//}

	return Middlewares{
		// O11y:    o11yMiddleware,
		General: []mux.MiddlewareFunc{},
	}, nil
}

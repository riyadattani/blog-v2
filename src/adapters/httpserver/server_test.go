package httpserver_test

import (
	"testing"
)

func TestNewWebServer(t *testing.T) {
	//router := httpserver.NewRouter(
	//	greethandler.GreeterServiceFunc(greet.HelloGreeter),
	//)
	//webServer := httpserver.NewWebServer(
	//	httpserver.ServerConfig{},
	//	router,
	//)
	//
	//svr := httptest.NewServer(webServer.Handler)
	//defer svr.Close()
	//
	//client := acceptance.NewAPIClient(http.DefaultTransport, svr.URL)
	//
	//specifications.Greeting{
	//	Subject: client,
	//	MakeContext: func(tb testing.TB) context.Context {
	//		return context.Background()
	//	},
	//}.Test(t)
}

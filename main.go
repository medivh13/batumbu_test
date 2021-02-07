package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/medivh13/batumbu_test/internal/transport/http/middleware"

	handlers "github.com/medivh13/batumbu_test/internal/transport/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/medivh13/batumbu_test/internal/services"
)

func main() {
	httpAddr := flag.String("http", ":8080", "http listen address")
	errChan := make(chan error)

	godotenv.Load()

	e := echo.New()
	m := middleware.NewMiddleware()

	e.Use(m.CORS)

	srv := services.NewService()
	handlers.NewHttpHandler(e, srv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- e.Start(*httpAddr)
	}()

	err := <-errChan
	log.Error(err.Error())
}

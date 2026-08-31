package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v5"
	mw "github.com/labstack/echo/v5/middleware"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "using -p <port>")
}

type GreetingResponse struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Timestamp string `json:"timestamp,omitempty"`
}

// Greetings add value from query parameter name and write to request
// greetings as string. Return nil by success, and error if any error
// detected
func Greetings(c *echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "Незнакомец"
	}

	response := GreetingResponse{
		Message:   fmt.Sprintf("Привет, %s!!!", name),
		Code:      http.StatusOK,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	flag.Parse()

	e := echo.New()

	e.Use(mw.RequestLogger())
	e.Use(mw.Recover())

	e.GET("/hello", Greetings)

	s := http.Server{
		Addr:              fmt.Sprintf("[::]:%s", port),
		Handler:           e,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Print(fmt.Sprintf("Start server on 127.0.0.1:%s ...", port))
	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Print(err.Error())
		}

		log.Fatal(err.Error())
	}
}

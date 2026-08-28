// simple
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    mw "github.com/labstack/echo/v5/middleware"
    echo "github.com/labstack/echo/v5"
)

// greetings add value from query parameter name and write to request
// greetings as string. Return nil by success, and error if any error
// detected
func greetings(c *echo.Context) error {
    name := c.QueryParam("name")
    if name == "" {
       name = "Stranger"
    }

    return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!!!", name))
}

func main() {
    e := echo.New()

    e.Use(mw.RequestLogger())
    e.Use(mw.Recover())

    e.GET("/hello", greetings)
    e.GET("/panic", panicHandler)

    s := http.Server{
        Addr: "[::]:8080",
        Handler: e,
        ReadHeaderTimeout: 5*time.Second,
        ReadTimeout: 15*time.Second,
        WriteTimeout: 30*time.Second,
        IdleTimeout: 120*time.Second,
    }

    log.Print("Start server on 127.0.0.1:8080 ...")
    if err := s.ListenAndServe(); err != nil {
        if err == http.ErrServerClosed {
            log.Print(err.Error())
        }

        log.Fatal(err.Error())
    }
}

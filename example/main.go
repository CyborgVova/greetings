package main

import (
    "context"
    "fmt"
    "time"

    sdk "github.com/cyborgvova/greetings/greetings_sdk"
)

func main() {
    client := sdk.NewClient(
        sdk.WithBaseURL("http://localhost:8080"),
//        sdk.WithDebug(true),
        sdk.WithAPIKey("my_secret_api_key"),
        sdk.WithTimeout(5*time.Second),
    )

    resp, err := client.Greeting(context.Background(), "Антон")
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.Message) // Привет, Антон!!!

    resp, err = client.Greeting(context.Background(), "")
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.Message) // Привет, Незнакомец!!!
}

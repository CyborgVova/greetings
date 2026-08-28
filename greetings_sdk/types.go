package greetings_sdk

// GreetingResponse - структура ответа от API
type GreetingResponse struct {
    Message     string  `json:"message"`
    Code        int     `json:"code"`
    Timestamp   string  `json:"timestamp,omitempty"`
}

type ErrorResponse struct {
    Error   string  `json:"error"`
    Status  string  `json:"status"`
}

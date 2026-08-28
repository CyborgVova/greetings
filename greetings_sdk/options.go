package greetings_sdk

import (
    "net/http"
    "time"
)

// Config - настройки клиента
type Config struct {
    BaseURL     string
    APIKey      string
    Timeout     time.Duration
    HTTPClient  *http.Client
    UserAgent   string
    Debug       bool
}

// DefaultConfig - значения по умолчанию
func DefaultConfig() *Config {
    return &Config{
        BaseURL:    "http://localhost",
        Timeout:    10*time.Second,
        UserAgent:  "greetings_sdk/1.0",
        Debug:      false,
    }
}

// Option - функция-опция для изменения конфига
type Option func(*Config)

// WithBaseURL - устанавливает базовый URL
func WithBaseURL(url string) Option {
    return func(cfg *Config) {
        cfg.BaseURL = url
    }
}

// WithAPIKey - устанавливает API-ключ
func WithAPIKey(key string) Option {
    return func(cfg *Config) {
        cfg.APIKey = key
    }
}

// WithTimeout - устанавливает таймаут
func WithTimeout(d time.Duration) Option {
    return func(cfg *Config) {
        cfg.Timeout = d
    }
}

// WithHTTPClient - позволяет передать свой http.Client
func WithHTTPClient(c *http.Client) Option {
    return func(cfg *Config) {
        cfg.HTTPClient = c
    }
}

// WithDebug - включает режим отладки (логирование)
func WithDebug(debug bool) Option {
    return func(cfg *Config) {
        cfg.Debug = debug
    }
}

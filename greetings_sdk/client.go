package greetings_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Client - основной клиент для работы с API
type Client struct {
	config     *Config
	httpClient *http.Client
	logger     *log.Logger
}

// NewClient - конструктор (инициализация SDK)
func NewClient(opts ...Option) *Client {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.HTTPClient == nil {
		if cfg.Timeout.Seconds() < 15.0 {
			cfg.Timeout = 15 * time.Second
		}

		cfg.HTTPClient = &http.Client{
			Timeout: cfg.Timeout,
		}
	}

	logger := log.New(io.Discard, "[GREETINGS_SDK] ", log.LstdFlags)
	if cfg.Debug {
		logger = log.New(io.Writer(os.Stdout), "[GREETINGS_SDK] ", log.LstdFlags)
	}

	return &Client{
		config:     cfg,
		httpClient: cfg.HTTPClient,
		logger:     logger,
	}
}

// Greeting - основной метод создает приветствие
func (c *Client) Greeting(ctx context.Context, name string) (*GreetingResponse, error) {
	// 1. Логируем начало запроса
	c.logger.Printf("Sending greeting request for: %s", name)

	// 2. Строим URL
	url := fmt.Sprintf("%s/hello?name=%s", c.config.BaseURL, name)
	c.logger.Printf("Request URL: %s", url)

	// 3. Создаем запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 4. Добавляем заголовки
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.config.UserAgent)
	if c.config.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	}

	// 5. Выполняем запрос
	start := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Printf("Request failed: %v", err)
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	c.logger.Printf("Response received in %v, status: %d", time.Since(start), resp.StatusCode)

	// 6. Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// 7. Проверяем статус
	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err == nil {
			return nil, &ErrInvalidStatus{
				StatusCode: resp.StatusCode,
				Body:       errResp.Error,
			}
		}

		return nil, &ErrInvalidStatus{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}

	// 8. Парсим JSON в структуру
	var result GreetingResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	c.logger.Printf("Successfully parsed response: %+v", result)

	return &result, nil
}

// GreetingsRaw - если нужен сырой ответ (для отладки или кастомной обработки)
func (c *Client) GreetingsRaw(ctx context.Context, name string) ([]byte, error) {
	url := fmt.Sprintf("%s/hello/%s", c.config.BaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.config.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &ErrInvalidStatus{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}

	return body, nil
}

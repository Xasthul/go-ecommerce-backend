package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ProductClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewProductClient(baseURL string, apiKey string) *ProductClient {
	return &ProductClient{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

type ProductResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PriceCents int    `json:"price_cents"`
	Currency   string `json:"currency"`
	Stock      int    `json:"stock"`
}

func (p *ProductClient) GetProduct(ctx context.Context, productID string) (*ProductResponse, error) {
	url := fmt.Sprintf("%s/internal/products/%s", p.baseURL, productID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-API-KEY", p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned %d", resp.StatusCode)
	}

	var pr ProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return nil, err
	}

	return &pr, nil
}

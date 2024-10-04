package client

import (
	"context"
	"encoding/json"
	"fmt"
	"get-price/types"
	"net/http"
)

type client struct {
	endPoint string
}

func NewClient(endPoint string) *client {
	return &client{endPoint: endPoint}
}

func (c *client) GetPrice(ctx context.Context, key string) (*types.PriceResponse, error) {
	endPoint := fmt.Sprintf("%s?key=%s", c.endPoint, key)

	req, err := http.NewRequest("get", endPoint, nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(types.PriceResponse)

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}

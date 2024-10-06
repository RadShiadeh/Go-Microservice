package client

import (
	"context"
	"encoding/json"
	"fmt"
	"get-price/types"
	"net/http"
)

type JSONClient struct {
	endPoint string
}

func NewJSONClient(endPoint string) *JSONClient {

	return &JSONClient{endPoint: endPoint}
}

func (c *JSONClient) GetPrice(ctx context.Context, key string, curr string) (*types.PriceResponse, error) {

	endPoint := fmt.Sprintf("%s?key=%s&curr=%s", c.endPoint, key, curr) //?key=bitcoin&curr=usd

	req, err := http.NewRequest("get", endPoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error in client GetPrice while making the new req: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in the GetPrice client where trying to get a resonse: %s", err)
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusAccepted) {
		httpErr := map[string]error{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, fmt.Errorf("error in GetPrice Client when encoding the resp: %s", err)
		}

		return nil, fmt.Errorf("service responsed with non ok status code: %s", httpErr["error"])
	}

	defer resp.Body.Close()

	res := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("error GetPrice Client when Decoing res: %s", err)
	}

	return res, nil
}

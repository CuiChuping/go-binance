package portfolio_pro

import (
	"context"
	"encoding/json"
	"net/http"
)

type MarginAssetLeverageService struct {
	c *Client
}

func (s *MarginAssetLeverageService) Do(ctx context.Context, opts ...RequestOption) ([]*MarginAssetLeverageResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/portfolio/margin-asset-leverage",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*MarginAssetLeverageResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MarginAssetLeverageResponse struct {
	Asset    string `json:"asset"`
	Leverage int32  `json:"leverage"`
}

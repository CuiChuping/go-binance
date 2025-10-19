package portfolio_pro

import (
	"context"
	"encoding/json"
	"net/http"
)

type CollateralRateService struct {
	c *Client
}

func (s *CollateralRateService) Do(ctx context.Context, opts ...RequestOption) ([]*CollateralRateResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/portfolio/collateralRate",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*CollateralRateResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CollateralRateResponse struct {
	Asset          string  `json:"asset"`
	CollateralRate float64 `json:"collateralRate"`
}

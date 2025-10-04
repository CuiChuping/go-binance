package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type InsuranceBalanceService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *InsuranceBalanceService) Symbol(symbol string) *InsuranceBalanceService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *InsuranceBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*InsuranceBalanceGroup, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/apiTradingStatus",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*InsuranceBalanceGroup, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type InsuranceBalance struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	UpdateTime    int64  `json:"updateTime"`
}

type InsuranceBalanceGroup struct {
	Symbols []string           `json:"symbols"`
	Assets  []InsuranceBalance `json:"assets"`
}

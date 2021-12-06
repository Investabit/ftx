package structs

import (
	"github.com/guregu/null"
	"time"
)

type HistoricalPrices struct {
	Success bool `json:"success"`
	Result  []struct {
		Close     float64   `json:"close"`
		High      float64   `json:"high"`
		Low       float64   `json:"low"`
		Open      float64   `json:"open"`
		StartTime time.Time `json:"startTime"`
		Volume    float64   `json:"volume"`
	} `json:"result"`
}

type Trades struct {
	Success bool `json:"success"`
	Result  []struct {
		ID          int64     `json:"id"`
		Liquidation bool      `json:"liquidation"`
		Price       float64   `json:"price"`
		Side        string    `json:"side"`
		Size        float64   `json:"size"`
		Time        time.Time `json:"time"`
	} `json:"result"`
}

type Orderbook struct {
	Success bool `json:"success"`
	Result  struct {
		Asks [][]float64 `json:"asks"`
		Bids [][]float64 `json:"bids"`
	} `json:"result"`
}

type MarketsResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		Name                  string      `json:"name"`
		Enabled               bool        `json:"enabled"`
		PostOnly              bool        `json:"postOnly"`
		PriceIncrement        float64     `json:"priceIncrement"`
		SizeIncrement         float64     `json:"sizeIncrement"`
		MinProvideSize        float64     `json:"minProvideSize"`
		Last                  float64     `json:"last"`
		Bid                   float64     `json:"bid"`
		Ask                   float64     `json:"ask"`
		Price                 float64     `json:"price"`
		Type                  string      `json:"type"`
		BaseCurrency          null.String `json:"baseCurrency"`
		QuoteCurrency         null.String `json:"quoteCurrency"`
		Underlying            string      `json:"underlying"`
		Restricted            bool        `json:"restricted"`
		HighLeverageFeeExempt bool        `json:"highLeverageFeeExempt"`
		Change1H              float64     `json:"change1h"`
		Change24H             float64     `json:"change24h"`
		ChangeBod             float64     `json:"changeBod"`
		QuoteVolume24H        float64     `json:"quoteVolume24h"`
		VolumeUsd24H          float64     `json:"volumeUsd24h"`
		TokenizedEquity       bool        `json:"tokenizedEquity,omitempty"`
	} `json:"result"`
}

type FuturesResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		Name                  string      `json:"name"`
		Underlying            string      `json:"underlying"`
		Description           string      `json:"description"`
		Type                  string      `json:"type"`
		Expiry                interface{} `json:"expiry"`
		Perpetual             bool        `json:"perpetual"`
		Expired               bool        `json:"expired"`
		Enabled               bool        `json:"enabled"`
		PostOnly              bool        `json:"postOnly"`
		PriceIncrement        float64     `json:"priceIncrement"`
		SizeIncrement         float64     `json:"sizeIncrement"`
		Last                  float64     `json:"last"`
		Bid                   float64     `json:"bid"`
		Ask                   float64     `json:"ask"`
		Index                 float64     `json:"index"`
		Mark                  float64     `json:"mark"`
		ImfFactor             float64     `json:"imfFactor"`
		LowerBound            float64     `json:"lowerBound"`
		UpperBound            float64     `json:"upperBound"`
		UnderlyingDescription string      `json:"underlyingDescription"`
		ExpiryDescription     string      `json:"expiryDescription"`
		MoveStart             interface{} `json:"moveStart"`
		MarginPrice           float64     `json:"marginPrice"`
		PositionLimitWeight   float64     `json:"positionLimitWeight"`
		Group                 string      `json:"group"`
		Change1H              float64     `json:"change1h"`
		Change24H             float64     `json:"change24h"`
		ChangeBod             float64     `json:"changeBod"`
		VolumeUsd24H          float64     `json:"volumeUsd24h"`
		Volume                float64     `json:"volume"`
		OpenInterest          float64     `json:"openInterest"`
		OpenInterestUsd       float64     `json:"openInterestUsd"`
	} `json:"result"`
}

type LendingRatesResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		Coin     string  `json:"coin"`
		Previous float64 `json:"previous"`
		Estimate float64 `json:"estimate"`
	} `json:"result"`
}

type FutureStatsResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Ask            int       `json:"ask"`
		Bid            float64   `json:"bid"`
		Change1H       int       `json:"change1h"`
		Change24H      int       `json:"change24h"`
		Description    string    `json:"description"`
		Enabled        bool      `json:"enabled"`
		Expired        bool      `json:"expired"`
		Expiry         time.Time `json:"expiry"`
		Index          float64   `json:"index"`
		Last           int       `json:"last"`
		LowerBound     float64   `json:"lowerBound"`
		Mark           float64   `json:"mark"`
		Name           string    `json:"name"`
		Perpetual      bool      `json:"perpetual"`
		PostOnly       bool      `json:"postOnly"`
		PriceIncrement float64   `json:"priceIncrement"`
		SizeIncrement  float64   `json:"sizeIncrement"`
		Underlying     string    `json:"underlying"`
		UpperBound     float64   `json:"upperBound"`
		Type           string    `json:"type"`
	} `json:"result"`
}

type LendingHistoryResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		Coin string    `json:"coin"`
		Time time.Time `json:"time"`
		Size float64   `json:"size"`
		Rate float64   `json:"rate"`
	} `json:"result"`
}

type HistoricalPricesResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		StartTime time.Time `json:"startTime"`
		Time      float64   `json:"time"`
		Open      float64   `json:"open"`
		High      float64   `json:"high"`
		Low       float64   `json:"low"`
		Close     float64   `json:"close"`
		Volume    float64   `json:"volume"`
	} `json:"result"`
}

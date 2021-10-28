package structs

import (
	"time"

	null "github.com/guregu/null"
)

type OrderHistoryRequest struct {
	Market    null.String `json:"market,omitempty"`
	StartTime null.Float  `json:"start_time,omitempty"`
	EndTime   null.Float  `json:"end_time,omitempty"`
	Limit     null.Int    `json:"limit,omitempty"`
}

type NewOrder struct {
	Market                  string      `json:"market"`
	Side                    string      `json:"side"`
	Price                   null.Float  `json:"price"`
	Type                    string      `json:"type"`
	Size                    float64     `json:"size"`
	ReduceOnly              bool        `json:"reduceOnly"`
	Ioc                     bool        `json:"ioc"`
	PostOnly                bool        `json:"postOnly"`
	ExternalReferralProgram string      `json:"externalReferralProgram"`
	ClientID                null.String `json:"clientId"`
}

type NewOrderResponse struct {
	Success bool  `json:"success"`
	Result  Order `json:"result"`
}

type Order struct {
	CreatedAt     time.Time `json:"createdAt"`
	FilledSize    float64   `json:"filledSize"`
	Future        string    `json:"future"`
	ID            int64     `json:"id"`
	Market        string    `json:"market"`
	Price         float64   `json:"price"`
	AvgFillPrice  float64   `json:"avgFillPrice"`
	RemainingSize float64   `json:"remainingSize"`
	Side          string    `json:"side"`
	Size          float64   `json:"size"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	ReduceOnly    bool      `json:"reduceOnly"`
	Ioc           bool      `json:"ioc"`
	PostOnly      bool      `json:"postOnly"`
	ClientID      string    `json:"clientId"`
}

type OpenOrders struct {
	Success bool    `json:"success"`
	Result  []Order `json:"result"`
}

type OrderHistory struct {
	Success     bool    `json:"success"`
	Result      []Order `json:"result"`
	HasMoreData bool    `json:"hasMoreData"`
}

type NewTriggerOrder struct {
	Market           string  `json:"market"`
	Side             string  `json:"side"`
	Size             float64 `json:"size"`
	Type             string  `json:"type"`
	ReduceOnly       bool    `json:"reduceOnly"`
	RetryUntilFilled bool    `json:"retryUntilFilled"`
	TriggerPrice     float64 `json:"triggerPrice,omitempty"`
	OrderPrice       float64 `json:"orderPrice,omitempty"`
	TrailValue       float64 `json:"trailValue,omitempty"`
}

type NewTriggerOrderResponse struct {
	Success bool         `json:"success"`
	Result  TriggerOrder `json:"result"`
}

type TriggerOrder struct {
	CreatedAt        time.Time `json:"createdAt"`
	Error            string    `json:"error"`
	Future           string    `json:"future"`
	ID               int64     `json:"id"`
	Market           string    `json:"market"`
	OrderID          int64     `json:"orderId"`
	OrderPrice       float64   `json:"orderPrice"`
	ReduceOnly       bool      `json:"reduceOnly"`
	Side             string    `json:"side"`
	Size             float64   `json:"size"`
	Status           string    `json:"status"`
	TrailStart       float64   `json:"trailStart"`
	TrailValue       float64   `json:"trailValue"`
	TriggerPrice     float64   `json:"triggerPrice"`
	TriggeredAt      string    `json:"triggeredAt"`
	Type             string    `json:"type"`
	OrderType        string    `json:"orderType"`
	FilledSize       float64   `json:"filledSize"`
	AvgFillPrice     float64   `json:"avgFillPrice"`
	OrderStatus      string    `json:"orderStatus"`
	RetryUntilFilled bool      `json:"retryUntilFilled"`
}

type OpenTriggerOrders struct {
	Success bool           `json:"success"`
	Result  []TriggerOrder `json:"result"`
}

type TriggerOrderHistory struct {
	Success     bool           `json:"success"`
	Result      []TriggerOrder `json:"result"`
	HasMoreData bool           `json:"hasMoreData"`
}

type Triggers struct {
	Success bool `json:"success"`
	Result  []struct {
		Error      string    `json:"error"`
		FilledSize float64   `json:"filledSize"`
		OrderSize  float64   `json:"orderSize"`
		OrderID    int64     `json:"orderId"`
		Time       time.Time `json:"time"`
	} `json:"result"`
}

type ModifyOrder struct {
	Price    float64 `json:"price,omitempty"`
	Size     float64 `json:"size,omitempty"`
	ClientID string  `json:"client_id,omitempty"`
}

type OrderFill struct {
	Fee           float64     `json:"fee"`
	FeeCurrency   string      `json:"feeCurrency"`
	FeeRate       float64     `json:"feeRate"`
	Future        string      `json:"future"`
	ID            int64       `json:"id"`
	Liquidity     string      `json:"liquidity"`
	Market        string      `json:"market"`
	BaseCurrency  interface{} `json:"baseCurrency"`
	QuoteCurrency interface{} `json:"quoteCurrency"`
	OrderID       int64       `json:"orderId"`
	TradeID       int64       `json:"tradeId"`
	Price         float64     `json:"price"`
	Side          string      `json:"side"`
	Size          float64     `json:"size"`
	Time          time.Time   `json:"time"`
	Type          string      `json:"type"`
}

type FillsResponse struct {
	Success bool        `json:"success"`
	Result  []OrderFill `json:"result"`
}

type Balance struct {
	Coin                   string  `json:"coin"`
	Free                   float64 `json:"free"`
	SpotBorrow             float64 `json:"spotBorrow"`
	Total                  float64 `json:"total"`
	UsdValue               float64 `json:"usdValue"`
	AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
}

type BalancesResponse struct {
	Success bool      `json:"success"`
	Result  []Balance `json:"result"`
}

type CancelResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

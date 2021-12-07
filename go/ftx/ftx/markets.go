package ftx

import (
	"github.com/Investabit/ftx/go/ftx/ftx/structs"
	"net/http"
	// "log"
	"github.com/guregu/null"
	"strconv"
)

type HistoricalPrices structs.HistoricalPrices
type Trades structs.Trades

func (client *FtxClient) GetHistoricalPrices(market string, resolution int64,
	limit int64, startTime int64, endTime int64) (HistoricalPrices, error) {
	var historicalPrices HistoricalPrices
	resp, err := client._get(
		"markets/"+market+
			"/candles?resolution="+strconv.FormatInt(resolution, 10)+
			"&limit="+strconv.FormatInt(limit, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		// log.Printf("Error GetHistoricalPrices", err)
		return historicalPrices, err
	}
	err = _processResponse(resp, &historicalPrices)
	return historicalPrices, err
}

func (client *FtxClient) GetTrades(market string, limit int64, startTime int64, endTime int64) (Trades, error) {
	var trades Trades
	resp, err := client._get(
		"markets/"+market+"/trades?"+
			"&limit="+strconv.FormatInt(limit, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return trades, err
	}
	err = _processResponse(resp, &trades)
	return trades, err
}

func (client *FtxClient) GetOrderbook(market string, depth int64) (structs.Orderbook, *http.Response, error) {
	var orderbook structs.Orderbook
	resp, err := client._get(
		"markets/"+market+"/orderbook?"+
			"&depth="+strconv.FormatInt(depth, 10),
		[]byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return orderbook, resp, err
	}
	err = _processResponse(resp, &orderbook)
	return orderbook, resp, err
}

func (client *FtxClient) Markets() (structs.MarketsResponse, *http.Response, error) {
	var markets_resp structs.MarketsResponse
	resp, err := client._get("markets", []byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return markets_resp, resp, err
	}
	err = _processResponse(resp, &markets_resp)
	return markets_resp, resp, err
}

func (client *FtxClient) Futures() (structs.FuturesResponse, *http.Response, error) {
	var future_resp structs.FuturesResponse
	resp, err := client._get("futures", []byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return future_resp, resp, err
	}
	err = _processResponse(resp, &future_resp)
	return future_resp, resp, err
}

func (client *FtxClient) FutureStat(future string) (structs.FutureStatsResponse, *http.Response, error) {
	var future_resp structs.FutureStatsResponse
	resp, err := client._get("futures/"+future+"/stats", []byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return future_resp, resp, err
	}
	err = _processResponse(resp, &future_resp)
	return future_resp, resp, err
}

func (client *FtxClient) LendingRates() (structs.LendingRatesResponse, *http.Response, error) {
	var lendingRates structs.LendingRatesResponse
	resp, err := client._get("spot_margin/lending_rates", []byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return lendingRates, resp, err
	}
	err = _processResponse(resp, &lendingRates)
	return lendingRates, resp, err
}

func (client *FtxClient) LendingHistory(startTime null.Int, endTime null.Int) (structs.LendingHistoryResponse, *http.Response, error) {
	var lendingRateHistory structs.LendingHistoryResponse
	params := ""
	if !startTime.IsZero() {
		params = "?start_time=" + strconv.FormatInt(startTime.ValueOrZero(), 10)
	}
	if !endTime.IsZero() {
		if len(params) == 0 {
			params = "?"
		} else {
			params += "&"
		}
		params += "end_time=" + strconv.FormatInt(endTime.ValueOrZero(), 10)
	}
	resp, err := client._get("spot_margin/history"+params, []byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return lendingRateHistory, resp, err
	}
	err = _processResponse(resp, &lendingRateHistory)
	return lendingRateHistory, resp, err
}

func (client *FtxClient) HistoricalPrices(market string, resolution string, startTime null.Int, endTime null.Int) (structs.HistoricalPricesResponse, *http.Response, error) {
	var historyPrice structs.HistoricalPricesResponse
	params := ""
	if !startTime.IsZero() {
		params = "&start_time=" + strconv.FormatInt(startTime.ValueOrZero(), 10)
	}
	if !endTime.IsZero() {
		params += "&end_time=" + strconv.FormatInt(endTime.ValueOrZero(), 10)
	}
	resp, err := client._get("markets/"+market+"/candles?resolution="+
		resolution+params, []byte(""))
	if err != nil {
		// log.Printf("Error GetTrades", err)
		return historyPrice, resp, err
	}
	err = _processResponse(resp, &historyPrice)
	return historyPrice, resp, err
}

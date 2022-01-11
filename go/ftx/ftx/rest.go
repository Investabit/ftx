package ftx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	// "log"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

const URL = "https://ftx.com/api/"

func (client *FtxClient) signRequest(method string, path string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, URL+path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if len(client.Api) == 0 || len(client.Secret) == 0 {
		ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
		signaturePayload := ts + method + "/api/" + path + string(body)
		signature := client.sign(signaturePayload)

		req.Header.Set("FTX-KEY", client.Api)
		req.Header.Set("FTX-SIGN", signature)
		req.Header.Set("FTX-TS", ts)
		if client.Subaccount != "" {
			req.Header.Set("FTX-SUBACCOUNT", client.Subaccount)
		}
	}
	return req
}

func (client *FtxClient) _get(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("GET", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *FtxClient) _post(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("POST", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *FtxClient) _delete(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("DELETE", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func _processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Printf("Error processing response:", err)
		return err
	}

	result := gjson.GetBytes(body, "success")
	if !result.Bool() {
		resErr := gjson.GetBytes(body, "error")
		resStr := resErr.String()
		if resStr == "" {
			return errors.New(resStr)
		}
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		// log.Printf("Error processing response:", err)
		return err
	}

	return nil
}

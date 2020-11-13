package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func newParap() *Param {
	return &Param{}
}

func (u *updateCfg) value() url.Values {
	val := url.Values{}
	if u.Offset != 0 {
		val.Add("offset", strconv.Itoa(u.Offset))
	}
	if u.Limit != 0 {
		val.Add("limit", strconv.Itoa(u.Limit))
	} else {
		val.Add("limit", strconv.Itoa(10))
	}
	val.Add("timeout", strconv.Itoa(u.Timeout))
	if u.AllowedUpdates != nil {
		for _, allow := range u.AllowedUpdates {
			val.Add("allowed_updates", allow)
		}
	}
	return val
}

func decodeResp(respBody io.Reader, resp interface{}) error {
	data, err := ioutil.ReadAll(respBody)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	return nil
}

// MakeRequest создаём простой запрос
func (b *bot) MakeRequest(endpoint string, params url.Values) (apiResponse, error) {
	apiResp := apiResponse{}
	method := fmt.Sprintf(b.Endpoint, b.Token, endpoint)
	req, err := http.NewRequest("POST", method, strings.NewReader(params.Encode()))
	if err != nil {
		return apiResp, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := b.client.Do(req)
	if err != nil {
		return apiResp, err
	}
	defer resp.Body.Close()

	err = decodeResp(resp.Body, &apiResp)
	if err != nil {
		return apiResp, err
	}
	if apiResp.OK != true {
		return apiResp, fmt.Errorf("not OK")
	}
	return apiResp, nil
}

// MakeJSONRequest создаём запрос с JSON параметрами
func (b *bot) MakeJSONRequest(end string, params []byte) (apiResponse, error) {
	apiResp := apiResponse{}
	method := fmt.Sprintf(endpoint, b.Token, end)
	reader := bytes.NewBuffer(params)
	req, err := http.NewRequest("POST", method, reader)

	if err != nil {
		return apiResp, err
	}
	fmt.Println(reader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return apiResp, err
	}
	defer resp.Body.Close()

	err = decodeResp(resp.Body, &apiResp)
	if err != nil {
		return apiResp, err
	}
	if apiResp.OK != true {
		return apiResp, nil
	}
	return apiResp, nil
}

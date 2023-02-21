package ratp

import (
	"net/url"
	"io/ioutil"
	"net/http"
	"github.com/mstch/go-ratp/logger"
)

func createUrl(functionality, location string) {
	URL.Path += functionality
	URL.RawQuery = "MonitoringRef=" + url.QueryEscape(location)
}

func sendRequest() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		logger.Error("creation of GET request failed.")
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("apikey", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("request failed")
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("ioutil.Readall fail")
		return nil, err
	}

	return data, nil
}

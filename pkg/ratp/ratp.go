package ratp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/mstch/go-ratp/logger"
)


type Stop struct {
	Location 	string
	Command 	string
	Result		NextPassResult
}

func (s *Stop) Request(location string) error {
	s.Location = location
	s.Command = NextPass
	createUrl(s.Command, s.Location)
	req, err :=sendRequest()
	if err != nil { return err }
	s.Result, _ = parseResult(req)
	return nil
}

func (s Stop) String() (string) {
	result := fmt.Sprintf("%s:\n", s.Result.DestinationName)
	location, _ := time.LoadLocation(Location)
	layout := "15:04:05"
	for _, v := range s.Result.ExpectedDeparture {
		result += fmt.Sprintf("%s\n", v.In(location).Format(layout))
	}
	return result
}

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

func parseResult(body []byte) (NextPassResult, error) {
	npStruct := NextPassStruct{}
	err := json.Unmarshal(body, &npStruct)
	if err != nil {
		logger.Error(err)
		return NextPassResult{}, err
	}
	if len(npStruct.Siri.ServiceDelivery.StopMonitoringDelivery) == 0 ||
		len(npStruct.Siri.ServiceDelivery.StopMonitoringDelivery[0].MonitoredStopVisit) == 0{
		err := errors.New("incomplete answer")
		logger.Error(err)
		return NextPassResult{}, err
	}

	base := npStruct.Siri.ServiceDelivery.StopMonitoringDelivery[0].MonitoredStopVisit
	result := NextPassResult{}
	result.MonitoringRef = base[0].MonitoringRef.Value
	result.DestinationName = base[0].MonitoredVehicleJourney.DestinationRef.Value

	for _, v := range base {
		result.ExpectedArrival = append(result.ExpectedArrival, v.MonitoredVehicleJourney.MonitoredCall.ExpectedArrivalTime)
		result.ExpectedDeparture = append(result.ExpectedDeparture, v.MonitoredVehicleJourney.MonitoredCall.ExpectedDepartureTime)
	}
	return result, nil
}

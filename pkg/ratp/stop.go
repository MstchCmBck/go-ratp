package ratp

import (
	"encoding/json"
	"errors"
	"fmt"
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
	result := fmt.Sprintf("%s:\n", s.Result.StopName)
	location, _ := time.LoadLocation(Location)
	layout := "15:04:05"
	for _, v := range s.Result.ExpectedDeparture {
		timeToWait := int(v.Sub(time.Now()).Minutes())
		result += fmt.Sprintf("\t%s (%d min)\n", v.In(location).Format(layout), timeToWait)
	}
	return result
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
	result.StopName = base[0].MonitoredVehicleJourney.DestinationName[0].Value

	for _, v := range base {
		result.ExpectedArrival = append(result.ExpectedArrival, v.MonitoredVehicleJourney.MonitoredCall.ExpectedArrivalTime)
		result.ExpectedDeparture = append(result.ExpectedDeparture, v.MonitoredVehicleJourney.MonitoredCall.ExpectedDepartureTime)
	}
	return result, nil
}

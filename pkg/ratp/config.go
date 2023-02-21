package ratp


import (
	"time"
	"net/url"
)

const token string = "YOUR-TOKEN"

// List of API function.
const (
	NextPass 			= "/stop-monitoring"
	DisplayedMessage 	= "/general-message"
)

// List of stop.
const (
	Foo = "STIF:StopPoint:Q:22389:"
)

var URL url.URL = url.URL{
	Scheme 	: "https",
	Host 	: "prim.iledefrance-mobilites.fr",
	Path    : "/marketplace",
}

const Location string = "Europe/Paris"

type NextPassResult struct {
	MonitoringRef		string
	DestinationName		string
	DepartureStatus		bool
	ExpectedArrival 	[]time.Time
	ExpectedDeparture 	[]time.Time
}

type NextPassStruct struct {
	Siri struct {
		ServiceDelivery struct {
			ResponseTimestamp         time.Time `json:"ResponseTimestamp"`
			ProducerRef               string    `json:"ProducerRef"`
			ResponseMessageIdentifier string    `json:"ResponseMessageIdentifier"`
			StopMonitoringDelivery    []struct {
				ResponseTimestamp  time.Time `json:"ResponseTimestamp"`
				Version            string    `json:"Version"`
				Status             string    `json:"Status"`
				MonitoredStopVisit []struct {
					RecordedAtTime time.Time `json:"RecordedAtTime"`
					ItemIdentifier string    `json:"ItemIdentifier"`
					MonitoringRef  struct {
						Value string `json:"value"`
					} `json:"MonitoringRef"`
					MonitoredVehicleJourney struct {
						LineRef struct {
							Value string `json:"value"`
						} `json:"LineRef"`
						OperatorRef struct {
							Value string `json:"value"`
						} `json:"OperatorRef"`
						FramedVehicleJourneyRef struct {
							DataFrameRef struct {
								Value string `json:"value"`
							} `json:"DataFrameRef"`
							DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef"`
						} `json:"FramedVehicleJourneyRef"`
						DirectionName []struct {
							Value string `json:"value"`
						} `json:"DirectionName"`
						DestinationRef struct {
							Value string `json:"value"`
						} `json:"DestinationRef"`
						DestinationName []struct {
							Value string `json:"value"`
						} `json:"DestinationName"`
						VehicleJourneyName []interface{} `json:"VehicleJourneyName"`
						JourneyNote        []interface{} `json:"JourneyNote"`
						MonitoredCall      struct {
							StopPointName []struct {
								Value string `json:"value"`
							} `json:"StopPointName"`
							VehicleAtStop      bool `json:"VehicleAtStop"`
							DestinationDisplay []struct {
								Value string `json:"value"`
							} `json:"DestinationDisplay"`
							ExpectedArrivalTime   time.Time `json:"ExpectedArrivalTime"`
							ExpectedDepartureTime time.Time `json:"ExpectedDepartureTime"`
							DepartureStatus       string    `json:"DepartureStatus"`
							ArrivalStatus         string    `json:"ArrivalStatus"`
						} `json:"MonitoredCall"`
						TrainNumbers struct {
							TrainNumberRef []interface{} `json:"TrainNumberRef"`
						} `json:"TrainNumbers"`
					} `json:"MonitoredVehicleJourney"`
				} `json:"MonitoredStopVisit"`
			} `json:"StopMonitoringDelivery"`
		} `json:"ServiceDelivery"`
	} `json:"Siri"`
}


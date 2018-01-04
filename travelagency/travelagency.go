package travelagency

import (
	"time"
)

// Weather .
type Weather struct{}

// Destinations .
type Destinations struct{}

// Quoting .
type Quoting struct{}

//Customers .
type Customers struct{}

// Services ..
type Services struct {
	WeatherService     Weather
	DestinationService Destinations
	QuotingService     Quoting
}

// GetWeatherService .
func GetWeatherService(d Destinations) Weather {
	time.Sleep(330 * time.Millisecond)
	return Weather{}
}

// GetDestinationService .
func GetDestinationService(c Customers) [10]Destinations {
	time.Sleep(250 * time.Millisecond)
	return [10]Destinations{}
}

// GetQuotingService .
func GetQuotingService(d Destinations) Quoting {
	time.Sleep(170 * time.Millisecond)
	return Quoting{}
}

// GetCustomerService ..
func GetCustomerService() Customers {
	time.Sleep(150 * time.Millisecond)
	return Customers{}
}

package synchronous

import (
	"Asynchronous-HTTPClient/travelagency"
)

// ComputeSynchronous -- compute the time that synchronous task cost
func ComputeSynchronous() [10]travelagency.Services {
	var services [10]travelagency.Services

	// get authenticated customer
	customer := travelagency.GetCustomerService()
	// get destination-list
	destination := travelagency.GetDestinationService(customer)

	for index, item := range destination {
		weatherservice := travelagency.GetWeatherService(item)
		quotingservice := travelagency.GetQuotingService(item)

		services[index] = travelagency.Services{weatherservice, item, quotingservice}
	}
	return services
}

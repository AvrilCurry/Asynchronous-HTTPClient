package asynchronous

import "Asynchronous-HTTPClient/travelagency"

// ComputeAsynchronous -- compute the time that asynchronous task cost
func ComputeAsynchronous() [10]travelagency.Services {
	var services [10]travelagency.Services

	// get authenticated customer
	customer := travelagency.GetCustomerService()
	// get destination-list
	destination := travelagency.GetDestinationService(customer)

	quotingData := [10]chan struct{}{}
	weatherData := [10]chan struct{}{}

	for index := 0; index < 10; index++ {
		weatherData[index] = make(chan struct{}, 0)
		quotingData[index] = make(chan struct{}, 0)
	}
	for index, item := range destination {
		var _index = index
		var _item = item
		go func() {
			services[_index].WeatherService = travelagency.GetWeatherService(_item)
			weatherData[_index] <- struct{}{} // 发送数据
		}()
		go func() {
			services[_index].QuotingService = travelagency.GetQuotingService(_item)
			quotingData[_index] <- struct{}{} // 发送数据
		}()
	}
	for _, data := range weatherData {
		<-data // 接收数据
	}
	for _, data := range quotingData {
		<-data // 接收数据
	}
	return services
}

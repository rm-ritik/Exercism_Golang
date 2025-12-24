// Package weather provides tools to forecast weater.
package weather

var (
    // CurrentCondition consists of a string representing current condition.
	CurrentCondition string
    // CurrentLocation consists of a string representing current location.
	CurrentLocation  string
)

// Forecast returns the string representing the current forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

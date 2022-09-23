// Package weather provides to forecast weather conditions
// via current condition and current location.
package weather

// CurrentCondition represents a your current weather condition.
var CurrentCondition string

// CurrentLocation represents a your current location.
var CurrentLocation string

// Forecast simply return current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

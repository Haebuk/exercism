package cars

import "math"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(math.Floor(CalculateWorkingCarsPerHour(productionRate, successRate) / 60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	num_of_groups := carsCount / 10
	remain_car := carsCount % 10
	return uint(num_of_groups)*95_000 + uint(remain_car)*10_000
}

package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

    successCarsPerHour := (float64(productionRate) * successRate)/100
    
    return successCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	 successCarsPerHours := (float64(productionRate) * successRate)/100
    successCarsPerMinutes := successCarsPerHours/60

    return int(successCarsPerMinutes)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

    groupOfTen := carsCount / 10
    restOfTen := carsCount - (groupOfTen * 10)

    return uint(groupOfTen * 95000) + uint(restOfTen * 10000)
}

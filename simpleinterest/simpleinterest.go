package simpleinterest

import (
	"fmt"
)

// init function added
func init() {
	fmt.Println("Simple interest package initialized")

}

// Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
func Calculate(principal float64, rate float64, duration float64) float64 {
	interest := principal * (rate / 100) * duration
	return interest
}

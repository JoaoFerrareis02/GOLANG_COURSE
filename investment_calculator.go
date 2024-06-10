package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// Formating strings - basic

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)

	// Formating strings

	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	// Creating strings

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// Building multiline strings

	// fmt.Printf(`Future Value: %.2f
	// Future Value (adjusted for Inflation): %.2f`, futureValue, futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

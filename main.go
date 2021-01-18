package main

import (
	"fmt"
	"math/rand"
)

func main() {
	daysOfCalculation := 2600
	losses := 7
	profits := 42
	breakevens := 9
	profitPercentage := profits * 100 / (profits + losses + breakevens)
	averageLossInPips := -14.5
	for i := 0; i < 10; i++ {
		averageWinInPips := 18.0
		var averages []float64
		for i := 0; i < daysOfCalculation; i++ {
			var numbers []float64
			numbers = append(numbers, 0.0)
			latestNumber := 0.0

			for i := 0; i < daysOfCalculation; i++ {
				randomProduction := rand.Intn(100-0) + 0
				if randomProduction < profitPercentage {
					latestNumber += averageWinInPips
				} else {
					latestNumber += averageLossInPips
				}
				numbers = append(numbers, latestNumber)
			}
			//fmt.Println(numbers[daysOfCalculation])
			averages = append(averages, numbers[daysOfCalculation])
		}
		sum := 0.0
		for _, average := range averages {
			sum += average
		}
		fmt.Println("")
		fmt.Print("Total average sum of pips after set days of calculation: ")
		fmt.Println(int(sum / float64(daysOfCalculation)))
		fmt.Print("Average daily sum of pips after set days of calculation: ")
		fmt.Printf("%.2f pips\n", sum/float64(daysOfCalculation)/float64(daysOfCalculation))
		fmt.Print("Daily percentage estimate after set days of calculation: ")
		fmt.Printf("%.2f %%\n", 3.79/18.0*(sum/float64(daysOfCalculation)/float64(daysOfCalculation)-1))
	}

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numberOfTrades := 100
	spreadWithCommission := 1.1
	totalSpread := float64(numberOfTrades) * spreadWithCommission
	losses := 40
	wins := 60
	profitPercentage := wins * 100 / (wins + losses)
	averageLossInPips := -26.5
	averageWinInPips := 28.3
	numberOfSimulations := 1000
	randomDrawDown := rand.Intn(numberOfTrades-1) + 1
	randomDrawDownInPips := -averageWinInPips * 2
	randomSlippage := rand.Intn(numberOfTrades-1) + 1
	randomSlippageInPips := -averageWinInPips / 2
	randomBadDay := rand.Intn(numberOfTrades-1) + 1

	var listOfResults []float64

	for i := 0; i < numberOfSimulations; i++ {
		fmt.Print("Monte Carlo Simulation No.      ")
		fmt.Printf("%d\n", i+1)
		if randomBadDay == rand.Intn(numberOfTrades-1)+1 {
			profitPercentage = 100 - profitPercentage
			fmt.Print("B")
		}
		sum := 0.0
		profitsCount := 0
		lossesCount := 0
		for i := 0; i < numberOfTrades; i++ {
			randomProduction := rand.Intn(numberOfTrades-1) + 1
			if randomProduction < profitPercentage {
				fmt.Print("+")
				profitsCount++
				sum += averageWinInPips
				if randomProduction == randomSlippage {
					sum += randomSlippageInPips
					fmt.Print("S")
				}
			} else {
				fmt.Print("-")
				lossesCount++
				sum += averageLossInPips
				if randomProduction == randomSlippage {
					sum += randomSlippageInPips
					fmt.Print("S")
				}
			}
		}
		if randomDrawDown == rand.Intn(numberOfTrades-1)+1 {
			sum += randomDrawDownInPips
			fmt.Print("+!!")
		}
		fmt.Println("")
		sum -= totalSpread
		fmt.Print("Number of profits               ")
		fmt.Printf("%d\n", profitsCount)
		fmt.Print("Number of losses                ")
		fmt.Printf("%d\n", lossesCount)
		fmt.Print("Total sum of pips               ")
		fmt.Printf("%.1f pips\n", sum)
		fmt.Print("Average trade will give you     ")
		fmt.Printf("%.1f pips\n", sum/float64(numberOfTrades))
		fmt.Println("")
		listOfResults = append(listOfResults, sum)
	}
	totalSum := 0.0
	min := 0.0
	max := 0.0
	for i, result := range listOfResults {
		if i == 0 {
			min = result
			max = result
		}
		totalSum += result
		if result < min {
			min = result
		}
		if result > max {
			max = result
		}
	}
	average := totalSum / float64(len(listOfResults))
	fmt.Println("")
	fmt.Println("")
	fmt.Print("Average sum                         ")
	fmt.Printf("%.1f pips\n", average)
	fmt.Print("Min sum                             ")
	fmt.Printf("%.1f pips\n", min)
	fmt.Print("Max sum                             ")
	fmt.Printf("%.1f pips\n", max)
}

package main

import (
	"fmt"
	"math/rand"
)

func SpinSlotMachine(reel []string, r int, c int) [][]string {
	result := [][]string{}

	for i := 0; i < r; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < c; col++ {
		selected := map[int]bool{}
		for row := 0; row < r; row++ {
			for {
				randomIndex := rand.Intn(len(reel)-1) + 0
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}

	return result
}

func DisplayWheels(spin [][]string) {
	for _, row := range spin {
		fmt.Println()
		for j, symbol := range row {
			fmt.Print(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
	}
	fmt.Println()
}

func CheckWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		currSymbol := row[0]
		for _, symbol := range row[1:] {
			if currSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[currSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

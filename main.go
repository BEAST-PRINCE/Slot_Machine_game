package main

import (
	"fmt"
)

func main() {
	symbols := map[string]uint{
		"A": 6,
		"B": 10,
		"C": 12,
		"D": 22,
		"X": 3,
	}

	multiplier := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 1,
		"X": 50,
	}

	name := GetName()

	balance := uint(250)

	symbolArr := GenerateSymbols(symbols)

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		// symbolArr := generateSymbols(symbols)
		spin := SpinSlotMachine(symbolArr, 3, 3)
		DisplayWheels(spin)
		winningLines := CheckWin(spin, multiplier)
		maxWin := uint(0)

		for _, i := range winningLines {
			if i > maxWin {
				maxWin = i
			}
		}

		if maxWin == 0 {
			fmt.Println("\nHard Luck!! You lost your bet money.")
		} else {
			balance = balance + (maxWin * bet)
			fmt.Printf("\nWoohhooooo!! You won %dx your bet. You got %d \n", maxWin, maxWin*bet)
			fmt.Printf("Your updated balance is %d\n", balance)
		}

	}

	if balance == 0 {
		fmt.Println()
		fmt.Println("You don't have enough money to play again.")
	} else {
		fmt.Printf("\n You left with %d.\n", balance)
	}
	fmt.Printf("Hope you had a great time. %s :)\n", name)
	fmt.Printf("\n\n\n")
}

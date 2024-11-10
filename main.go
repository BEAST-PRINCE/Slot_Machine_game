package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	name := ""

	fmt.Println("\nWelcome to Prince' Casino!!!")
	fmt.Println("Let's Start Gambling!!")
	fmt.Printf("\nPlease enter your name: ")

	if _, err := fmt.Scanln(&name); err != nil {
		return ""
	}
	fmt.Printf("\nWelcome %s, HAppy to have you here!\nLets Play!\n", name)

	return name
}

func getBet(balance uint) uint {
	var bet uint

	for {
		fmt.Printf("\n Enter your bet or 0 to quit (balance = %d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("\nBet cannot be larger than the current balance")
		} else {
			break
		}

	}
	return bet
}

func generateSymbols(symbols map[string]uint) []string {
	arr := []string{}
	for symbol, cnt := range symbols {
		for i := 0; i < int(cnt); i++ {
			arr = append(arr, symbol)
		}
	}

	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	// fmt.Println(arr)
	return arr
}

func spinSlotMachine(reel []string, r int, c int) [][]string {
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

func displayWheels(spin [][]string) {
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

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
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

	name := getName()

	balance := uint(250)

	symbolArr := generateSymbols(symbols)

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		// symbolArr := generateSymbols(symbols)
		spin := spinSlotMachine(symbolArr, 3, 3)
		displayWheels(spin)
		winningLines := checkWin(spin, multiplier)
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

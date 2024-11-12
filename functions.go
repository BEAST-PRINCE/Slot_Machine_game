package main

import (
	"fmt"
	"math/rand"
)

func GetName() string {
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

func GetBet(balance uint) uint {
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

func GenerateSymbols(symbols map[string]uint) []string {
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

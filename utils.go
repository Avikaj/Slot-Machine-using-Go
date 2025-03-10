package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to the Game")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, Let's play!\n", name)
	return name

}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet or zero to quit (balance =%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bets cannot be larger than balance.")

		} else {
			break
		}

	}
	return bet
}

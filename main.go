package main

import "fmt"

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}
	for _, row := range spin {
		win := true
		checksymbol := row[0]
		for _, symbol := range row[1:] {
			if checksymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checksymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 6,
		"D": 2,
	}
	symbolArr := GenerateSymbolArray(symbols)
	balance := uint(200)
	GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		winningLines := checkWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("You won $%d, (%dx) on Line #%d\n", win, multi, i+1)
				fmt.Println("Check your updated balance!")
			}
		}
	}
	fmt.Printf("You are left with, $%d\n.", balance)
}

# Slot-Machine-in-Go

## 📜 Description
This is a simple command-line Slot Machine Game written in Go. Players can place bets, spin the reels, and win based on matching symbols. The game manages user balance, generates random spins, and applies multipliers for winnings.

## 🚀 Features
Interactive command-line gameplay.
Randomized slot machine spins using Go’s math/rand package.
Betting system with balance management.
Different symbols with unique multipliers.
Win detection with dynamic payouts.

## 📁 Project Structure
`main.go` - Controls the game loop and win-checking logic.

`spin.go` - Handles slot machine spins and symbol generation.

`utils.go` - Manages user input for betting and player name.

## To Run the game:
Run:
`go run main.go spin.go utils.go`

OR

Run: `go mod init _name_`  to create a .mod file package and then run `go run .`


## 🎮 How to Play
1. Enter your name when prompted.
2. Place your bet (must be within your balance).
3. Spin the reels and check if you win!
4. The game continues until your balance reaches 0 or you decide to quit.

## 🖥️ Gameplay Example
```
Welcome to the Game

Enter your name: James

Welcome James, Let's play!

Enter your bet or zero to quit (balance =200): 10

A | B | C

C | A | A

B | C | B

You are left with, $190.
```



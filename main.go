package main

import (
	"fmt"
	"MyApp/match"
)

func main() {
	playerA, err := match.NewPlayer("Player A", 50, 5, 10)
	if err != nil {
		fmt.Println("Error creating Player A:", err)
		return
	}

	playerB, err := match.NewPlayer("Player B", 100, 10, 5)
	if err != nil {
		fmt.Println("Error creating Player B:", err)
		return
	}

	result := match.Fight(playerA, playerB)
	fmt.Println(result)
}

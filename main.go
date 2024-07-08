package main

import (
	"MyApp/match"
	"fmt"
)

func main() {
	playerA := match.NewPlayer("Player A", 50, 5, 10)
	playerB := match.NewPlayer("Player B", 100, 10, 5)

	result := match.Fight(playerA, playerB)
	fmt.Println(result)
}

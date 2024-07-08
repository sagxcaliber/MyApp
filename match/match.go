package arena

import (
	"fmt"
	"math/rand"
	"time"
)

// Player represents a player in the arena
type Player struct {
	Name    string
	Health  int
	Strength int
	Attack   int
}

// NewPlayer creates a new player with given attributes
func NewPlayer(name string, health, strength, attack int) *Player {
	return &Player{Name: name, Health: health, Strength: strength, Attack: attack}
}

// RollDie simulates rolling a six-sided die
func RollDie() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

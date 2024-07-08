package match

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


// Fight simulates a fight between two players and returns the result
func Fight(playerA, playerB *Player) string {
	for playerA.Health > 0 && playerB.Health > 0 {
		attackFirst, defendFirst := decideTurn(playerA, playerB)
		attack(attackFirst, defendFirst)
		if defendFirst.Health <= 0 {
			return fmt.Sprintf("%s wins!", attackFirst.Name)
		}
		attack(defendFirst, attackFirst)
		if attackFirst.Health <= 0 {
			return fmt.Sprintf("%s wins!", defendFirst.Name)
		}
	}
	return "It's a draw!"
}

// decideTurn decides which player attacks first based on health
func decideTurn(playerA, playerB *Player) (*Player, *Player) {
	if playerA.Health < playerB.Health {
		return playerA, playerB
	}
	return playerB, playerA
}

// attack performs an attack from the attacker to the defender
func attack(attacker, defender *Player) {
	attackRoll := RollDie()
	defendRoll := RollDie()

	attackDamage := attacker.Attack * attackRoll
	defendStrength := defender.Strength * defendRoll

	damage := attackDamage - defendStrength
	if damage > 0 {
		defender.Health -= damage
	}
	fmt.Printf("%s attacks (%d) and %s defends (%d) - Damage: %d, %s Health: %d\n",
		attacker.Name, attackRoll, defender.Name, defendRoll, damage, defender.Name, defender.Health)
}

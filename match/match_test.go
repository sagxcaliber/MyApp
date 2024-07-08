package match

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	player, err := NewPlayer("Test Player", 100, 10, 20)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if player.Name != "Test Player" || player.Health != 100 || player.Strength != 10 || player.Attack != 20 {
		t.Errorf("Player creation failed: %+v", player)
	}
}

func TestNewPlayerInvalidAttributes(t *testing.T) {
	_, err := NewPlayer("Invalid Player", -100, 10, 20)
	if err == nil {
		t.Error("Expected error for negative health, got none")
	}

	_, err = NewPlayer("Invalid Player", 100, 0, 20)
	if err == nil {
		t.Error("Expected error for zero strength, got none")
	}

	_, err = NewPlayer("Invalid Player", 100, 10, -5)
	if err == nil {
		t.Error("Expected error for negative attack, got none")
	}
}

func TestFight(t *testing.T) {
	playerA, _ := NewPlayer("Player A", 50, 5, 10)
	playerB, _ := NewPlayer("Player B", 100, 10, 5)

	result := Fight(playerA, playerB)
	if result == "" {
		t.Errorf("Expected a winner, got empty result")
	}
}

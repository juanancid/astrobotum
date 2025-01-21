package systems

import (
	"fmt"
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// HealthSystem manages entity health and triggers game-over logic.
type HealthSystem struct {
	PlayerEntity ecs.Entity // The player entity
	GameOver     bool       // Tracks whether the game is over
}

func (hs *HealthSystem) Update(w *ecs.World, dt float64) {
	// Get the player's health component
	health := w.GetComponent(hs.PlayerEntity, reflect.TypeOf(&components.Health{})).(*components.Health)

	// Check if the player's health has reached zero
	if health.CurrentHealth <= 0 {
		hs.GameOver = true
		fmt.Println("Game Over!")
		return
	}
}

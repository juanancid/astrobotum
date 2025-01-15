package systems

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
)

// InputSystem handles player input to update velocity components.
type InputSystem struct{}

func (is *InputSystem) Update(w *ecs.World, dt float64) {
	// Get all entities with the PlayerControlled component
	playerControlledEntities := w.GetComponents(reflect.TypeOf(&components.PlayerControlled{}))

	for entity := range playerControlledEntities {
		// Ensure the entity has a Velocity component
		velocity := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)

		// Reset velocity
		velocity.DX = 0
		velocity.DY = 0

		// Update velocity based on key presses
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			velocity.DX = -100 // Move left
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			velocity.DX = 100 // Move right
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			velocity.DY = -100 // Move up
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			velocity.DY = 100 // Move down
		}
	}
}

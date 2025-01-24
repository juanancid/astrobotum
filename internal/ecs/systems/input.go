package systems

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
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

		// Update velocity based on key presses
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			velocity.DX = -100 // Move left
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			velocity.DX = 100 // Move right
		}
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			onGround := w.GetComponent(entity, reflect.TypeOf(&components.OnGround{})).(*components.OnGround)

			if onGround.IsGrounded { // Only allow jumping if the player is on the ground
				velocity.DY = -50 // Apply an upward velocity
			}
		}
	}
}

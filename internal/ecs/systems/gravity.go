package systems

import (
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// GravitySystem applies gravity to entities with a Velocity component.
type GravitySystem struct {
	Gravity float64 // Gravity constant (e.g., 9.8 or a game-appropriate value)
}

func (gs *GravitySystem) Update(w *ecs.World, dt float64) {
	velocities := w.GetComponents(reflect.TypeOf(&components.Velocity{}))

	for _, component := range velocities {
		velocity := component.(*components.Velocity)

		// Apply gravity to the Y-axis velocity
		velocity.DY += gs.Gravity * dt
	}
}

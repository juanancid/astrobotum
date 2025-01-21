package systems

import (
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// DynamicObstacleMovementSystem updates the position of dynamic obstacles based on velocity.
type DynamicObstacleMovementSystem struct{}

func (d *DynamicObstacleMovementSystem) Update(w *ecs.World, dt float64) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	velocities := w.GetComponents(reflect.TypeOf(&components.Velocity{}))
	dynamicObstacles := w.GetComponents(reflect.TypeOf(&components.DynamicObstacle{}))

	for entity := range dynamicObstacles {
		// Ensure the entity has Position and Velocity components
		pos := positions[entity].(*components.Position)
		vel := velocities[entity].(*components.Velocity)

		// Update position based on velocity
		pos.X += vel.DX * dt
		pos.Y += vel.DY * dt
	}
}

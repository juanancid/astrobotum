package systems

import (
	"reflect"

	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
)

// BoundarySystem enforces screen boundaries for entities with a Position component.
type BoundarySystem struct {
	ScreenWidth, ScreenHeight float64 // Dimensions of the screen
}

func (bs *BoundarySystem) Update(w *ecs.World, dt float64) {
	for _, component := range w.GetComponents(reflect.TypeOf(&components.Position{})) {
		position := component.(*components.Position)

		// Clamp X position within screen boundaries
		if position.X < 0 {
			position.X = 0
		}
		if position.X > bs.ScreenWidth-16 { // Adjust for entity width
			position.X = bs.ScreenWidth - 16
		}

		// Clamp Y position within screen boundaries
		if position.Y < 0 {
			position.Y = 0
		}
		if position.Y > bs.ScreenHeight-16 { // Adjust for entity height
			position.Y = bs.ScreenHeight - 16
		}
	}
}

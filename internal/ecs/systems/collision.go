package systems

import (
	"reflect"

	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
)

type CollisionSystem struct{}

func (cs *CollisionSystem) Update(w *ecs.World, dt float64) {
	// Get all entities with Position and Size components
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))

	// Iterate over all entity pairs to check for collisions
	for e1, pos1 := range positions {
		size1 := sizes[e1].(*components.Size)
		for e2, pos2 := range positions {
			// Skip self-collision
			if e1 == e2 {
				continue
			}

			size2 := sizes[e2].(*components.Size)

			// Check for overlap (basic AABB collision detection)
			if isColliding(
				pos1.(*components.Position), size1,
				pos2.(*components.Position), size2,
			) {
				// For now, log collisions
				println("Collision detected between entities", e1, "and", e2)
			}
		}
	}
}

// isColliding checks if two entities' bounding boxes overlap.
func isColliding(pos1 *components.Position, size1 *components.Size, pos2 *components.Position, size2 *components.Size) bool {
	return pos1.X < pos2.X+size2.Width &&
		pos1.X+size1.Width > pos2.X &&
		pos1.Y < pos2.Y+size2.Height &&
		pos1.Y+size1.Height > pos2.Y
}

package systems

import (
	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
	"reflect"
)

// CollisionSystem checks for overlaps and handles collision responses.
type CollisionSystem struct {
	previousPositions map[ecs.Entity]*components.Position
}

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{
		previousPositions: make(map[ecs.Entity]*components.Position),
	}
}

// SavePreviousPositions stores the positions of all entities before movement.
func (cs *CollisionSystem) SavePreviousPositions(w *ecs.World) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))

	for entity, pos := range positions {
		cs.previousPositions[entity] = &components.Position{
			X: pos.(*components.Position).X,
			Y: pos.(*components.Position).Y,
		}
	}
}

func (cs *CollisionSystem) Update(w *ecs.World, dt float64) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))
	staticObstacles := w.GetComponents(reflect.TypeOf(&components.StaticObstacle{}))

	for e1, pos1 := range positions {
		size1 := sizes[e1].(*components.Size)
		for e2, pos2 := range positions {
			if e1 == e2 {
				continue
			}

			size2 := sizes[e2].(*components.Size)

			// Check if the second entity is a static obstacle
			if _, isObstacle := staticObstacles[e2]; isObstacle {
				// Handle player collisions with static obstacles
				if _, isPlayer := w.GetComponent(e1, reflect.TypeOf(&components.PlayerControlled{})).(*components.PlayerControlled); isPlayer {
					if isColliding(pos1.(*components.Position), size1, pos2.(*components.Position), size2) {
						// Rollback the player's position
						if prevPos, exists := cs.previousPositions[e1]; exists {
							pos1.(*components.Position).X = prevPos.X
							pos1.(*components.Position).Y = prevPos.Y
						}
					}
				}
			}
		}
	}
}

func isColliding(pos1 *components.Position, size1 *components.Size, pos2 *components.Position, size2 *components.Size) bool {
	return pos1.X < pos2.X+size2.Width &&
		pos1.X+size1.Width > pos2.X &&
		pos1.Y < pos2.Y+size2.Height &&
		pos1.Y+size1.Height > pos2.Y
}

func isSweptColliding(
	prevPos1, currPos1 *components.Position, size1 *components.Size,
	pos2 *components.Position, size2 *components.Size,
) bool {
	return isColliding(prevPos1, size1, pos2, size2) || isColliding(currPos1, size1, pos2, size2)
}

package systems

import (
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// MovementSystem updates entities' positions based on their velocity.
type MovementSystem struct{}

func (ms *MovementSystem) Update(w *ecs.World, dt float64) {
	for entity, component := range w.GetComponents(reflect.TypeOf(&components.Velocity{})) {
		velocity := component.(*components.Velocity)
		position := w.GetComponent(entity, reflect.TypeOf(&components.Position{})).(*components.Position)

		position.X += velocity.DX * dt
		position.Y += velocity.DY * dt
	}
}

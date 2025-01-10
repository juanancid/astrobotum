package systems

import (
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
)

// RenderingSystem draws entities based on their Position.
type RenderingSystem struct{}

func (rs *RenderingSystem) Render(w *ecs.World, screen *ebiten.Image) {
	for _, component := range w.GetComponents(reflect.TypeOf(&components.Position{})) {
		position := component.(*components.Position)

		// Draw a simple rectangle for the entity
		ebitenutil.DrawRect(screen, position.X, position.Y, 16, 16, color.White)
	}
}

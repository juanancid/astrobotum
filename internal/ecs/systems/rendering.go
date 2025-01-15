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
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))
	collectibles := w.GetComponents(reflect.TypeOf(&components.Collectible{}))

	for entity, pos := range positions {
		position := pos.(*components.Position)

		// Check if the entity has a Size component
		size, hasSize := sizes[entity]
		width, height := 16.0, 16.0 // Default size if no Size component is present
		if hasSize {
			width = size.(*components.Size).Width
			height = size.(*components.Size).Height
		}

		var itemColor color.Color
		itemColor = color.White

		_, isCollectible := collectibles[entity]
		if isCollectible {
			// itemColor is yellow
			itemColor = color.RGBA{
				R: 234,
				G: 239,
				B: 44,
				A: 0,
			}
		}

		// Render the entity
		ebitenutil.DrawRect(screen, position.X, position.Y, width, height, itemColor)
	}
}

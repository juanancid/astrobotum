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

		if _, isObstacle := w.GetComponent(entity, reflect.TypeOf(&components.StaticObstacle{})).(*components.StaticObstacle); isObstacle {
			itemColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}
		} else if _, isCollectible := w.GetComponent(entity, reflect.TypeOf(&components.Collectible{})).(*components.Collectible); isCollectible {
			itemColor = color.RGBA{R: 234, G: 239, B: 44, A: 0}
		} else if _, isDynamicObstacle := w.GetComponent(entity, reflect.TypeOf(&components.DynamicObstacle{})).(*components.DynamicObstacle); isDynamicObstacle {
			itemColor = color.RGBA{R: 0, G: 255, B: 0, A: 0}
		}

		// Render the entity
		ebitenutil.DrawRect(screen, position.X, position.Y, width, height, itemColor)
	}
}

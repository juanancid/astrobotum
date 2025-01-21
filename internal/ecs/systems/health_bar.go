package systems

import (
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// HealthBarSystem renders a health bar above the player.
type HealthBarSystem struct {
	PlayerEntity ecs.Entity // The player entity
}

func (hbs *HealthBarSystem) Render(w *ecs.World, screen *ebiten.Image) {
	// Get the player's position and health components
	position := w.GetComponent(hbs.PlayerEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	health := w.GetComponent(hbs.PlayerEntity, reflect.TypeOf(&components.Health{})).(*components.Health)

	// Calculate health bar dimensions
	barWidth := 32.0 // Total width of the health bar
	barHeight := 4.0 // Height of the health bar
	healthRatio := float64(health.CurrentHealth) / float64(health.MaxHealth)
	healthWidth := barWidth * healthRatio // Width based on current health

	// Render the health bar background (gray)
	gray := color.RGBA{R: 128, G: 128, B: 128, A: 255}
	ebitenutil.DrawRect(screen, position.X-(barWidth/2), position.Y-8, barWidth, barHeight, gray)

	// Render the health bar foreground (red)
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	ebitenutil.DrawRect(screen, position.X-(barWidth/2), position.Y-8, healthWidth, barHeight, red)
}

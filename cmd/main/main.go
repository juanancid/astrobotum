package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
	"astrobotum/internal/ecs/systems"
)

// Game represents the overall game state.
type Game struct {
	world           *ecs.World
	renderingSystem *systems.RenderingSystem
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0 // Assume a fixed frame rate for simplicity
	g.world.UpdateSystems(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Black background

	// Call the rendering system's Render method
	g.renderingSystem.Render(g.world, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // Game screen resolution
}

func main() {
	// Initialize the ECS world
	world := ecs.NewWorld()

	// Create the player entity and add components
	player := world.AddEntity()
	world.AddComponent(player, &components.Position{X: 160, Y: 120}) // Centered position
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})   // Initial velocity

	// Add systems
	world.AddSystem(&systems.InputSystem{})
	world.AddSystem(&systems.MovementSystem{})
	renderingSystem := &systems.RenderingSystem{}
	// Note: RenderingSystem is not added to the World since it’s invoked separately

	// Start the game
	game := &Game{
		world:           world,
		renderingSystem: renderingSystem,
	}
	ebiten.SetWindowSize(640, 480) // Window size
	ebiten.SetWindowTitle("Astrobotum")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

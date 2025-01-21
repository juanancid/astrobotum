package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
	"github.com/juanancid/astrobotum/internal/ecs/systems"
)

// Game represents the overall game state.
type Game struct {
	world           *ecs.World
	renderingSystem *systems.RenderingSystem
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0 // Assume a fixed frame rate for simplicity

	// Retrieve and save previous positions for collision handling
	if collisionSystem := g.world.GetSystem(&systems.CollisionSystem{}); collisionSystem != nil {
		cs := collisionSystem.(*systems.CollisionSystem) // Type assertion
		cs.SavePreviousPositions(g.world)
	}

	g.world.UpdateSystems(dt)

	// Check if the game is over
	healthSystem := g.world.GetSystem(&systems.HealthSystem{}).(*systems.HealthSystem)
	if healthSystem.GameOver {
		return ebiten.Termination
	}

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
	world.AddComponent(player, &components.Position{X: 160, Y: 120})                   // Centered position
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})                     // Initial velocity
	world.AddComponent(player, &components.Size{Width: 16, Height: 16})                // Entity dimensions
	world.AddComponent(player, &components.PlayerControlled{})                         // Mark as player-controlled
	world.AddComponent(player, &components.Health{CurrentHealth: 100, MaxHealth: 100}) // Health component

	// Add static obstacles
	for i := 0; i < 3; i++ {
		obstacle := world.AddEntity()
		world.AddComponent(obstacle, &components.Position{X: float64(100 + i*50), Y: 150})
		world.AddComponent(obstacle, &components.Size{Width: 32, Height: 32})
		world.AddComponent(obstacle, &components.StaticObstacle{})
	}

	// Create collectible entities
	for i := 0; i < 5; i++ {
		collectible := world.AddEntity()
		world.AddComponent(collectible, &components.Position{X: float64(50 + i*40), Y: 100})
		world.AddComponent(collectible, &components.Size{Width: 16, Height: 16})
		world.AddComponent(collectible, &components.Collectible{Value: 10})
	}

	// Create dynamic obstacles
	for i := 0; i < 3; i++ {
		obstacle := world.AddEntity()
		world.AddComponent(obstacle, &components.Position{X: float64(50 + i*80), Y: 100})
		world.AddComponent(obstacle, &components.Velocity{DX: float64((i + 1) * 20), DY: 0}) // Horizontal movement
		world.AddComponent(obstacle, &components.Size{Width: 16, Height: 16})
		world.AddComponent(obstacle, &components.DynamicObstacle{Damage: 10}) // Inflicts 10 damage on collision
	}

	// Add systems
	world.AddSystem(&systems.InputSystem{})
	world.AddSystem(&systems.MovementSystem{})
	world.AddSystem(&systems.BoundarySystem{ScreenWidth: 320, ScreenHeight: 240}) // Screen dimensions
	dynamicObstacleCollisionSystem := &systems.DynamicObstacleCollisionSystem{PlayerEntity: player}
	world.AddSystem(dynamicObstacleCollisionSystem)
	collectibleSystem := &systems.CollectibleSystem{PlayerEntity: player}
	world.AddSystem(collectibleSystem)
	collisionSystem := systems.NewCollisionSystem()
	world.AddSystem(collisionSystem)
	healthSystem := &systems.HealthSystem{PlayerEntity: player}
	world.AddSystem(healthSystem)

	renderingSystem := &systems.RenderingSystem{}

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

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juanancid/astrobotum/internal/levels"
	"image/color"
	"log"
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
	"github.com/juanancid/astrobotum/internal/ecs/systems"
)

// Game represents the overall game state.
type Game struct {
	world *ecs.World
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

	scoreSystem := g.world.GetSystem(&systems.ScoreSystem{}).(*systems.ScoreSystem)
	if scoreSystem.Victory {
		if g.world.GetCurrentLevel() == 1 {
			levels.LoadLevel2(g.world, scoreSystem.PlayerEntity)
		} else {
			player := scoreSystem.PlayerEntity
			victoryScreen := g.world.GetRenderable(&systems.VictoryScreen{}).(*systems.VictoryScreen)

			victoryScreen.Active = true
			victoryScreen.Score = g.world.GetComponent(player, reflect.TypeOf(&components.Score{})).(*components.Score).Points
		}

		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Black background

	// Render all renderable systems
	g.world.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // Game screen resolution
}

func main() {
	// Initialize the ECS world
	world := ecs.NewWorld()

	// Create the player entity and add components
	player := world.AddEntity()
	world.AddComponent(player, &components.Position{X: 100, Y: 120})                   // Centered position
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})                     // Initial velocity
	world.AddComponent(player, &components.Size{Width: 16, Height: 16})                // Entity dimensions
	world.AddComponent(player, &components.PlayerControlled{})                         // Mark as player-controlled
	world.AddComponent(player, &components.Health{CurrentHealth: 100, MaxHealth: 100}) // Health component
	world.AddComponent(player, &components.Score{Points: 0})                           // Score
	world.AddComponent(player, &components.OnGround{IsGrounded: false})

	// Load level 1
	levels.LoadLevel1(world, player)

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
	scoreSystem := &systems.ScoreSystem{PlayerEntity: player, TargetScore: 100}
	world.AddSystem(scoreSystem)
	gravitySystem := &systems.GravitySystem{Gravity: 100}
	world.AddSystem(gravitySystem)

	renderingSystem := &systems.RenderingSystem{}
	world.AddRenderable(renderingSystem)
	healthBarSystem := &systems.HealthBarSystem{PlayerEntity: player}
	world.AddRenderable(healthBarSystem)
	scoreRenderer := &systems.ScoreRenderer{PlayerEntity: player}
	world.AddRenderable(scoreRenderer)
	victoryScreen := &systems.VictoryScreen{PlayerEntity: player}
	world.AddRenderable(victoryScreen)

	// Start the game
	game := &Game{
		world: world,
	}
	ebiten.SetWindowSize(640, 480) // Window size
	ebiten.SetWindowTitle("Astrobotum")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

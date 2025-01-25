package levels

import (
	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

func LoadLevel1(world *ecs.World, playerEntity ecs.Entity) {
	world.Clear(playerEntity)
	world.SetCurrentLevel(1)

	// Add static obstacles
	for i := 0; i < 3; i++ {
		obstacle := world.AddEntity()
		world.AddComponent(obstacle, &components.Position{X: float64(100 + i*50), Y: 150})
		world.AddComponent(obstacle, &components.Size{Width: 32, Height: 32})
		world.AddComponent(obstacle, &components.StaticObstacle{})
	}

	// Add near-to-ground obstacle
	obstacle := world.AddEntity()
	world.AddComponent(obstacle, &components.Position{X: 50, Y: 200})
	world.AddComponent(obstacle, &components.Size{Width: 64, Height: 16})
	world.AddComponent(obstacle, &components.StaticObstacle{})

	// Add ground
	ground := world.AddEntity()
	world.AddComponent(ground, &components.Position{X: 0, Y: 224})
	world.AddComponent(ground, &components.Size{Width: 320, Height: 16})
	world.AddComponent(ground, &components.StaticObstacle{})

	// Create collectible entities
	for i := 0; i < 5; i++ {
		collectible := world.AddEntity()
		world.AddComponent(collectible, &components.Position{X: float64(50 + i*40), Y: 100})
		world.AddComponent(collectible, &components.Size{Width: 16, Height: 16})
		world.AddComponent(collectible, &components.Collectible{Value: 10})
	}

	// Create healing items
	for i := 0; i < 3; i++ {
		healingItem := world.AddEntity()
		world.AddComponent(healingItem, &components.Position{X: float64(150 + i*50), Y: 200})
		world.AddComponent(healingItem, &components.Size{Width: 16, Height: 16})
		world.AddComponent(healingItem, &components.Collectible{Value: 10})
		world.AddComponent(healingItem, &components.HealingCollectible{HealAmount: 20})
	}

	// Create dynamic obstacles
	for i := 0; i < 3; i++ {
		dynamicObstacle := world.AddEntity()
		world.AddComponent(dynamicObstacle, &components.Position{X: float64(50 + i*80), Y: 100})
		world.AddComponent(dynamicObstacle, &components.Velocity{DX: float64((i + 1) * 20), DY: 0}) // Horizontal movement
		world.AddComponent(dynamicObstacle, &components.Size{Width: 16, Height: 16})
		world.AddComponent(dynamicObstacle, &components.DynamicObstacle{Damage: 10}) // Inflicts 10 damage on collision
		world.AddComponent(dynamicObstacle, &components.OnGround{IsGrounded: false})
	}
}

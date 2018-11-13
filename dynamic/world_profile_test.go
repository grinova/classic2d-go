package dynamic

import (
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/grinova/classic2d-server/physics"
	. "github.com/grinova/classic2d-server/physics/shapes"
	. "github.com/grinova/classic2d-server/vmath"
)

var w = CreateWorld()
var world = resetWorld(&w)

func BenchmarkWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		world.Step(time.Second / 60)
	}
}

var r = rand.New(rand.NewSource(42))

func random(max float64) float64 {
	return r.Float64() * max
}

func resetWorld(world *World) *World {
	world.Clear()
	const ArenaRadius = 3
	const ActorRadius = 0.05
	const ActorsCount = 100
	createArena(world, ArenaRadius)
	createActors(world, ActorsCount, ActorRadius, ArenaRadius)
	return world
}

func createArena(world *World, radius float64) {
	body := world.CreateBody(BodyDef{Inverse: true})
	body.Type = StaticBody
	body.SetFixture(FixtureDef{Shape: CircleShape{Radius: radius}, Density: 1})
}

func createActors(world *World, count int, actorRadius float64, arenaRadius float64) {
	bodies := []*Body{}
	fixtureDef := FixtureDef{Shape: CircleShape{Radius: actorRadius}, Density: 1}
	for i := 0; i < count; i++ {
		position, ok := findEmptyPlace(actorRadius, arenaRadius, bodies, 20)
		if !ok {
			return
		}
		linearVelocity := Vec2{X: random(1)}.Rotate(RotFromAngle(random(2 * math.Pi)))
		bodyDef := BodyDef{Position: position, LinearVelocity: linearVelocity}
		body := world.CreateBody(bodyDef)
		body.SetFixture(fixtureDef)
		bodies = append(bodies, body)
	}
}

func findEmptyPlace(radius float64, arenaRadius float64, bodies []*Body, iterations int) (v Vec2, ok bool) {
	maxEmptyRadius := arenaRadius - 2*radius
	for i := 0; i < iterations; i++ {
		position := Vec2{X: random(maxEmptyRadius), Y: 0}.Rotate(RotFromAngle(random(2 * math.Pi)))
		found := false
		for _, body := range bodies {
			found = body.GetPosition().Sub(position).Length() < math.Pow(radius+body.GetRadius(), 2)
			if found {
				break
			}
		}
		if !found {
			return position, true
		}
	}
	return Vec2{}, false
}

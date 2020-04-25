package main

import (
	"korok.io/korok"
	"korok.io/korok/asset"
	"korok.io/korok/effect"
	"korok.io/korok/game"
	"korok.io/korok/math"
	"korok.io/korok/math/f32"
)

type MainScene struct {
}

func (*MainScene) Load() {
	asset.Texture.Load("particle.png")
}

func (*MainScene) OnEnter(g *game.Game) {
	cfg := &effect.GravityConfig{
		Config: effect.Config{
			Max:      2048,
			Rate:     200,
			Duration: math.MaxFloat32,
			Life:     effect.Var{40.1, 0.4},
			Size:     effect.Range{effect.Var{10, 20}, effect.Var{20, 40}},
			X:        effect.Var{0, 0}, Y: effect.Var{0, 0},
			A: effect.Range{effect.Var{1, 0}, effect.Var{0, 0}},
		},
		Speed:   effect.Var{70, 10},
		Angel:   effect.Var{math.Radian(90), math.Radian(30)},
		Gravity: f32.Vec2{0, -10},
	}
	gravity := korok.Entity.New()
	gParticle := korok.ParticleSystem.NewComp(gravity)
	gParticle.SetSimulator(effect.NewGravitySimulator(cfg))
	gParticle.SetTexture(asset.Texture.Get("particle.png"))
	xf := korok.Transform.NewComp(gravity)
	xf.SetPosition(f32.Vec2{400, 300})
}

func (*MainScene) Update(dt float32) {

}

func (*MainScene) OnExit() {
}

func main() {
	options := &korok.Options{
		Title:  "ParticleSystem",
		Width:  800,
		Height: 600,
	}
	korok.Run(options, &MainScene{})
}

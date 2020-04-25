package main

import (
	"korok.io/korok"
	"korok.io/korok/asset"
	"korok.io/korok/game"
	"korok.io/korok/gfx/dbg"
	"korok.io/korok/math/f32"
)

type MainScene struct {
}

func (m *MainScene) OnEnter(g *game.Game) {

	asset.Texture.Load("avator.jpg")

}

func (m *MainScene) Update(dt float32) {
	dbg.DrawStr(180, 160, "Hello, World")
	face := asset.Texture.Get("avator.jpg")
	entity := korok.Entity.New()
	spr := korok.Sprite.NewComp(entity)
	spr.SetSprite(face)
	xf := korok.Transform.NewComp(entity)
	xf.SetPosition(f32.Vec2{100, 100})
}

func (m *MainScene) OnExit() {

}
func main() {

	options := &korok.Options{
		Title:  "Hello, Korok",
		Width:  480,
		Height: 320,
	}

	korok.Run(options, &MainScene{})

}

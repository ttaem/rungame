package scene

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	//"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/ttaem/rungame/rungame1/animation"
	"github.com/ttaem/rungame/rungame1/global"
)

type GameScene struct {
	runnerImg *ebiten.Image
	backImg   *ebiten.Image
	animation *animation.Handler
}

func (g *GameScene) StartUp() {
	var err error

	runnerImg, _, err := ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
	g.backImg, _, err = ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}

	g.animation = animation.New()

	sprites := make([]*ebiten.Image, global.RunningFrames)
	for i := 0; i < global.RunningFrames; i++ {
		sx := 0 + i*global.FrameWidth
		sy := global.FrameHeight
		sprites[i] = runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	}
	speed := global.RunningAnimSpeed

	g.animation.Add("Run", sprites, speed)
	g.animation.Play("Run")
}

func (g *GameScene) Update(screen *ebiten.Image) error {
	frameCount++

	bgWidth, _ := g.backImg.Size()

	opt := &ebiten.DrawImageOptions{}
	backX := (frameCount / 2) % bgWidth

	opt.GeoM.Translate(-float64(backX), 0)
	screen.DrawImage(g.backImg, opt)

	opt.GeoM.Translate(float64(bgWidth), 0)
	screen.DrawImage(g.backImg, opt)

	/* Draw Running Animation */
	g.animation.Update(screen, 0, float64(global.ScreenHeight/2))

	return nil
}

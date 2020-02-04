package scene

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	//"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/ttaem/rungame/rungame1/global"
)

type GameScene struct {
	runnerImg *ebiten.Image
	backImg   *ebiten.Image
}

func (g *GameScene) StartUp() {
	var err error

	frameCount = 0
	g.runnerImg, _, err = ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
	g.backImg, _, err = ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
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
	frameIdx := (frameCount / global.RunningAnimSpeed) % global.RunningFrames

	sx := frameIdx * global.FrameWidth
	sy := global.FrameHeight

	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, float64(global.ScreenHeight/2))

	screen.DrawImage(g.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image), opt)
	return nil
}

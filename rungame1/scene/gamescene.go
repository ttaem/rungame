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
}

func (g *GameScene) StartUp() {
	var err error
	g.runnerImg, _, err = ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
}

func (g *GameScene) Update(screen *ebiten.Image) error {
	frameCount++
	/* Draw Running Animation */
	frameIdx := (frameCount / global.RunningAnimSpeed) % global.RunningFrames

	sx := frameIdx * global.FrameWidth
	sy := global.FrameHeight

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(global.ScreenWidth/2), float64(global.ScreenHeight/2))
	opt.GeoM.Translate(-float64(global.FrameWidth/2), -float64(global.FrameHeight/2))

	screen.DrawImage(g.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image), opt)
	return nil
}

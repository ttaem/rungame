package scene

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/ttaem/rungame/rungame1/global"
)

type StartScene struct {
	runnerImg *ebiten.Image
}

func (s *StartScene) StartUp() {
	var err error
	s.runnerImg, _, err = ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
}

var frameCount = 0

func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	frameIdx := (frameCount / 5) % global.RunningFrames

	sx := frameIdx * global.FrameWidth
	sy := global.FrameHeight

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(global.ScreenWidth/2), float64(global.ScreenHeight/2))
	opt.GeoM.Translate(-float64(global.FrameWidth/2), -float64(global.FrameHeight/2))

	screen.DrawImage(s.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image), opt)

	return nil
}

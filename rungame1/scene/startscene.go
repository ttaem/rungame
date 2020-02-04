package scene

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/ttaem/rungame/rungame1/font"
	"github.com/ttaem/rungame/rungame1/global"
	"github.com/ttaem/rungame/rungame1/scenemanager"
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

	/* Draw Idle Animation */
	frameIdx := (frameCount / global.IdleAnimSpeed) % global.IdleFrames
	sx := global.IdleX + frameIdx*global.FrameWidth
	sy := global.IdleY

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, float64(global.ScreenHeight/2))
	opt.GeoM.Translate(0, -float64(global.FrameHeight/2))

	screen.DrawImage(s.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image), opt)

	/* Draw Text */
	size := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText,
		global.ScreenWidth/2-size/2, global.ScreenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})

	}

	return nil
}

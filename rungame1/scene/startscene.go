package scene

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/ttaem/rungame/rungame1/animation"
	"github.com/ttaem/rungame/rungame1/font"
	"github.com/ttaem/rungame/rungame1/global"
	"github.com/ttaem/rungame/rungame1/scenemanager"
)

type StartScene struct {
	runnerImg *ebiten.Image
	backImg   *ebiten.Image
	animation *animation.Handler
}

func (s *StartScene) StartUp() {
	runnerImg, _, err := ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
	s.backImg, _, err = ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}

	s.animation = animation.New()

	sprites := make([]*ebiten.Image, global.IdleFrames)
	for i := 0; i < global.IdleFrames; i++ {
		sx := global.IdleX + i*global.FrameWidth
		sy := global.IdleY
		sprites[i] = runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	}
	speed := global.IdleAnimSpeed

	s.animation.Add("Idle", sprites, speed)
	s.animation.Play("Idle")
}

var frameCount = 0

func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	screen.DrawImage(s.backImg, nil)

	/* Draw Idle Animation */
	s.animation.Update(screen, 0, float64(global.ScreenHeight/2))

	/* Draw Text */
	size := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText,
		global.ScreenWidth/2-size/2, global.ScreenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})

	}

	return nil
}

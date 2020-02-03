// ex1
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ttaem/rungame/rungame1/global"
	"github.com/ttaem/rungame/rungame1/scene"
	"github.com/ttaem/rungame/rungame1/scenemanager"
)

func update(screen *ebiten.Image) error {
	/* Input Process */
	return nil
}

func main() {
	var err error

	scenemanager.SetScene(&scene.StartScene{})

	err = ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 2, "Run Game")
	if err != nil {
		log.Fatalf("Run Error: %v", err)
	}
}

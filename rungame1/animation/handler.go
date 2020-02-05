package animation

import (
	"github.com/hajimehoshi/ebiten"
)

type animInfo struct {
	sprites []*ebiten.Image
	speed   int
}

type Handler struct {
	animMap     map[string]animInfo
	currentAnim animInfo
	lastIdx     int
	remainFrame int
}

func New() *Handler {
	h := &Handler{}
	h.animMap = make(map[string]animInfo)
	return h
}

func (h *Handler) Add(name string, sprites []*ebiten.Image, speed int) {
	h.animMap[name] = animInfo{sprites, speed}
}

func (h *Handler) Play(name string) {
	h.currentAnim = h.animMap[name]
	h.lastIdx = 0
	h.remainFrame = h.currentAnim.speed
}

func (h *Handler) Update(screen *ebiten.Image, x, y float64) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(x, y)

	screen.DrawImage(h.currentAnim.sprites[h.lastIdx], opt)

	h.remainFrame--
	if h.remainFrame == 0 {
		h.lastIdx++
		if len(h.currentAnim.sprites) == h.lastIdx {
			h.lastIdx = 0
		}
		h.remainFrame = h.currentAnim.speed
	}

}

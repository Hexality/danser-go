package render

import (
	"github.com/wieku/glhf"
	"github.com/wieku/danser/utils"
)

var Circle *glhf.Texture
var ApproachCircle *glhf.Texture
var CircleFull *glhf.Texture
var CircleOverlay *glhf.Texture
var SliderGradient *glhf.Texture
var SliderTick *glhf.Texture
var SliderBall *glhf.Texture
var CursorTex *glhf.Texture
var CursorTop *glhf.Texture
var CursorTrail *glhf.Texture
var SpinnerMiddle *glhf.Texture
var SpinnerMiddle2 *glhf.Texture
var SpinnerAC *glhf.Texture

func LoadTextures() {
	Circle, _ = utils.LoadTexture("assets/textures/hitcircle.png")
	ApproachCircle, _ = utils.LoadTexture("assets/textures/approachcircle.png")
	CircleFull, _ = utils.LoadTexture("assets/textures/hitcircle-full.png")
	CircleOverlay, _ = utils.LoadTexture("assets/textures/hitcircleoverlay.png")
	SliderGradient, _ = utils.LoadTexture("assets/textures/slidergradient.png")
	SliderTick, _ = utils.LoadTexture("assets/textures/sliderscorepoint.png")
	SliderBall, _ = utils.LoadTexture("assets/textures/sliderball.png")
	CursorTex, _ = utils.LoadTexture("assets/textures/cursor.png")
	CursorTop, _ = utils.LoadTexture("assets/textures/cursor-top.png")
	CursorTrail, _ = utils.LoadTexture("assets/textures/cursortrail.png")
	SpinnerMiddle, _ = utils.LoadTextureU("assets/textures/spinner-middle.png")
	SpinnerMiddle2, _ = utils.LoadTextureU("assets/textures/spinner-middle2.png")
	SpinnerAC, _ = utils.LoadTextureU("assets/textures/spinner-approachcircle.png")
}

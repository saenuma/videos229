package sprites

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/saenuma/videos229/v2shared"
)

func GetRootPath() (string, error) {
	return v2shared.GetRootPath()
}

func timeFormatToSeconds(s string) int {
	return v2shared.TimeFormatToSeconds(s)
}

func pasteWithoutTransparentBackground(backgroundImg *image.NRGBA, spriteImg image.Image, xOrigin, yOrigin int) *image.NRGBA {
	newRectangle := image.Rect(xOrigin, yOrigin, xOrigin+spriteImg.Bounds().Dx(), yOrigin+spriteImg.Bounds().Dy())
	draw.DrawMask(backgroundImg, newRectangle, spriteImg, image.Pt(0, 0),
		image.NewUniform(color.RGBA{255, 255, 255, 255}), image.Pt(0, 0),
		draw.Over)

	return backgroundImg
}

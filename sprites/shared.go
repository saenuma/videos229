package sprites

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/kovidgoyal/imaging"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/saenuma/videos229/v2shared"
)

func GetRootPath() (string, error) {
	return v2shared.GetRootPath()
}

func pasteWithoutTransparentBackground(backgroundImg *image.NRGBA, spriteImg image.Image, xOrigin, yOrigin int) *image.NRGBA {
	newRectangle := image.Rect(xOrigin, yOrigin, xOrigin+spriteImg.Bounds().Dx(), yOrigin+spriteImg.Bounds().Dy())
	draw.DrawMask(backgroundImg, newRectangle, spriteImg, image.Pt(0, 0),
		image.NewUniform(color.RGBA{255, 255, 255, 255}), image.Pt(0, 0),
		draw.Over)

	return backgroundImg
}

func scaleSprite(img image.Image, scale float64) image.Image {
	newWidth, newHeight := float64(img.Bounds().Dx())*scale, float64(img.Bounds().Dy())*scale
	return imaging.Resize(img, int(newWidth), int(newHeight), imaging.Lanczos)
}

func addGutter(img image.Image, gutter int, bgColor colorful.Color) image.Image {
	newWidth, newHeight := img.Bounds().Dx()+(2*gutter), img.Bounds().Dy()+(2*gutter)
	newImg := imaging.New(newWidth, newHeight, bgColor)
	return imaging.PasteCenter(newImg, img)
}

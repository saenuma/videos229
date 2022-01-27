package main

import (
  "image/draw"
  "image/color"
  "github.com/disintegration/imaging"
  "github.com/lucasb-eyer/go-colorful"
  "image"
  // "fmt"
  "math"
)


func main() {

  // paste test
  newColor, err := colorful.Hex("#DFBABA")
  if err != nil {
    panic(err)
  }
  backgroundImg := imaging.New(1366, 768, newColor)

  spriteImg, err := imaging.Open("/home/bankole/videos229_data/v229_sprite_t3.png")
  if err != nil {
    panic(err)
  }

  numberOfXIterations := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx() )
  numberOfYIternations := int(backgroundImg.Bounds().Dy() / spriteImg.Bounds().Dy())

  newBackgroundImg := imaging.New(1366, 768, color.White)
  newBackgroundImg = imaging.Paste(newBackgroundImg, backgroundImg, image.Pt(0, 0))

  for x := 0; x < numberOfXIterations + 1; x++ {
    for y := 0; y < numberOfYIternations + 1; y++ {
      newX := x * spriteImg.Bounds().Dx()
      newY := y * spriteImg.Bounds().Dy()

      if int(math.Mod(float64(x), float64(2))) == 0 {
        newBackgroundImg = pasteWithoutTransparentBackground2(newBackgroundImg, spriteImg, newX, newY, 255)
      } else {
        newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, spriteImg, newX, newY)
      }
    }
  }
  imaging.Save(newBackgroundImg, "out_t3_1.png")
}


func pasteWithoutTransparentBackground(backgroundImg *image.NRGBA, spriteImg image.Image, xOrigin, yOrigin int) *image.NRGBA {

  newRectangle := image.Rect(xOrigin, yOrigin, xOrigin + spriteImg.Bounds().Dx(), yOrigin + spriteImg.Bounds().Dy())
  draw.DrawMask(backgroundImg, newRectangle, spriteImg, image.Pt(0,0),
    image.NewUniform(color.RGBA{255, 255, 255, 255 }), image.Pt(0,0),
    draw.Over)

  return backgroundImg

}


func pasteWithoutTransparentBackground2(backgroundImg *image.NRGBA, spriteImg image.Image, xOrigin, yOrigin int, transparency uint8) *image.NRGBA {

  newRectangle := image.Rect(xOrigin, yOrigin, xOrigin + spriteImg.Bounds().Dx(), yOrigin + spriteImg.Bounds().Dy())
  draw.DrawMask(backgroundImg, newRectangle, spriteImg, image.Pt(0,0),
    image.NewUniform(color.RGBA{255, 255, 255, uint8(transparency) }), image.Pt(0,0),
    draw.Over)

  return backgroundImg
}

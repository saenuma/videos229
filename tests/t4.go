package main

import (
  // "image/draw"
  "image/color"
  "github.com/disintegration/imaging"
  "github.com/lucasb-eyer/go-colorful"
  "image"
  // "fmt"
  // "math"
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

      newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, spriteImg, newX, newY)
    }
  }
  imaging.Save(newBackgroundImg, "out_t3_1.png")
}


func pasteWithoutTransparentBackground(backgroundImg *image.NRGBA, spriteImg image.Image, xOrigin, yOrigin int) *image.NRGBA {

  for y := 0; y < spriteImg.Bounds().Max.Y; y++ {
    for x := 0; x < spriteImg.Bounds().Max.X; x++ {
      tmpColor := spriteImg.At(x, y)
      r, g, b, a := tmpColor.RGBA()
      if r == 0 && g == 0 && b == 0 && a == 0 {
        continue
      } else {
        backgroundImg.Set(xOrigin + x, yOrigin + y, tmpColor)
      }
    }
  }

  return backgroundImg
}

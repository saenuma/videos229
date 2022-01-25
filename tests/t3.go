package main

import (
  // "image/draw"
  // "image/color"
  "github.com/disintegration/imaging"
  "github.com/lucasb-eyer/go-colorful"
  // "image"
  // "fmt"
  // "math"
)


func main() {

  newColor, err := colorful.Hex("#DFBABA")
  if err != nil {
    panic(err)
  }
  backgroundImg := imaging.New(1366, 768, newColor)

  spriteImg, err := imaging.Open("/home/bankole/videos229_data/v229_sprite_t3.png")
  if err != nil {
    panic(err)
  }

  xOrigin := 600;
  yOrigin := 200;

  for x := 0; x < spriteImg.Bounds().Max.X; x++ {
    for y := 0; y < spriteImg.Bounds().Max.Y; y++ {
      tmpColor := spriteImg.At(x, y)
      r, g, b, a := tmpColor.RGBA()
      if r == 0 && g == 0 && b == 0 && a == 0 {
        continue
      } else if r == 0 && g == 0 && b == 0 {
        continue
      } else {
        backgroundImg.Set(xOrigin + x, yOrigin + y, tmpColor)
      }
    }
  }

  imaging.Save(backgroundImg, "out_t3.png")
}

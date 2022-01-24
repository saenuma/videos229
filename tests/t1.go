package main

import (
  "github.com/disintegration/imaging"
  "image"
  "image/color"
  "fmt"
)


func main() {
  img1, err := imaging.Open("/home/bankole/videos229_data/v229_sprite.png")
  if err != nil {
    panic(err)
  }

  // begin the joining proper
  joinedImg := image.NewNRGBA(image.Rect(0,0,1366,768))

  for x := 0; x < joinedImg.Bounds().Dx(); x++ {
    for y := 0; y < joinedImg.Bounds().Dy(); y++ {
      joinedImg.Set(x, y, color.White)
    }
  }

  rotatedImg1 := imaging.Rotate(img1, 0.0, color.White)

  joinedImg = imaging.Paste(joinedImg, rotatedImg1, image.Pt(50,50))

  err = imaging.Save(joinedImg, "out.png")
  if err != nil {
    panic(err)
  }

  fmt.Println("all done.")
}

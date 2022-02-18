package main

import (
  "image"
  "image/color"
  "math"
  "fmt"
  "github.com/disintegration/imaging"
)

func main() {
  anImage, err := imaging.Open("t.jpg")
  if err != nil {
    panic(err)
  }

  newBackgroundImg := imaging.New(1366, 768, color.White)

  var anImageResized image.Image
  if anImage.Bounds().Dx() > newBackgroundImg.Bounds().Dx() {
    newH := (float64(anImage.Bounds().Dx()) * float64(768)) / float64(1366)
    anImageResized = imaging.Fit(anImage, 1366, int(math.Ceil(newH)), imaging.NearestNeighbor)
  } else if anImage.Bounds().Dy() > newBackgroundImg.Bounds().Dy() {
    newW := (float64(1366) * float64(768)) / float64(anImage.Bounds().Dy())
    anImageResized = imaging.Fit(anImage, int(math.Ceil(newW)), 768, imaging.NearestNeighbor)
  } else {
    anImageResized = anImage
  }

  toWriteImage := imaging.PasteCenter(newBackgroundImg, anImageResized)

  imaging.Save(toWriteImage, "out.jpg")
}

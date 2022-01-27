package main

import (
  "os"
  // "fmt"
  "time"
  "path/filepath"
  "github.com/bankole7782/zazabul"
  "github.com/disintegration/imaging"
  "math/rand"
  "image"
  "image/color"
  "strconv"
  "math"
  "github.com/lucasb-eyer/go-colorful"
)


// method1 generates a video with the sprite dancing round a circle
func method1(conf zazabul.Config) string {
  rootPath, _ := GetRootPath()

  outName := "s" + time.Now().Format("20060102T150405")
  renderPath := filepath.Join(rootPath, outName)
  os.MkdirAll(renderPath, 0777)

  spriteImg, err := imaging.Open(filepath.Join(rootPath, conf.Get("sprite_file")))
  if err != nil {
    panic(err)
  }

  backgroundColor, err := colorful.Hex(conf.Get("background_color"))
  if err != nil {
    panic(err)
  }
  backgroundImg := imaging.New(1366, 768, backgroundColor)

  var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

  radius := 200
  xOrigin := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
  yOrigin := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
  xOrigin2 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
  yOrigin2 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
  xOrigin3 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
  yOrigin3 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
  xOrigin4 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
  yOrigin4 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
  xOrigin5 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
  yOrigin5 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
  xOrigin6 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
  yOrigin6 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
  var tinyAngle float64
  var angleIncrement float64 = float64(0.5)

  for seconds := 0; seconds < 60; seconds++ {

    for i := 1; i <= 60; i++ {
      out := (60 * seconds) + i
      outPath := filepath.Join(renderPath, strconv.Itoa(out) + ".png")

      tinyAngle += angleIncrement

      toWriteImage := writeRotation(backgroundImg, spriteImg, xOrigin, yOrigin, radius, tinyAngle, 1)
      toWriteImage = writeRotation(toWriteImage, spriteImg, xOrigin2, yOrigin2, radius, tinyAngle, 2)
      toWriteImage = writeRotation(toWriteImage, spriteImg, xOrigin3, yOrigin3, radius, tinyAngle, 1)
      toWriteImage = writeRotation(toWriteImage, spriteImg, xOrigin4, yOrigin4, radius, tinyAngle, 2)
      toWriteImage = writeRotation(toWriteImage, spriteImg, xOrigin5, yOrigin5, radius, tinyAngle, 1)
      toWriteImage = writeRotation(toWriteImage, spriteImg, xOrigin6, yOrigin6, radius, tinyAngle, 2)
      imaging.Save(toWriteImage, outPath)
    }


  }

  return outName
}


func writeRotation(background, sprite image.Image, xOrigin, yOrigin, radius int, angle float64, rotationStyle int) image.Image {
  angleInRadians := angle * (math.Pi / 180)
  var xCircle float64
  var yCircle float64
  if rotationStyle == 1 {
    xCircle = float64(radius) * math.Sin(-angleInRadians)
    yCircle = float64(radius) * math.Cos(-angleInRadians)
  } else {
    xCircle = float64(radius) * math.Sin(angleInRadians)
    yCircle = float64(radius) * math.Cos(angleInRadians)
  }

  newBackgroundImg := imaging.New(1366, 768, color.White)
  newBackgroundImg = imaging.Paste(newBackgroundImg, background, image.Pt(0, 0))

  return pasteWithoutTransparentBackground(newBackgroundImg, sprite, xOrigin + int(xCircle), yOrigin + int(yCircle))
}

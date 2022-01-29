package main

import (
  "os"
  // "fmt"
  "time"
  "path/filepath"
  "github.com/bankole7782/zazabul"
  "github.com/disintegration/imaging"
  // "math/rand"
  "image"
  "image/color"
  "image/draw"
  "strconv"
  "math"
  "github.com/lucasb-eyer/go-colorful"
)


// method1 generates a video with the sprite dancing round a circle
func method2(conf zazabul.Config) string {
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

  // var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))


  var increment uint8 = 2
  var transparency  uint8 = 255
  totalSeconds := timeFormatToSeconds(conf.Get("video_length"))
  for seconds := 0; seconds < totalSeconds; seconds++ {

    for i := 1; i <= 60; i++ {
      out := (60 * seconds) + i
      outPath := filepath.Join(renderPath, strconv.Itoa(out) + ".png")

      transparency -= increment
      if transparency <= 0 {
        transparency = 255
      }
      toWriteImage := makePattern(backgroundImg, spriteImg, transparency)
      imaging.Save(toWriteImage, outPath)
    }


  }

  return outName
}

func makePattern(backgroundImg, spriteImg image.Image, transparency uint8) *image.NRGBA {
  numberOfXIterations := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx() )
  numberOfYIternations := int(backgroundImg.Bounds().Dy() / spriteImg.Bounds().Dy())

  newBackgroundImg := imaging.New(1366, 768, color.White)
  newBackgroundImg = imaging.Paste(newBackgroundImg, backgroundImg, image.Pt(0, 0))

  for x := 0; x < numberOfXIterations + 1; x++ {
    for y := 0; y < numberOfYIternations + 1; y++ {
      newX := x * spriteImg.Bounds().Dx()
      newY := y * spriteImg.Bounds().Dy()

      if int(math.Mod(float64(x), float64(2))) == 0 {
        newBackgroundImg = pasteWithoutTransparentBackground2(newBackgroundImg, spriteImg, newX, newY, transparency)
      } else {
        newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, spriteImg, newX, newY)
      }
    }
  }

  return newBackgroundImg
}



func pasteWithoutTransparentBackground2(backgroundImg *image.NRGBA, spriteImg image.Image, xOrigin, yOrigin int, transparency uint8) *image.NRGBA {

  newRectangle := image.Rect(xOrigin, yOrigin, xOrigin + spriteImg.Bounds().Dx(), yOrigin + spriteImg.Bounds().Dy())
  draw.DrawMask(backgroundImg, newRectangle, spriteImg, image.Pt(0,0),
    image.NewUniform(color.RGBA{255, 255, 255, uint8(transparency) }), image.Pt(0,0),
    draw.Over)

  return backgroundImg
}

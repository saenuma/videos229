package main

import (
  "os"
  // "fmt"
  color2 "github.com/gookit/color"
  "time"
  "path/filepath"
  "github.com/bankole7782/zazabul"
  "github.com/disintegration/imaging"
  "math/rand"
  "image"
  "strconv"
  "math"
)


// method1 generates a video with the sprite dancing round a circle
func method1(args []string) string {
  rootPath, _ := GetRootPath()

  if len(args) != 3 {
    color2.Red.Println("The run command expects a file created by the init1 command")
    os.Exit(1)
  }

  confPath := filepath.Join(rootPath, args[2])

  conf, err := zazabul.LoadConfigFile(confPath)
  if err != nil {
    panic(err)
    os.Exit(1)
  }

  for _, item := range conf.Items {
    if item.Value == "" {
      color2.Red.Println("Every field in the launch file is compulsory.")
      os.Exit(1)
    }
  }


  outName := "s" + time.Now().Format("20060102T150405")
  renderPath := filepath.Join(rootPath, outName)
  os.MkdirAll(renderPath, 0777)

  spriteImg, err := imaging.Open(filepath.Join(rootPath, conf.Get("sprite_file")))
  if err != nil {
    panic(err)
  }

  backgroundImg := image.NewNRGBA(image.Rect(0,0,1366,768))
  backgroundColor := hexToNRGBA(conf.Get("background_color"))

  for x := 0; x < backgroundImg.Bounds().Dx(); x++ {
    for y := 0; y < backgroundImg.Bounds().Dy(); y++ {
      backgroundImg.Set(x, y, backgroundColor)
    }
  }

  rand.Seed(time.Now().UnixNano())

  xOrigin := (backgroundImg.Bounds().Dx() / 2 ) - (spriteImg.Bounds().Dx() / 2)
  yOrigin := (backgroundImg.Bounds().Dy() / 2 ) - (spriteImg.Bounds().Dy() / 2)

  radius := 100

  var tinyAngle float64
  var angleIncrement float64 = float64(0.5)

  for seconds := 0; seconds < 60; seconds++ {

    for i := 1; i <= 60; i++ {
      out := (60 * seconds) + i
      outPath := filepath.Join(renderPath, strconv.Itoa(out) + ".png")

      tinyAngle += angleIncrement
      toWriteImage := writeRotation(backgroundImg, spriteImg, 0, yOrigin, radius, tinyAngle, "pos")
      toWriteImage = writeRotation(toWriteImage, spriteImg, xOrigin, yOrigin, radius, tinyAngle, "neg")
      thirdXOrigin := backgroundImg.Bounds().Dx() - (spriteImg.Bounds().Dx() / 2)
      toWriteImage = writeRotation(toWriteImage, spriteImg, thirdXOrigin, yOrigin, radius, tinyAngle, "pos")
      imaging.Save(toWriteImage, outPath)
    }


  }

  return outName
}


func writeRotation(background, sprite image.Image, xOrigin, yOrigin, radius int, angle float64, rotationStyle string) image.Image {
  angleInRadians := angle * (math.Pi / 180)
  var x float64
  var y float64
  if rotationStyle == "pos" {
    x = float64(radius) * math.Sin(angleInRadians)
    y = float64(radius) * math.Cos(angleInRadians)
  } else {
    x = float64(radius) * math.Sin(-angleInRadians)
    y = float64(radius) * math.Cos(-angleInRadians)
  }

  return imaging.Paste(background, sprite, image.Pt(xOrigin + int(x), yOrigin + int(y)))
}

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
    color2.Red.Println("The run command expects a file created by the init command")
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

  backgroundColor := hexToNRGBA(conf.Get("background_color"))
  backgroundImg := imaging.New(1366, 768, backgroundColor)

  rand.Seed(time.Now().UnixNano())
  radius := 200
  xOrigin := rand.Intn(backgroundImg.Bounds().Dx() / 3)
  yOrigin := rand.Intn(backgroundImg.Bounds().Dy() / 3)
  xOrigin2 := rand.Intn(backgroundImg.Bounds().Dx() * 2 / 3)
  yOrigin2 := rand.Intn(backgroundImg.Bounds().Dy() * 2 / 3)
  xOrigin3 := rand.Intn(backgroundImg.Bounds().Dx())
  yOrigin3 := rand.Intn(backgroundImg.Bounds().Dy())
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
      imaging.Save(toWriteImage, outPath)
    }


  }

  return outName
}


func writeRotation(background, sprite image.Image, xOrigin, yOrigin, radius int, angle float64, rotationStyle int) image.Image {
  angleInRadians := angle * (math.Pi / 180)
  var x float64
  var y float64
  if rotationStyle == 1 {
    x = float64(radius) * math.Sin(-angleInRadians)
    y = float64(radius) * math.Cos(-angleInRadians)
  } else {
    x = float64(radius) * math.Sin(angleInRadians)
    y = float64(radius) * math.Cos(angleInRadians)
  }

  return imaging.Paste(background, sprite, image.Pt(xOrigin + int(x), yOrigin + int(y)))
}

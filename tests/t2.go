package main

import (
  "image/draw"
  "image/color"
  "github.com/disintegration/imaging"
  "image"
  "fmt"
  "math"
)


func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
    x, y, dx, dy := r-1, 0, 1, 1
    err := dx - (r * 2)

    for x > y {
        img.Set(x0+x, y0+y, c)
        img.Set(x0+y, y0+x, c)
        img.Set(x0-y, y0+x, c)
        img.Set(x0-x, y0+y, c)
        img.Set(x0-x, y0-y, c)
        img.Set(x0-y, y0-x, c)
        img.Set(x0+y, y0-x, c)
        img.Set(x0+x, y0-y, c)

        if err <= 0 {
            y++
            err += dy
            dy += 2
        }
        if err > 0 {
            x--
            dx += 2
            err += dx - (r * 2)
        }
    }
}


func main() {
  // begin the joining proper
  joinedImg := image.NewNRGBA(image.Rect(0,0,1366,768))

  for x := 0; x < joinedImg.Bounds().Dx(); x++ {
    for y := 0; y < joinedImg.Bounds().Dy(); y++ {
      joinedImg.Set(x, y, color.White)
    }
  }

  pointWiseDrawCurve(joinedImg, 200, color.Black)
  err := imaging.Save(joinedImg, "out_t2.png")
  if err != nil {
    panic(err)
  }

  fmt.Println("all done.")
}


func pointWiseDrawCurve(img draw.Image, r int, c color.Color) {

  for angle := 0; angle < 360; angle++ {
    angleInRadians := float64(angle) * (math.Pi / 180)
    x := float64(r) * math.Sin(angleInRadians)
    y := float64(r) * math.Cos(angleInRadians)

    img.Set(1366/2 + int(x), 768/2 + int(y), color.Black)
  }
}

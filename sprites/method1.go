package sprites

import (
	"os"
	// "fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	color2 "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/saenuma/zazabul"
)

// method1 generates a video with the sprite dancing round a circle
func Method1(conf zazabul.Config) string {
	rootPath, _ := GetRootPath()

	outName := "sp_" + time.Now().Format("20060102T150405")
	renderPath := filepath.Join(rootPath, outName)
	os.MkdirAll(renderPath, 0777)

	spriteImg, err := imaging.Open(filepath.Join(rootPath, conf.Get("sprite_file")))
	if err != nil {
		color2.Red.Printf("The sprite file '%s' does not exist.\n Exiting.\n", filepath.Join(rootPath, conf.Get("sprite_file")))
		os.Exit(1)
	}

	backgroundColor, err := colorful.Hex(conf.Get("background_color"))
	if err != nil {
		color2.Red.Printf("The color code '%s' is not valid.\nExiting.\n", conf.Get("background_color"))
		os.Exit(1)
	}
	backgroundImg := imaging.New(1366, 768, backgroundColor)

	totalSeconds := timeFormatToSeconds(conf.Get("video_length"))

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	radius := 200
	xOrigin := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
	yOrigin := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
	xOrigin2 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
	yOrigin2 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
	xOrigin3 := seededRand.Intn(backgroundImg.Bounds().Dx()) - (spriteImg.Bounds().Dx() / 2)
	yOrigin3 := seededRand.Intn(backgroundImg.Bounds().Dy()) - (spriteImg.Bounds().Dy() / 2)
	var angleIncrement float64 = float64(2)

	var tinyAngle float64

	for seconds := 0; seconds < totalSeconds; seconds++ {
		for i := 1; i <= 60; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

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

	return pasteWithoutTransparentBackground(newBackgroundImg, sprite, xOrigin+int(xCircle), yOrigin+int(yCircle))
}

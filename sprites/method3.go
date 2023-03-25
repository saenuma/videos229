package sprites

import (
	"os"
	// "fmt"
	"image"
	"image/color"
	"math"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	color2 "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/saenuma/zazabul"
)

// method3 for rotation in place style
func Method3(conf zazabul.Config) string {
	rootPath, _ := GetRootPath()

	outName := ".tmp_sp_" + time.Now().Format("20060102T150405")
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

	videoWidth, _ := strconv.Atoi(conf.Get("video_width"))
	videoHeight, _ := strconv.Atoi(conf.Get("video_height"))

	backgroundImg := imaging.New(videoWidth, videoHeight, backgroundColor)

	// var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	increment := 5.0
	totalSeconds, _ := strconv.Atoi(conf.Get("video_length"))

	var rotationAngle float64
	for seconds := 0; seconds < totalSeconds; seconds++ {

		for i := 1; i <= 60; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			rotationAngle += increment

			toWriteImage := makePatternWithRotations(backgroundImg, spriteImg, rotationAngle, backgroundColor)
			imaging.Save(toWriteImage, outPath)
		}

	}

	return outName
}

func makePatternWithRotations(backgroundImg, spriteImg image.Image, rotationAngle float64, bgColor color.Color) *image.NRGBA {
	numberOfXIterations := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx())
	numberOfYIternations := int(backgroundImg.Bounds().Dy() / spriteImg.Bounds().Dy())

	newBackgroundImg := imaging.New(backgroundImg.Bounds().Dx(), backgroundImg.Bounds().Dy(), color.White)
	newBackgroundImg = imaging.Paste(newBackgroundImg, backgroundImg, image.Pt(0, 0))

	for x := 0; x < numberOfXIterations+1; x++ {
		for y := 0; y < numberOfYIternations+1; y++ {
			newX := x * spriteImg.Bounds().Dx()
			newY := y * spriteImg.Bounds().Dy()

			if int(math.Mod(float64(x), float64(2))) == 0 {
				rotatedSpriteImage := imaging.Rotate(spriteImg, rotationAngle, bgColor)
				newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, rotatedSpriteImage, newX, newY)
			} else {
				newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, spriteImg, newX, newY)
			}
		}
	}

	return newBackgroundImg
}

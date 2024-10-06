package sprites

import (
	"image"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"

	color2 "github.com/gookit/color"
	"github.com/kovidgoyal/imaging"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/saenuma/zazabul"
)

var Method2Conf = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768


// video_length is the length of the output video in seconds
video_length: 10

// scale. a float to resize the sprite
scale: 1.0

// gutter. a int to add padding to the sprite
gutter: 20

// increment is the increment to add to each rotation in this movement
increment: 10
	`

// method3 for rotation in place style
func Method2(conf zazabul.Config) string {
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

	videoWidth, _ := strconv.Atoi(conf.Get("video_width"))
	videoHeight, _ := strconv.Atoi(conf.Get("video_height"))

	gutter, _ := strconv.Atoi(conf.Get("gutter"))
	scale, err := strconv.ParseFloat(conf.Get("scale"), 64)
	if err != nil {
		color2.Red.Println("Invalid scale value. Expecting a float")
		os.Exit(1)
	}
	spriteImg = scaleSprite(spriteImg, scale)
	spriteImg = addGutter(spriteImg, gutter, backgroundColor)

	backgroundImg := imaging.New(videoWidth, videoHeight, backgroundColor)

	increment, _ := strconv.Atoi(conf.Get("increment"))
	totalSeconds, _ := strconv.Atoi(conf.Get("video_length"))

	var rotationAngle float64
	for seconds := 0; seconds < totalSeconds; seconds++ {

		for i := 1; i <= 24; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			rotationAngle += float64(increment)

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

	for x := 0; x <= numberOfXIterations; x++ {
		for y := 0; y <= numberOfYIternations; y++ {
			if int(math.Mod(float64(x), 2)) == 0 && int(math.Mod(float64(y), 2)) == 0 {
				newX := (x * spriteImg.Bounds().Dx())
				newY := (y * spriteImg.Bounds().Dy())

				rotatedSpriteImage := imaging.Rotate(spriteImg, rotationAngle, bgColor)
				newWidth, newHeight := float64(spriteImg.Bounds().Dx())*1.5, float64(spriteImg.Bounds().Dy())*1.5
				newSpriteImg := imaging.New(int(newWidth), int(newHeight), color.Transparent)
				newSpriteImg = imaging.PasteCenter(newSpriteImg, rotatedSpriteImage)

				newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, newSpriteImg, newX, newY)

			}
		}
	}

	return newBackgroundImg
}

package sprites

import (
	"os"
	// "fmt"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/saenuma/zazabul"

	// "math/rand"
	"image"
	"image/color"
	"math"
	"strconv"

	color2 "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
)

// method4 generates a video with the sprite moving upwards
func Method4(conf zazabul.Config) string {
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
	backgroundImg := imaging.New(videoWidth, videoHeight, backgroundColor)

	totalSeconds := timeFormatToSeconds(conf.Get("video_length"))
	numberOfObjects := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx())

	// load up sprites locations into objectsState
	objectsState := make([]image.Point, 0)
	displacement := 100
	for i := 0; i <= numberOfObjects; i++ {
		if int(math.Mod(float64(i), float64(2))) == 0 {
			newX := i * spriteImg.Bounds().Dx()
			objectsState = append(objectsState, image.Pt(newX, backgroundImg.Bounds().Dy()-displacement))
		} else {
			newX := i * spriteImg.Bounds().Dx()
			objectsState = append(objectsState, image.Pt(newX, backgroundImg.Bounds().Dy()))
		}
	}

	for seconds := 0; seconds < totalSeconds; seconds++ {
		for i := 1; i <= 60; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			toWriteImage := writeCurrentState(backgroundImg, spriteImg, objectsState)
			// update state
			objectsState = updateStateUpwards(backgroundImg, spriteImg, objectsState, numberOfObjects)
			imaging.Save(toWriteImage, outPath)
		}
	}

	return outName
}

func writeCurrentState(backgroundImg, spriteImg image.Image, objectsState []image.Point) *image.NRGBA {
	newBackgroundImg := imaging.New(backgroundImg.Bounds().Dx(), backgroundImg.Bounds().Dy(), color.White)
	newBackgroundImg = imaging.Paste(newBackgroundImg, backgroundImg, image.Pt(0, 0))

	for _, point := range objectsState {
		newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, spriteImg, point.X, point.Y)
	}
	return newBackgroundImg
}

func updateStateUpwards(backgroundImg, spriteImg image.Image, objectsState []image.Point, numberOfObjects int) []image.Point {
	displacement2 := 10

	for i, point := range objectsState {
		newPoint := image.Pt(point.X, point.Y-displacement2)
		objectsState[i] = newPoint
	}

	// append objects if necessary
	almostLastPt := objectsState[len(objectsState)-2]
	lastPt := objectsState[len(objectsState)-1]

	truthValue3 := lastPt.Y+spriteImg.Bounds().Dy() < backgroundImg.Bounds().Dy()
	truthValue4 := almostLastPt.Y+spriteImg.Bounds().Dy() < backgroundImg.Bounds().Dy()

	if truthValue3 && truthValue4 {
		// load up sprites locations into objectsState
		displacement := 100
		for i := 0; i <= numberOfObjects; i++ {
			if int(math.Mod(float64(i), float64(2))) == 0 {
				newX := i * spriteImg.Bounds().Dx()
				objectsState = append(objectsState, image.Pt(newX, backgroundImg.Bounds().Dy()-displacement+50))
			} else {
				newX := i * spriteImg.Bounds().Dx()
				objectsState = append(objectsState, image.Pt(newX, backgroundImg.Bounds().Dy()+50))
			}
		}
	}

	if len(objectsState) > (numberOfObjects * 10) {
		// remove top objects if necessary
		firstPt := objectsState[0]
		secondPt := objectsState[1]
		truthValue1 := firstPt.Y+spriteImg.Bounds().Dy() <= 0
		truthValue2 := secondPt.Y+spriteImg.Bounds().Dy() <= 0
		if truthValue1 || truthValue2 {
			objectsState = append(objectsState[:numberOfObjects], objectsState[numberOfObjects+1:]...)
		}
	}

	return objectsState
}

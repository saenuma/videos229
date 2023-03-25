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
	"strconv"

	color2 "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
)

// method4 generates a video with the sprite moving upwards
func Method4(conf zazabul.Config) string {
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

	totalSeconds, _ := strconv.Atoi(conf.Get("video_length"))
	numberOfObjects := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx())

	// load up sprites locations into objectsState
	objectsState := make([]image.Point, 0)
	increment, _ := strconv.Atoi(conf.Get("increment"))

	for i := 0; i <= numberOfObjects; i++ {
		newX := i * spriteImg.Bounds().Dx()
		objectsState = append(objectsState, image.Pt(newX, backgroundImg.Bounds().Dy()))
	}

	for seconds := 0; seconds < totalSeconds; seconds++ {
		for i := 1; i <= 60; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			toWriteImage := writeCurrentState(backgroundImg, spriteImg, backgroundColor, objectsState)
			// update state
			objectsState = updateStateUpwards(backgroundImg, spriteImg, objectsState, increment, numberOfObjects)
			imaging.Save(toWriteImage, outPath)
		}
	}

	return outName
}

func writeCurrentState(backgroundImg, spriteImg image.Image, backgroundColor color.Color, objectsState []image.Point) *image.NRGBA {
	newBackgroundImg := imaging.New(backgroundImg.Bounds().Dx(), backgroundImg.Bounds().Dy(), backgroundColor)
	newBackgroundImg = imaging.Paste(newBackgroundImg, backgroundImg, image.Pt(0, 0))

	for _, point := range objectsState {
		newBackgroundImg = pasteWithoutTransparentBackground(newBackgroundImg, spriteImg, point.X, point.Y)
	}
	return newBackgroundImg
}

func updateStateUpwards(backgroundImg, spriteImg image.Image, objectsState []image.Point, increment, numberOfObjects int) []image.Point {

	// append objects if necessary
	lastPt := objectsState[len(objectsState)-1]

	shouldAppendBool := lastPt.Y+spriteImg.Bounds().Dy()-10 < backgroundImg.Bounds().Dy()

	if shouldAppendBool {
		// load up sprites locations into objectsState
		for i := 0; i <= numberOfObjects; i++ {
			newX := i * spriteImg.Bounds().Dx()
			objectsState = append(objectsState, image.Pt(newX, backgroundImg.Bounds().Dy()))
		}
	}

	for i, point := range objectsState {
		newPoint := image.Pt(point.X, point.Y-increment)
		objectsState[i] = newPoint
	}

	if len(objectsState) > (numberOfObjects * 10) {
		// remove top objects if necessary
		firstPt := objectsState[0]
		truthValue1 := firstPt.Y+spriteImg.Bounds().Dy() <= 0
		if truthValue1 {
			objectsState = append(objectsState[:numberOfObjects], objectsState[numberOfObjects+1:]...)
		}
	}

	return objectsState
}

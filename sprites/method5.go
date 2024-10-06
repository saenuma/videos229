package sprites

import (
	"image"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kovidgoyal/imaging"
	color2 "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/saenuma/zazabul"
)

// method4 generates a video with the sprite moving downwards
func Method5(conf zazabul.Config) string {
	rootPath, _ := GetRootPath()

	outName := "sp_" + time.Now().Format("20060102T150405")
	renderPath := filepath.Join(rootPath, outName)
	os.MkdirAll(renderPath, 0777)
	totalSeconds, _ := strconv.Atoi(conf.Get("video_length"))

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

	numberOfObjects := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx())

	// load up sprites locations into objectsState
	objectsState := make([]image.Point, 0)
	increment, _ := strconv.Atoi(conf.Get("increment"))
	for i := 0; i <= numberOfObjects; i++ {
		newX := i * spriteImg.Bounds().Dx()
		objectsState = append(objectsState, image.Pt(newX, -spriteImg.Bounds().Dy()))
	}

	for seconds := 0; seconds < totalSeconds; seconds++ {
		for i := 1; i <= 24; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			toWriteImage := writeCurrentState(backgroundImg, spriteImg, backgroundColor, objectsState)
			// update state
			objectsState = updateStateDownwards(backgroundImg, spriteImg, objectsState, increment, numberOfObjects)
			imaging.Save(toWriteImage, outPath)
		}
	}

	return outName
}

func updateStateDownwards(backgroundImg, spriteImg image.Image, objectsState []image.Point, increment, numberOfObjects int) []image.Point {
	// append objects if necessary
	refPt := objectsState[len(objectsState)-1]
	truthValue3 := refPt.Y > 0
	if truthValue3 {
		for i := 0; i <= numberOfObjects; i++ {
			newX := i * spriteImg.Bounds().Dx()
			objectsState = append(objectsState, image.Pt(newX, -spriteImg.Bounds().Dy()))
		}
	}

	for i, point := range objectsState {
		newPoint := image.Pt(point.X, point.Y+increment)
		objectsState[i] = newPoint
	}

	// remove top objects if necessary
	if len(objectsState) > (numberOfObjects * 20) {
		firstPt := objectsState[0]
		truthValue1 := firstPt.Y+spriteImg.Bounds().Dy() > backgroundImg.Bounds().Dy()
		if truthValue1 {
			objectsState = append(objectsState[:numberOfObjects], objectsState[numberOfObjects+1:]...)
		}
	}

	return objectsState
}

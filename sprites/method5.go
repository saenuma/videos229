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
	// "image/color"
	"math"
	"runtime"
	"strconv"
	"sync"

	color2 "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
)

// method4 generates a video with the sprite moving downwards
func Method5(conf zazabul.Config) string {
	rootPath, _ := GetRootPath()

	outName := "s" + time.Now().Format("20060102T150405")
	renderPath := filepath.Join(rootPath, outName)
	os.MkdirAll(renderPath, 0777)
	totalSeconds := timeFormatToSeconds(conf.Get("video_length"))

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

	numberOfObjects := int(backgroundImg.Bounds().Dx() / spriteImg.Bounds().Dx())

	// load up sprites locations into objectsState
	objectsState := make([]image.Point, 0)
	displacement := 10
	for i := 0; i <= numberOfObjects; i++ {
		if int(math.Mod(float64(i), float64(2))) == 0 {
			newX := i * spriteImg.Bounds().Dx()
			objectsState = append(objectsState, image.Pt(newX, displacement-spriteImg.Bounds().Dy()))
		} else {
			newX := i * spriteImg.Bounds().Dx()
			objectsState = append(objectsState, image.Pt(newX, -40-spriteImg.Bounds().Dy()))
		}
	}

	var wg sync.WaitGroup
	numberOfCPUS := runtime.NumCPU()
	jobsPerThread := int(math.Floor(float64(totalSeconds) / float64(numberOfCPUS)))

	for threadIndex := 0; threadIndex < numberOfCPUS; threadIndex++ {
		wg.Add(1)

		startSeconds := threadIndex * jobsPerThread
		endSeconds := (threadIndex + 1) * jobsPerThread

		go func(startSeconds, endSeconds int, objectsState []image.Point, wg *sync.WaitGroup) {
			defer wg.Done()

			// begin pasting and making displacements upwards
			for seconds := 0; seconds < totalSeconds; seconds++ {
				for i := 1; i <= 60; i++ {
					out := (24 * seconds) + i
					outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

					toWriteImage := writeCurrentState(backgroundImg, spriteImg, objectsState)
					// update state
					objectsState = updateStateDownwards(backgroundImg, spriteImg, objectsState, numberOfObjects)
					imaging.Save(toWriteImage, outPath)
				}

			}

		}(startSeconds, endSeconds, objectsState, &wg)

	}
	wg.Wait()

	for seconds := (jobsPerThread * numberOfCPUS); seconds < totalSeconds; seconds++ {
		for i := 1; i <= 60; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			toWriteImage := writeCurrentState(backgroundImg, spriteImg, objectsState)
			// update state
			objectsState = updateStateDownwards(backgroundImg, spriteImg, objectsState, numberOfObjects)
			imaging.Save(toWriteImage, outPath)
		}
	}

	return outName
}

func updateStateDownwards(backgroundImg, spriteImg image.Image, objectsState []image.Point, numberOfObjects int) []image.Point {
	displacement2 := 10

	for i, point := range objectsState {
		newPoint := image.Pt(point.X, point.Y+displacement2)
		objectsState[i] = newPoint
	}

	// append objects if necessary
	almostLastPt := objectsState[len(objectsState)-2]
	lastPt := objectsState[len(objectsState)-1]

	truthValue3 := lastPt.Y > ((spriteImg.Bounds().Dy() / 2) - spriteImg.Bounds().Dy() + 10)
	truthValue4 := almostLastPt.Y > ((spriteImg.Bounds().Dy() / 2) - spriteImg.Bounds().Dy() + 10)

	if truthValue3 && truthValue4 {
		displacement := 10
		for i := 0; i <= numberOfObjects; i++ {
			if int(math.Mod(float64(i), float64(2))) == 0 {
				newX := i * spriteImg.Bounds().Dx()
				objectsState = append(objectsState, image.Pt(newX, displacement-spriteImg.Bounds().Dy()))
			} else {
				newX := i * spriteImg.Bounds().Dx()
				objectsState = append(objectsState, image.Pt(newX, -40-spriteImg.Bounds().Dy()))
			}
		}
	}

	if len(objectsState) > (numberOfObjects * 10) {
		// remove top objects if necessary
		firstPt := objectsState[0]
		secondPt := objectsState[1]
		truthValue1 := firstPt.Y+spriteImg.Bounds().Dy() > backgroundImg.Bounds().Dy()
		truthValue2 := secondPt.Y+spriteImg.Bounds().Dy() > backgroundImg.Bounds().Dy()
		if truthValue1 || truthValue2 {
			objectsState = append(objectsState[:numberOfObjects], objectsState[numberOfObjects+1:]...)
		}
	}

	return objectsState
}

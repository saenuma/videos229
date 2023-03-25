package slideshow

import (
	"image"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	color2 "github.com/gookit/color"
	"github.com/saenuma/videos229/v2shared"
	"github.com/saenuma/zazabul"
)

// immediate appearance slideshow
func Method1(conf zazabul.Config) string {
	rootPath, _ := v2shared.GetRootPath()

	outName := "sl_" + time.Now().Format("20060102T150405")
	renderPath := filepath.Join(rootPath, outName)
	os.MkdirAll(renderPath, 0777)

	fullPicsPath := filepath.Join(rootPath, conf.Get("pictures_dir"))
	if !v2shared.DoesPathExists(fullPicsPath) {
		color2.Red.Printf("The pictures folder '%s' does not exist.\n Exiting.\n", fullPicsPath)
		os.Exit(1)
	}

	videoWidth, _ := strconv.Atoi(conf.Get("video_width"))
	videoHeight, _ := strconv.Atoi(conf.Get("video_height"))

	dirFIs, err := os.ReadDir(fullPicsPath)
	if err != nil {
		color2.Red.Printf("Error reading directory '%s'.\nFull Error: %s\n.Exiting", fullPicsPath, err.Error())
		os.Exit(1)
	}
	picsPaths := make([]string, 0)
	picsBytes := make(map[int]image.Image)
	for i, dirFI := range dirFIs {
		aPicPath := filepath.Join(fullPicsPath, dirFI.Name())
		aPicOpened, _ := imaging.Open(aPicPath)
		if aPicOpened.Bounds().Dx() != videoWidth || aPicOpened.Bounds().Dy() != videoHeight {
			color2.Red.Printf("The width of the picture '%s'\n is not '%s' or the height is not '%s'.\nExiting.\n",
				videoWidth, videoHeight, aPicPath)
			os.Exit(1)
		}
		picsBytes[i] = aPicOpened
		picsPaths = append(picsPaths, aPicPath)
	}

	// var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	switchFrequency, _ := strconv.Atoi(conf.Get("switch_frequency"))
	totalSeconds, _ := strconv.Atoi(conf.Get("video_length"))
	totalThreadsF64 := float64(totalSeconds) / float64(switchFrequency)
	totalThreads := int(math.Floor(totalThreadsF64))

	for threadIndex := 0; threadIndex < totalThreads; threadIndex++ {
		wg.Add(1)

		startSeconds := threadIndex * switchFrequency
		endSeconds := (threadIndex + 1) * switchFrequency

		lengthOfPics := len(picsPaths)
		currentIndexF64 := math.Mod(float64(threadIndex), float64(lengthOfPics))
		currentIndex := int(currentIndexF64)

		go func(startSeconds, endSeconds, currentIndex int, wg *sync.WaitGroup) {
			defer wg.Done()

			for seconds := startSeconds; seconds < endSeconds; seconds++ {
				for i := 1; i <= 24; i++ {
					out := (24 * seconds) + i
					outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

					imaging.Save(picsBytes[currentIndex], outPath)
				}
			}

		}(startSeconds, endSeconds, currentIndex, &wg)
	}
	wg.Wait()

	for seconds := (totalThreads * switchFrequency); seconds < totalSeconds; seconds++ {
		lengthOfPics := len(picsPaths)
		currentIndexF64 := math.Mod(float64(1+(totalThreads*switchFrequency)), float64(lengthOfPics))
		currentIndex := int(currentIndexF64)

		for i := 1; i <= 24; i++ {
			out := (24 * seconds) + i
			outPath := filepath.Join(renderPath, strconv.Itoa(out)+".png")

			imaging.Save(picsBytes[currentIndex], outPath)
		}

	}

	return outName

}

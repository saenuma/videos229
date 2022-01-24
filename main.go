package main

import (
  "os"
  "fmt"
  color2 "github.com/gookit/color"
  "time"
  "path/filepath"
  "github.com/bankole7782/zazabul"
  "github.com/disintegration/imaging"
  "math/rand"
  "image"
  // "image/color"
  "math"
  "strconv"

)

func main() {
  rootPath, err := GetRootPath()
  if err != nil {
    panic(err)
    os.Exit(1)
  }

  if len(os.Args) < 2 {
		color2.Red.Println("Expecting a command. Run with help subcommand to view help.")
		os.Exit(1)
	}


	switch os.Args[1] {
	case "--help", "help", "h":
  		fmt.Println(`videos229 generates videos that could be used for the background of ads
and lyrics videos.

The number of frames per seconds is 60. This is what this program uses.

Directory Commands:
  pwd     Print working directory. This is the directory where the files needed by any command
          in this cli program must reside.

Main Commands:
  init    Creates a config file describing your video. Edit to your own requirements.
          The file from init1 is expected for r1.

  run     Renders a project with the config created above. It expects a blender file and a
          launch file (created from 'init' above)
          All files must be placed in the working directory.

  			`)

	case "pwd":
		fmt.Println(rootPath)

  case "init":
    var	tmplOfMethod1 = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

  	`
		configFileName := "s" + time.Now().Format("20060102T150405") + ".zconf"
		writePath := filepath.Join(rootPath, configFileName)

		conf, err := zazabul.ParseConfig(tmplOfMethod1)
    if err != nil {
    	panic(err)
    }

    err = conf.Write(writePath)
    if err != nil {
      panic(err)
    }

    fmt.Printf("Edit the file at '%s' before launching.\n", writePath)


  case "run":
  	if len(os.Args) != 3 {
  		color2.Red.Println("The run command expects a file created by the init1 command")
  		os.Exit(1)
  	}

  	confPath := filepath.Join(rootPath, os.Args[2])

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

    backgroundImg := image.NewNRGBA(image.Rect(0,0,1366,768))
    backgroundColor := hexToNRGBA(conf.Get("background_color"))

    for x := 0; x < backgroundImg.Bounds().Dx(); x++ {
      for y := 0; y < backgroundImg.Bounds().Dy(); y++ {
        backgroundImg.Set(x, y, backgroundColor)
      }
    }

    rand.Seed(time.Now().UnixNano())

    xOrigin := (backgroundImg.Bounds().Dx() / 2 ) - (spriteImg.Bounds().Dx() / 2)
    yOrigin := (backgroundImg.Bounds().Dy() / 2 ) - (spriteImg.Bounds().Dy() / 2)

    radius := 200

    var tinyAngle float64
    var angleIncrement float64 = float64(0.5)

    for seconds := 0; seconds < 60; seconds++ {

      // if tinyAngle == 360 {
      //   break
      // }
      //
      for i := 1; i <= 60; i++ {
        out := (60 * seconds) + i
        outPath := filepath.Join(renderPath, strconv.Itoa(out) + ".png")

        tinyAngle += angleIncrement
        toWriteImage := writeRotation(backgroundImg, spriteImg, xOrigin, yOrigin, radius, tinyAngle)
        imaging.Save(toWriteImage, outPath)
      }


    }


	default:
		color2.Red.Println("Unexpected command. Run the cli with --help to find out the supported commands.")
		os.Exit(1)
	}

}


func writeRotation(background, sprite image.Image, xOrigin, yOrigin, radius int, angle float64) image.Image {
  // fmt.Println(angle)
  angleInRadians := angle * (math.Pi / 180)
  x := float64(radius) * math.Sin(-angleInRadians)
  y := float64(radius) * math.Cos(-angleInRadians)

  return imaging.Paste(background, sprite, image.Pt(xOrigin + int(x), yOrigin + int(y)))
}

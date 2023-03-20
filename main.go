package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	color2 "github.com/gookit/color"
	"github.com/saenuma/videos229/slideshow"
	"github.com/saenuma/videos229/sprites"
	"github.com/saenuma/videos229/v2shared"
	"github.com/saenuma/zazabul"
)

const VersionFormat = "20060102T150405MST"

func main() {
	rootPath, err := v2shared.GetRootPath()
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		color2.Red.Println("Expecting a command. Run with help subcommand to view help.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--help", "help", "h":
		fmt.Println(`videos229 generates videos that could be used for the background of adverts
and lyrics videos.

Directory Commands:
  pwd     Print working directory. This is the directory where the files needed by any command
          in this cli program must reside.

Main Commands:
  initsp    Initialize Sprites Video. Creates a config file describing your video.
            Edit to your own requirements.

  initsl    Initialize Slideshow Video. Creates a config file describing your video.
            Edit to your own requirements.

  run       Renders a project with the config created above. It expects a config file
            created from either 'initsp' or 'initsl' above. Command 'run' would
            generate an mp4 video.
            All files must be placed in the working directory.

  			`)

	case "pwd":
		fmt.Println(rootPath)

	case "initsp":
		var tmplOfMethod1 = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_length is the length of the output video in this format (mm:ss)
video_length:

// method. The method are in numbers. Allowed values are 1, 2, 3, 4, 5.
// 1: for movement around a circle style
// 2: for disappearing pattern style
// 3: for rotation in place style
// 4: for upward movement
// 5: for downward movement
method: 1

  	`
		configFileName := "sp_" + time.Now().Format("20060102T150405") + ".zconf"
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

	case "initsl":
		var tmplOfMethod1 = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// The directory containing the pictures for a slideshow. It must be stored in the working directory
// of videos229.
// All pictures here must be of width 1366px and height 768px
pictures_dir:

// video_length is the length of the output video in this format (mm:ss)
video_length:

// method. The method are in numbers. Allowed values are 1
// 1: for immediate appearance slideshow
// 2: for fade in slideshow
method: 1

  	`
		configFileName := "sl_" + time.Now().Format("20060102T150405") + ".zconf"
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
		rootPath, _ := v2shared.GetRootPath()

		if len(os.Args) != 3 {
			color2.Red.Println("The run command expects a file created by the init command")
			os.Exit(1)
		}

		confFileName := os.Args[2]
		confPath := filepath.Join(rootPath, confFileName)

		conf, err := zazabul.LoadConfigFile(confPath)
		if err != nil {
			panic(err)
		}

		for _, item := range conf.Items {
			if item.Value == "" {
				color2.Red.Println("Every field in the launch file is compulsory.")
				os.Exit(1)
			}
		}

		var outName string

		if strings.HasPrefix(confFileName, "sp_") {
			if conf.Get("method") == "1" {
				outName = sprites.Method1(conf)
			} else if conf.Get("method") == "2" {
				outName = sprites.Method2(conf)
			} else if conf.Get("method") == "3" {
				outName = sprites.Method3(conf)
			} else if conf.Get("method") == "4" {
				outName = sprites.Method4(conf)
			} else if conf.Get("method") == "5" {
				outName = sprites.Method5(conf)
			} else {
				color2.Red.Println("The method code is invalid.")
				os.Exit(1)
			}

		} else if strings.HasPrefix(confFileName, "sl_") {
			if conf.Get("method") == "1" {
				outName = slideshow.Method1(conf)
			} else if conf.Get("method") == "2" {
				outName = slideshow.Method2(conf)
			}
		}

		fmt.Println("Finished generating frames.")

		command := v2shared.GetFFMPEGCommand()

		out, err := exec.Command(command, "-framerate", "24", "-i", filepath.Join(rootPath, outName, "%d.png"),
			"-pix_fmt", "yuv420p",
			filepath.Join(rootPath, outName+".mp4")).CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			panic(err)
		}

		os.RemoveAll(filepath.Join(rootPath, outName))
		fmt.Println("View the generated video at: ", filepath.Join(rootPath, outName+".mp4"))

	default:
		color2.Red.Println("Unexpected command. Run the cli with --help to find out the supported commands.")
		os.Exit(1)
	}

}

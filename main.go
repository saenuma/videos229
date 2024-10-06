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

var configTemplates map[string]string = make(map[string]string)

func main() {
	rootPath, err := v2shared.GetRootPath()
	if err != nil {
		panic(err)
	}

	// register functions
	sprites.RegisterAll(&configTemplates)
	slideshow.RegisterAll(&configTemplates)

	if len(os.Args) < 2 {
		color2.Red.Println("Expecting a command. Run with help subcommand to view help.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--help", "help", "h":
		fmt.Println(`videos229 generates videos that could be used for the background of adverts.

Directory Commands:
  pwd     Print working directory. This is the directory where the files needed by any command
          in this cli program must reside.

Main Commands:
  lf        Lists the available folders

  lm        Lists the methods in a folder. Expects the folder name as the only arguments

  br        Begins a render. It expects the folder/number combo.
            For example: videos229 br sprites/1

  run       Renders a project with the config created above. It expects a config file
            created from 'br' above. Command 'run' would generate an mp4 video.
            All files must be placed in the working directory.

  			`)

	case "pwd":
		fmt.Println(rootPath)

	case "lf":
		fmt.Println("sprites")
		fmt.Println("slideshows")

	case "lm":
		if os.Args[2] == "sprites" {
			fmt.Println(`Method Numbers in Folder 'Sprites':
  1     sprites dancing round a circle
  2     disappearing pattern style
  3     rotation in place style
  4     upward movements of sprites
  5     downward movements of sprites
			`)

		} else {
			fmt.Println(`Method Numbers in Folder 'Sprites':
  1     immediate appearance slideshow
  2     fade in slideshow
			`)

		}

	case "br":
		inputMethod := os.Args[2]
		tmpl, ok := configTemplates[os.Args[2]]
		if ok {
			stub := strings.ReplaceAll(inputMethod, "/", "_") + "_"
			configFileName := stub + time.Now().Format("20060102T150405") + ".zconf"
			writePath := filepath.Join(rootPath, configFileName)

			conf, err := zazabul.ParseConfig(tmpl)
			if err != nil {
				panic(err)
			}

			err = conf.Write(writePath)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Edit the file at '%s' before launching.\n", writePath)

		} else {
			color2.Red.Println("Invalid method: " + inputMethod)
			return
		}

	case "run":
		startTime := time.Now()
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
		confFilenameparts := strings.Split(confFileName, "_")
		if strings.HasPrefix(confFileName, "sprites_") {

			if confFilenameparts[1] == "1" {
				outName = sprites.Method1(conf)
			} else if confFilenameparts[1] == "2" {
				outName = sprites.Method2(conf)
			} else if confFilenameparts[1] == "3" {
				outName = sprites.Method3(conf)
			} else if confFilenameparts[1] == "4" {
				outName = sprites.Method4(conf)
				// } else if confFilenameparts[1] == "5" {
				// 	outName = sprites.Method5(conf)
			} else {
				color2.Red.Println("The method code is invalid.")
				os.Exit(1)
			}

		} else if strings.HasPrefix(confFileName, "slideshows_") {
			if confFilenameparts[1] == "1" {
				outName = slideshow.Method1(conf)
			} else if confFilenameparts[1] == "2" {
				outName = slideshow.Method2(conf)
			}
		}

		fmt.Println("Finished generating frames.")

		command := v2shared.GetFFMPEGCommand()

		outPath := filepath.Join(rootPath, "video_"+time.Now().Format("20060102T150405")+".mp4")
		out, err := exec.Command(command, "-framerate", "24", "-i", filepath.Join(rootPath, outName, "%d.png"),
			"-pix_fmt", "yuv420p", outPath).CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			panic(err)
		}

		os.RemoveAll(filepath.Join(rootPath, outName))

		fmt.Printf("took %ds\n", int(time.Since(startTime).Seconds()))
		fmt.Println("View the generated video at: ", outPath)

	default:
		color2.Red.Println("Unexpected command. Run the cli with --help to find out the supported commands.")
		os.Exit(1)
	}

}

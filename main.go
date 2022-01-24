package main

import (
  "os"
  "fmt"
  color2 "github.com/gookit/color"
  "time"
  "path/filepath"
  "github.com/bankole7782/zazabul"
  "os/exec"
  "strings"
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
  		fmt.Println(`videos229 generates videos that could be used for the background of adverts
and lyrics videos.

The number of frames per seconds is 60. This is what this program uses.

Directory Commands:
  pwd     Print working directory. This is the directory where the files needed by any command
          in this cli program must reside.

Main Commands:
  init    Creates a config file describing your video. Edit to your own requirements.
          The file from init is expected for trun and frun.

  run     Renders a project with the config created above. It expects a config file (created from 'init' above)
          run would generate an mp4 video.
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
    outName := method1(os.Args)
    fmt.Println("Finished generating frames.")

    begin := os.Getenv("SNAP")
    command := "ffmpeg"
    if begin != "" && ! strings.HasPrefix(begin, "/snap/go/") {
      command = filepath.Join(begin, "bin", "ffmpeg")
    }

    out, err := exec.Command(command, "-framerate", "60", "-i", filepath.Join(rootPath, outName, "%d.png"),
      filepath.Join(rootPath, outName + ".mp4")).CombinedOutput()
    if err != nil {
      fmt.Println(string(out))
      panic(err)
    }

    os.RemoveAll(filepath.Join(rootPath, outName))
    fmt.Println("View the generated video at: ", filepath.Join(rootPath, outName + ".mp4"))

	default:
		color2.Red.Println("Unexpected command. Run the cli with --help to find out the supported commands.")
		os.Exit(1)
	}

}

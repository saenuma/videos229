package main

import (
  "os"
  "fmt"
  color2 "github.com/gookit/color"
  "time"
  "path/filepath"
  "github.com/saenuma/zazabul"
  "os/exec"
  "strings"
  "runtime"
  "io"
  "net/http"
  "github.com/saenuma/videos229/sprites"
  "github.com/saenuma/videos229/slideshow"
  v229s "github.com/saenuma/videos229/videos229_shared"
)


const VersionFormat = "20060102T150405MST"

func main() {
  rootPath, err := v229s.GetRootPath()
  if err != nil {
    panic(err)
    os.Exit(1)
  }

  if len(os.Args) < 2 {
		color2.Red.Println("Expecting a command. Run with help subcommand to view help.")
		os.Exit(1)
	}

  if runtime.GOOS == "windows" {
    newVersionStr := ""
    resp, err := http.Get("https://sae.ng/static/wapps/videos229.txt")
    if err != nil {
      fmt.Println(err)
    }
    if err == nil {
      defer resp.Body.Close()
      body, err := io.ReadAll(resp.Body)
      if err == nil && resp.StatusCode == 200 {
        newVersionStr = string(body)
      }
    }

    newVersionStr = strings.TrimSpace(newVersionStr)
    currentVersionStr = strings.TrimSpace(currentVersionStr)

    hnv := false
    if newVersionStr != "" && newVersionStr != currentVersionStr {
      time1, err1 := time.Parse(VersionFormat, newVersionStr)
      time2, err2 := time.Parse(VersionFormat, currentVersionStr)

      if err1 == nil && err2 == nil && time2.Before(time1) {
        hnv = true
      }
    }

    if hnv == true {
      fmt.Println("videos229 has an update.")
      fmt.Println("please visit 'https://sae.ng/videos229' for update instructions." )
      fmt.Println()
    }

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
    var	tmplOfMethod1 = `// background_color is the color of the background image. Example is #af1382
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


  case "initsl":
    var	tmplOfMethod1 = `// background_color is the color of the background image. Example is #af1382
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
    rootPath, _ := v229s.GetRootPath()

    if len(os.Args) != 3 {
      color2.Red.Println("The run command expects a file created by the init command")
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

    var outName string
    if conf.Get("sprite_file") != "" {
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

    } else if conf.Get("pictures_dir") != "" {
      if conf.Get("method") == "1" {
        outName = slideshow.Method1(conf)
      } else if conf.Get("method") == "2" {
        outName = slideshow.Method2(conf)
      }
    }

    fmt.Println("Finished generating frames.")

    command := v229s.GetFFMPEGCommand()

    out, err := exec.Command(command, "-framerate", "60", "-i", filepath.Join(rootPath, outName, "%d.png"),
      "-pix_fmt",  "yuv420p",
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

package videos229_shared

import (
  "os"
  "strings"
  "github.com/pkg/errors"
  "path/filepath"
  "strconv"
)



func GetRootPath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "os error")
	}
  dd := filepath.Join(hd, "Videos229")
  os.MkdirAll(dd, 0777)

	return dd, nil
}



func TimeFormatToSeconds(s string) int {
  // calculate total duration of the song
  parts := strings.Split(s, ":")
  minutesPartConverted, err := strconv.Atoi(parts[0])
  if err != nil {
    panic(err)
  }
  secondsPartConverted, err := strconv.Atoi(parts[1])
  if err != nil {
    panic(err)
  }
  totalSecondsOfSong := (60 * minutesPartConverted) + secondsPartConverted
  return totalSecondsOfSong
}


func DoesPathExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}


func GetFFMPEGCommand() string {
  // get the right ffmpeg command
  homeDir, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }

  devPath := filepath.Join(homeDir, "bin", "ffmpeg.exe")
  bundledPath := filepath.Join("C:\\Program Files (x86)\\Videos229", "ffmpeg.exe")
  if DoesPathExists(devPath) {
    return devPath
  }

  return bundledPath
}

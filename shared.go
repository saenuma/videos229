package main

import (
  "os"
  "strings"
  "github.com/pkg/errors"
  "path/filepath"
  "strconv"
  // "fmt"
)



func GetRootPath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "os error")
	}
	dd := os.Getenv("SNAP_USER_COMMON")
	if strings.HasPrefix(dd, filepath.Join(hd, "snap", "go")) || dd == "" {
		dd = filepath.Join(hd, "videos229_data")
    os.MkdirAll(dd, 0777)
	}

	return dd, nil
}



func timeFormatToSeconds(s string) int {
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

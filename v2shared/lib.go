package v2shared

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func GetRootPath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "os error")
	}

	dd := os.Getenv("SNAP_USER_COMMON")
	if strings.HasPrefix(dd, filepath.Join(hd, "snap", "go")) || dd == "" {
		dd = filepath.Join(hd, "Videos229")
		os.MkdirAll(dd, 0777)
	}

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
	totalSecondsOfSong := (24 * minutesPartConverted) + secondsPartConverted
	return totalSecondsOfSong
}

func DoesPathExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFFMPEGCommand() string {
	if runtime.GOOS == "windows" {
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
	} else {
		var cmdPath string
		begin := os.Getenv("SNAP")
		cmdPath = "ffmpeg"
		if begin != "" && !strings.HasPrefix(begin, "/snap/go/") {
			cmdPath = filepath.Join(begin, "bin", "ffmpeg")
		}

		return cmdPath
	}

}

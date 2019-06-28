package picam

import (
	"fmt"
	"os/exec"
	"time"
)

type Cam struct {
	Path string
}

// TakePicture takes a picture and saves it.
// If filename is empty, it will generate a filename with a timestamp
func (r *Cam) TakePicture(filename string) (string, error) {
	var fullPath string
	if filename != "" {
		fullPath = r.Path + "/" + filename
	} else {
		fullPath = r.Path + "/" + buildFilename()
	}
	cmd := exec.Command("raspistill", "-o", fullPath)
	e := cmd.Run()
	if e != nil {
		return "", e
	}
	return fullPath, nil
}

func NewCam(path string) *Cam {
	return &Cam{Path: path}
}

func buildFilename() string {
	now := time.Now()
	return fmt.Sprintf("picture-%04d%02d%02d%02d%02d%02d%03d.jpg", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond() / 1000000)
}

package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

type facedetector struct {
	vidfeed        *gocv.VideoCapture
	window         *gocv.Window
	classifierFile string // xml file
	classifier     *gocv.CascadeClassifier
	timer          *timer
	lastErr        string
}

func newFaceDetector(deviceID int) error /* *facedetector */ {
	var f = new(facedetector)
	cam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		return err
	}
	f.vidfeed = cam
	f.classifierFile = "lbpcascade_frontalface_improved.xml"
	var c = gocv.NewCascadeClassifier()
	f.classifier = &c
	if !f.classifier.Load(f.classifierFile) {
		// f.lastErr = fmt.Errorf("error loading classifier file %s\n", f.classifierFile).Error()
		return fmt.Errorf("error loading classifier file %s\n", f.classifierFile)
	}
	f.timer = newTimer()
	return f

}

func (fd *facedetector) Error() string {
	return fd.lastErr
}

func (fd *facedetector) Start() {
	if fd.window == nil {
		fd.window = gocv.NewWindow("timeclock")
	}
	m := gocv.NewMat()
	for fd.vidfeed.Read(&m) {
		fd.classifier.DetectMultiScale(m)
	}

}
func (fd *facedetector) Show() *gocv.Window {
	return fd.window
}

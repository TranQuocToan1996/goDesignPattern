package anhthi

import "testing"

func TestImage(t *testing.T) {
	highResolutionImageProxy := NewImageProxy("sample/img1.png")
	// the realImage won't be loaded until user calls display on proxy
	// later
	highResolutionImageProxy.display()
}

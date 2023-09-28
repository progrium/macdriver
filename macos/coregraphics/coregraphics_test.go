package coregraphics

import "testing"

func TestCoreGraphicsValid(t *testing.T) {}

func TestDisplayCreateImage(t *testing.T) {
	id := MainDisplayID()
	t.Logf("MainDisplayID: %d", id)
	img := DisplayCreateImage(id)
	t.Logf("DisplayCreateImage: %v", img)
}

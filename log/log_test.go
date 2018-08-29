package log

import (
	"testing"
)

func TestTitleConsole(t *testing.T) {
	ten := TitleConsole("Chao Ban")

	Log(ten)
}

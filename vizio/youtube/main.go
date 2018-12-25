package youtube

import (
	"time"

	"github.com/midn/alexa/vizio/key_command"
	"github.com/midn/alexa/vizio/volume"
)

func OpenAndPlay() {
	// Last menu could be something else, first call to Home doesn't open SmartCast, it just turns on TV
	key_command.Home()
	// It takes about 15s to open SmartCast menu with animation loading
	key_command.Home()
	time.Sleep(15 * time.Second)

	// Just incase Vizio menu was frozen, go left 10x times :shrug:
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()
	key_command.Left()

	// Self explanatory
	key_command.Down()
	key_command.Down()
	key_command.Down()

	// Youtube is 3rd from the left
	key_command.Right()
	key_command.Right()

	time.Sleep(1 * time.Second)
	key_command.Click()

	// Youtube opening takes about 12s
	time.Sleep(12 * time.Second)
	volume.SetVolume(8)
	key_command.Click()
}

func NextVideo() {
	key_command.Click()
	time.Sleep(1 * time.Second)
	key_command.Right()
	key_command.Click()
}

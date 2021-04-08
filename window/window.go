package window

import (
	"github.com/elizarpif/cryptoswitch"
	"github.com/elizarpif/cryptoswitch-ui/ui"
)

type Window struct {
	uiWindow *ui.UICryptoswitchMainWindow

	privKey *cryptoswitch.PrivateKey
	file    *File
}

type File struct {
	in  string
	out string
}

func NewWindow(ui *ui.UICryptoswitchMainWindow) *Window {
	return &Window{
		uiWindow: ui,
	}
}

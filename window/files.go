package window

import (
	"io/ioutil"
	"strings"

	"github.com/therecipe/qt/widgets"
)

func (w *Window) selectFile() string {
	filename := widgets.NewQFileDialog2(nil, "Open Dialog", "", "").
		GetOpenFileName(nil, "", "", "", "", 0)

	return filename
}

func (w *Window) SelectInFile() {
	filename := w.selectFile()
	if filename == "" {
		return
	}

	if w.file != nil {
		w.file.in = filename
	} else {
		w.file = &File{in: filename}
	}

	w.uiWindow.LabelInFile.SetText(filename)
	//w.uiWindow.LabelOutFile.SetText(filename)
}

func (w *Window) SelectOutFile() {
	filename := w.selectFile()
	if filename == "" {
		return
	}

	if w.file != nil {
		w.file.out = filename
	} else {
		w.file = &File{out: filename}
	}

	w.uiWindow.LabelOutFile.SetText(filename)
}

func openFile(filename string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func writeFile(data []byte, filename string) error {
	return ioutil.WriteFile(filename, data, 0777)
}

func fileNameWithoutExtension(fileName string) string {
	ff := strings.TrimSuffix(fileName, ".enc")
	return ff
}

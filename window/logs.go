package window

import (
	"fmt"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (w *Window) addErrLog(err error) {
	if err == nil {
		return
	}

	msg := widgets.NewQMessageBox2(widgets.QMessageBox__Warning, "Ошибка!", err.Error(), widgets.QMessageBox__Ok, w.uiWindow.Centralwidget, core.Qt__Widget)
	msg.Show()
	w.addLog(err.Error())
}

func (w *Window) addLog(msg string) {
	//box := widgets.NewQMessageBox2(widgets.QMessageBox__Information, "Ошибка!", msg, widgets.QMessageBox__Ok, w.uiWindow.Centralwidget, core.Qt__Widget)
	//box.Show()

	str := fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), msg)
	w.uiWindow.Logs.Append(str)
}

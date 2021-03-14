package window

import (
	"fmt"
	"log"
	"time"
)

func (w *Window) addErrLog(err error) {
	if err == nil {
		return
	}

	w.addLog(err.Error())

}

func (w *Window) addLog(msg string) {
	log.Println(msg)
	str := fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), msg)
	w.uiWindow.Logs.Append(str)
}

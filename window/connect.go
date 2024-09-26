package window

import (
	"fmt"
	"runtime"
	"unicode/utf8"

	"github.com/elizarpif/cryptoswitch"
)

func (w *Window) Connect() {
	ww := w.uiWindow

	ww.EncryptBtn.ConnectClicked(func(checked bool) {
		go w.Encrypt()
	})
	ww.DecryptBtn.ConnectClicked(func(checked bool) {
		go w.Decrypt()
	})

	ww.GenerateBtn.ConnectClicked(func(checked bool) {
		go w.GenerateKey()
	})

	ww.PlainText.ConnectTextChanged(func() {
		w.countPlainTextSymbols()
	})
	ww.CipherText.ConnectTextChanged(func() {
		w.countCipherTextSymbols()
	})

	ww.SelectInFileBtn.ConnectClicked(func(checked bool) {
		w.SelectInFile()
	})

	ww.CipherSwitch.ConnectCurrentIndexChanged(func(index int) {
		if index == int(cryptoswitch.DES) {
			ww.CbcRadio.SetChecked(true)
			ww.GcmRadio.SetEnabled(false)
		} else {
			ww.GcmRadio.SetEnabled(true)
		}
	})

	ww.SelectOutFileBtn.ConnectClicked(func(checked bool) {
		w.SelectOutFile()
	})

	//w.uiWindow.Menubar = widgets.NewQMenuBar(w.uiWindow.Centralwidget)
	//menu := widgets.NewQMenu2("Меню", nil)
	//menu.AddAction("Some")
	//
	//w.uiWindow.Menubar.AddMenu(menu)
}

func (w *Window) selectCipher() cryptoswitch.Cipher {
	index := w.uiWindow.CipherSwitch.CurrentIndex()
	switch index {
	case 0:
		return cryptoswitch.AES
	case 1:
		return cryptoswitch.DES
	case 2:
		return cryptoswitch.Camellia
	case 3:
		return cryptoswitch.Twofish
	}

	w.uiWindow.Logs.Append("Неизвестный шифр, использую AES")
	return 0
}

func (w *Window) selectMode() cryptoswitch.Mode {
	if w.uiWindow.GcmRadio.IsChecked() {
		return cryptoswitch.GCM
	}

	return cryptoswitch.CBC
}

func (w *Window) GenerateKey() {
	var err error
	defer w.addErrLog(err)

	privKey, err := cryptoswitch.GenerateKey()
	if err != nil {
		return
	}

	w.privKey = privKey

	w.uiWindow.ParamXEdit.SetText(privKey.X.String())
	w.uiWindow.ParamYEdit.SetText(privKey.Y.String())
}

func (w *Window) Encrypt() {
	if w.uiWindow.TabWidget.CurrentIndex() == 0 {
		w.EncryptText()
	} else {
		w.addLog("Шифрование файла начинается")
		w.EncryptFile()
		runtime.GC()
	}
}

func (w *Window) Decrypt() {
	if w.uiWindow.TabWidget.CurrentIndex() == 0 {
		w.DecryptText()
	} else {
		w.addLog("Дешифрование файла начинается")
		w.DecryptFile()
		runtime.GC()
	}
}

func (w *Window) countCipherTextSymbols() {
	text := w.uiWindow.CipherText.ToPlainText()

	n := w.countSymbols(text)

	w.uiWindow.LabelCountCipherText.SetText(
		fmt.Sprintf("%s %d %s", countText, n, incline(n)),
	)
}

func (w *Window) countPlainTextSymbols() {
	text := w.uiWindow.PlainText.ToPlainText()

	n := w.countSymbols(text)

	w.uiWindow.LabelCountPlainText.SetText(
		fmt.Sprintf("%s %d %s", countText, n, incline(n)),
	)
}

func (w *Window) countSymbols(text string) int {
	return utf8.RuneCountInString(text)
}

const countText = "Размер введенного текста:"

var countSymbols = [3]string{"символ", "символа", "символов"}

// склонение существительного после числительного
func incline(n int) string {
	n %= 100
	n1 := n % 10

	if n > 10 && n < 20 {
		return countSymbols[2]
	}
	if n1 > 1 && n1 < 5 {
		return countSymbols[1]
	}
	if n1 == 1 {
		return countSymbols[0]
	}

	return countSymbols[2]
}

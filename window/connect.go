package window

import (
	"fmt"

	"github.com/elizarpif/diploma-elliptic/cryptoswitch"
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

	ww.SelectOutFileBtn.ConnectClicked(func(checked bool) {
		w.SelectOutFile()
	})
}

func (w *Window) selectCipher() cryptoswitch.Cipher {
	index := w.uiWindow.CipherSwitch.CurrentIndex()
	switch index {
	case 0:
		return cryptoswitch.AES
	case 1:
		return cryptoswitch.Camellia
	case 2:
		return cryptoswitch.Twofish
	}

	w.uiWindow.Logs.Append("Неизвестный шифр, использую AES")
	return 0
}

func (w *Window) selectMode() cryptoswitch.Mode {
	index := w.uiWindow.CipherMode.CurrentIndex()
	switch index {
	case 0:
		return cryptoswitch.GCM
	case 1:
		return cryptoswitch.CBC
	}

	w.uiWindow.Logs.Append("Неизвестный режим шифрования, использую GCM")
	return 0
}

func (w *Window) countCipherTextSymbols() {
	n := len(w.uiWindow.CipherText.ToPlainText())

	w.uiWindow.LabelCountCipherText.SetText(
		fmt.Sprintf("%s %d %s", countText, n, incline(n)),
	)
}
func (w *Window) countPlainTextSymbols() {
	n := len(w.uiWindow.PlainText.ToPlainText())

	w.uiWindow.LabelCountPlainText.SetText(
		fmt.Sprintf("%s %d %s", countText, n, incline(n)),
	)
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
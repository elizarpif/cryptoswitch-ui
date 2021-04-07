package window

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/widgets"

	"github.com/elizarpif/cryptoswitch"
)

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

func (w *Window) EncryptText() {
	data := []byte(w.uiWindow.PlainText.ToPlainText())
	if len(data) == 0 {
		return
	}

	if w.privKey == nil {
		w.addLog("отсутствует ключ шифрования")
		return
	}

	sw := cryptoswitch.NewCryptoSwitch(w.selectCipher(), w.selectMode())

	encrypt, err := sw.Encrypt(w.privKey.PublicKey, &data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	w.uiWindow.CipherText.SetText(hex.EncodeToString(*encrypt))
}

func (w *Window) progress(pbar *widgets.QProgressDialog, fileLen int, done chan bool, ticker *time.Ticker) {

	for {

		select {
		case <-done:
			fmt.Println("ticker done")
			pbar.SetValue(pbar.Maximum())
			pbar.Close()
			return
		case t := <-ticker.C:
			if pbar.WasCanceled() {
				continue
			}

			pbar.SetValue(pbar.Value() + fileLen/70)
			fmt.Printf("Tick at %d , value %d\n", t, pbar.Value())
		}
	}
}

func (w *Window) EncryptFile() {
	if w.file == nil {
		w.addLog("Входной и выходной файлы не выбраны")
		return
	}

	filename := w.file.in

	data, err := openFile(filename)
	if err != nil {
		w.addLog(fmt.Sprintf("Не получается открыть файл: %s", filename))
		return
	}

	dataLen := len(*data)
	if dataLen == 0 {
		return
	}

	if w.privKey == nil {
		w.addLog("Отсутствует ключ шифрования")
		return
	}

	ticker := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)
	pbar := widgets.NewQProgressDialog2("Шифрование...", "Отмена", 0, dataLen, w.uiWindow.Centralwidget, core.Qt__Dialog)
	pbar.ConnectCanceled(func() {
		w.addLog("Шифрование остановлено")
	})

	go w.progress(pbar, dataLen, done, ticker)

	sw := cryptoswitch.NewCryptoSwitch(w.selectCipher(), w.selectMode())

	encrypt, err := sw.Encrypt(w.privKey.PublicKey, data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	if pbar.WasCanceled() {
		return
	}

	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	fileOut := w.file.out
	if fileOut == "" {
		fileOut = filename + ".enc"
	}

	err = writeFile(encrypt, fileOut)
	if err != nil {
		w.addErrLog(err)
		w.addLog("Не получается записать файл")
		return
	}

	w.addLog(fmt.Sprintf("Файл успешно зашифрован и записан в %s", fileOut))
}

func (w *Window) Encrypt() {
	if w.uiWindow.TabWidget.CurrentIndex() == 0 {
		w.EncryptText()
	} else {
		w.addLog("Шифрование файла начинается")
		w.EncryptFile()
	}
}

func (w *Window) DecryptFile() {
	if w.file == nil {
		w.addLog("Входной и выходной файлы не выбраны")
		return
	}

	filename := w.file.in

	data, err := openFile(filename)
	if err != nil {
		w.addLog(fmt.Sprintf("Не получается открыть файл: %s", filename))
		return
	}

	if len(*data) == 0 {
		return
	}

	if w.privKey == nil {
		w.addLog("отсутствует ключ шифрования")
		return
	}

	sw := cryptoswitch.NewCryptoSwitch(w.selectCipher(), w.selectMode())

	encrypt, err := sw.Decrypt(w.privKey, data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	if w.stopCipher {
		w.stopCipher = false
		w.addLog("Шифрование остановлено")
		return
	}

	fileOut := w.file.out
	if fileOut == "" {
		fileOut = fileNameWithoutExtension(filename)
		w.addLog(fmt.Sprintf("Выходной файл не указан, будет создан файл с названием %s", fileOut))
	}

	err = writeFile(encrypt, fileOut)
	if err != nil {
		w.addErrLog(err)
		w.addLog("Не получается записать файл")
		return
	}

	w.addLog("Файл успешно расшифрован")
}

func (w *Window) DecryptText() {
	data, err := hex.DecodeString(w.uiWindow.PlainText.ToPlainText())
	if err != nil {
		w.addErrLog(err)
		return
	}

	if w.privKey == nil {
		w.addLog("отсутствует ключ шифрования")
		return
	}

	sw := cryptoswitch.NewCryptoSwitch(w.selectCipher(), w.selectMode())

	encrypt, err := sw.Decrypt(w.privKey, &data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	w.uiWindow.CipherText.SetText(string(*encrypt))
}

func (w *Window) Decrypt() {
	if w.uiWindow.TabWidget.CurrentIndex() == 0 {
		w.DecryptText()
	} else {
		w.addLog("Дешифрование файлы начинается")
		w.DecryptFile()
	}
}

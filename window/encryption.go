package window

import (
	"encoding/hex"
	"fmt"

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
	w.uiWindow.SecretKey.SetText(hex.EncodeToString(privKey.D.Bytes()))
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

	encrypt, err := sw.Encrypt(w.privKey.PublicKey, data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	w.uiWindow.CipherText.SetText(hex.EncodeToString(encrypt))
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

	if len(data) == 0 {
		return
	}

	if w.privKey == nil {
		w.addLog("отсутствует ключ шифрования")
		return
	}

	sw := cryptoswitch.NewCryptoSwitch(w.selectCipher(), w.selectMode())

	encrypt, err := sw.Encrypt(w.privKey.PublicKey, data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	fileOut := w.file.out
	if fileOut == "" {
		w.addLog(fmt.Sprintf("Выходной файл не указан, будет создан файл с названием %s%s", filename, ".enc"))
		fileOut = filename + ".enc"
	}

	if w.stopCipher {
		w.addLog("Шифрование остановлено")
		return
	}

	err = writeFile(encrypt, fileOut)
	if err != nil {
		w.addErrLog(err)
		w.addLog("Не получается записать файл")
		return
	}

	w.addLog("Файл успешно зашифрован")
}

func (w *Window) Encrypt() {
	if w.uiWindow.TabWidget.CurrentIndex() == 0 {
		w.EncryptText()
	} else {
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

	if len(data) == 0 {
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

	fileOut := w.file.out
	if fileOut == "" {
		fileOut = fileNameWithoutExtension(filename)
		w.addLog(fmt.Sprintf("Выходной файл не указан, будет создан файл с названием %s", fileOut))
	}

	if w.stopCipher {
		w.addLog("Шифрование остановлено")
		return
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

	encrypt, err := sw.Decrypt(w.privKey, data)
	if err != nil {
		w.addErrLog(err)
		return
	}

	w.uiWindow.CipherText.SetText(string(encrypt))
}

func (w *Window) Decrypt() {
	if w.uiWindow.TabWidget.CurrentIndex() == 0 {
		w.DecryptText()
	} else {
		w.DecryptFile()
	}
}

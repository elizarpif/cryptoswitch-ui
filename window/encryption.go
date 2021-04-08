package window

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/elizarpif/cryptoswitch"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (w *Window) EncryptText() {
	w.operationText(true)
}

func (w *Window) DecryptText() {
	w.operationText(false)
}

func getTextData(text string, encrypt bool) (*[]byte, error) {
	if encrypt {
		data := []byte(text)
		return &data, nil
	}

	data, err := hex.DecodeString(text)
	return &data, err
}

func resultTextData(text *[]byte, encrypt bool) string {
	if encrypt {
		return hex.EncodeToString(*text)
	}

	return string(*text)
}

func (w *Window) operationText(encrypt bool) {
	data, err := getTextData(w.uiWindow.PlainText.ToPlainText(), encrypt)
	if err != nil {
		w.addErrLog(err)
		return
	}

	if w.privKey == nil {
		w.addLog("отсутствует ключ шифрования")
		return
	}

	encrypted, err := w.operation(data, encrypt)
	if err != nil {
		w.addErrLog(err)
		return
	}

	w.uiWindow.CipherText.SetText(resultTextData(encrypted, encrypt))
}

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
)

func getTickerTime(dataLen int, cipher cryptoswitch.Cipher) time.Duration {
	t := 4
	if cipher == cryptoswitch.AES {
		t = 1
	}

	if dataLen >= GIGABYTE {
		return time.Duration(200*t) * time.Millisecond
	}
	if dataLen >= MEGABYTE {
		return time.Duration(80*t) * time.Millisecond
	}
	return time.Millisecond
}

func (w *Window) getOutFilename(filename string, encrypt bool) string {
	fileOut := w.file.out
	if fileOut != "" {
		return fileOut
	}

	if encrypt {
		return filename + ".enc"
	}

	return fileNameWithoutExtension(filename)
}

// шифрование либо дешифрование
func (w *Window) operation(data *[]byte, encrypt bool) (*[]byte, error) {
	sw := cryptoswitch.NewCryptoSwitch(w.selectCipher(), w.selectMode())
	if encrypt {
		return sw.Encrypt(w.privKey.PublicKey, data)
	}

	return sw.Decrypt(w.privKey, data)
}

func resultLog(encrypt bool, fileOut string) string {
	if encrypt {
		return fmt.Sprintf("Файл успешно зашифрован и записан в %s", fileOut)
	}

	return fmt.Sprintf("Файл успешно расшифрован и записан в %s", fileOut)
}

func operationType(encrypt bool) string {
	if encrypt {
		return "Шифрование"
	}

	return "Дешифрование"
}

// изменение прогресса и проверка "готово ли"
func (w *Window) progress(pbar *widgets.QProgressDialog, fileLen int, done chan bool, ticker *time.Ticker) {
	for {
		select {
		case <-done:
			if !pbar.WasCanceled() {
				pbar.SetValue(pbar.Maximum() - 1)
				time.Sleep(time.Millisecond * 10)
				pbar.SetValue(pbar.Maximum())
				pbar.Close()
			}

			return
		case <-ticker.C:
			if !pbar.WasCanceled() {
				pbar.SetValue(pbar.Value() + fileLen/50)
			}
		}
	}
}

func (w *Window) operationFile(encrypt bool) {
	// читаем файлик
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

	// флаг для лога при остановке прогрессбара
	isProgress := true

	// тикер и канал для работы прогресс бара
	ticker := time.NewTicker(getTickerTime(dataLen, w.selectCipher()))
	done := make(chan bool)

	pbar := widgets.NewQProgressDialog2(fmt.Sprintf("%s...", operationType(encrypt)), "Отмена", 0,
		dataLen, w.uiWindow.Centralwidget, core.Qt__Widget)
	pbar.ConnectCanceled(func() {
		if isProgress {
			w.addLog(fmt.Sprintf("%s остановлено", operationType(encrypt)))
		}
	})

	// в горутине занимаемся отображением прогресс бара
	go w.progress(pbar, dataLen, done, ticker)

	encrypted, err := w.operation(data, encrypt)
	if err != nil {
		w.addErrLog(err)
		return
	}

	// смотрим, вдруг пользователь нажал отмену
	if pbar.WasCanceled() {
		return
	}

	// иначе завершаем прогресс бар
	isProgress = false
	ticker.Stop()
	done <- true

	// собираем выходной файл
	fileOut := w.getOutFilename(filename, encrypt)

	err = writeFile(encrypted, fileOut)
	if err != nil {
		w.addErrLog(err)
		w.addLog("Не получается записать файл")
		return
	}

	w.addLog(resultLog(encrypt, fileOut))
}

func (w *Window) EncryptFile() {
	w.operationFile(true)
}

func (w *Window) DecryptFile() {
	w.operationFile(false)
}

package window

import (
	"encoding/hex"
	"github.com/elizarpif/cryptoswitch-ui/ui"
	"github.com/elizarpif/diploma-elliptic/cryptoswitch"
)

type Window struct {
	uiWindow *ui.UICryptoswitchMainWindow

	stopCipher bool
	privKey    *cryptoswitch.PrivateKey
}

func NewWindow(ui *ui.UICryptoswitchMainWindow) *Window {
	return &Window{
		uiWindow: ui,
	}
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
	w.uiWindow.SecretKey.SetText(hex.EncodeToString(privKey.D.Bytes()))
}

func (w *Window) EncryptData() {
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

func (w *Window) DecryptData() {
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
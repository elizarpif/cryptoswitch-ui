// WARNING! All changes made in this file will be lost!
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type UICryptoswitchMainWindow struct {
	Centralwidget        *widgets.QWidget
	TabWidget            *widgets.QTabWidget
	Text                 *widgets.QWidget
	PlainText            *widgets.QTextEdit
	LabelPlainText       *widgets.QLabel
	CipherText           *widgets.QTextEdit
	LabelCipherText      *widgets.QLabel
	LabelCountPlainText  *widgets.QLabel
	LabelCountCipherText *widgets.QLabel
	File                 *widgets.QWidget
	SelectInFileBtn      *widgets.QPushButton
	LabelInFile          *widgets.QLabel
	LabelOutFile         *widgets.QLabel
	SelectOutFileBtn     *widgets.QPushButton
	DecryptBtn           *widgets.QPushButton
	EncryptBtn           *widgets.QPushButton
	CipherSwitch         *widgets.QComboBox
	LabelAlgorithm       *widgets.QLabel
	LabelCipherMode      *widgets.QLabel
	CipherMode           *widgets.QComboBox
	ParamsGroup          *widgets.QGroupBox
	EllipticCurve        *widgets.QComboBox
	LabelCurve           *widgets.QLabel
	LabelX               *widgets.QLabel
	ParamXEdit           *widgets.QLineEdit
	LabelY               *widgets.QLabel
	ParamYEdit           *widgets.QLineEdit
	GenerateBtn          *widgets.QPushButton
	SelectFileBtn3       *widgets.QPushButton
	SecretKey            *widgets.QLineEdit
	LabelSecretKey       *widgets.QLabel
	Logs                 *widgets.QTextEdit
	Menubar              *widgets.QMenuBar
	Statusbar            *widgets.QStatusBar
}

func (this *UICryptoswitchMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 721, 979))
	MainWindow.SetMinimumSize(core.NewQSize2(721, 979))
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	MainWindow.SetFont(font)
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetStyleSheet("")
	this.TabWidget = widgets.NewQTabWidget(this.Centralwidget)
	this.TabWidget.SetObjectName("TabWidget")
	this.TabWidget.SetGeometry(core.NewQRect4(50, 310, 601, 421))
	this.Text = widgets.NewQWidget(this.TabWidget, core.Qt__Widget)
	this.Text.SetObjectName("Text")
	this.PlainText = widgets.NewQTextEdit(this.Text)
	this.PlainText.SetObjectName("PlainText")
	this.PlainText.SetGeometry(core.NewQRect4(20, 50, 541, 101))
	this.LabelPlainText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelPlainText.SetObjectName("LabelPlainText")
	this.LabelPlainText.SetGeometry(core.NewQRect4(20, 20, 291, 21))
	font = gui.NewQFont()
	font.SetFamily("Zapfino")
	this.LabelPlainText.SetFont(font)
	this.CipherText = widgets.NewQTextEdit(this.Text)
	this.CipherText.SetObjectName("CipherText")
	this.CipherText.SetGeometry(core.NewQRect4(20, 220, 541, 101))
	this.CipherText.SetReadOnly(true)
	this.LabelCipherText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelCipherText.SetObjectName("LabelCipherText")
	this.LabelCipherText.SetGeometry(core.NewQRect4(20, 190, 301, 21))
	font = gui.NewQFont()
	font.SetFamily("Zapfino")
	this.LabelCipherText.SetFont(font)
	this.LabelCountPlainText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelCountPlainText.SetObjectName("LabelCountPlainText")
	this.LabelCountPlainText.SetGeometry(core.NewQRect4(20, 150, 541, 21))
	this.LabelCountCipherText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelCountCipherText.SetObjectName("LabelCountCipherText")
	this.LabelCountCipherText.SetGeometry(core.NewQRect4(20, 320, 541, 21))
	this.TabWidget.AddTab(this.Text, "")
	this.File = widgets.NewQWidget(this.TabWidget, core.Qt__Widget)
	this.File.SetObjectName("File")
	this.SelectInFileBtn = widgets.NewQPushButton(this.File)
	this.SelectInFileBtn.SetObjectName("SelectInFileBtn")
	this.SelectInFileBtn.SetGeometry(core.NewQRect4(30, 40, 221, 41))
	this.LabelInFile = widgets.NewQLabel(this.File, core.Qt__Widget)
	this.LabelInFile.SetObjectName("LabelInFile")
	this.LabelInFile.SetGeometry(core.NewQRect4(40, 80, 531, 21))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.LabelInFile.SetFont(font)
	this.LabelOutFile = widgets.NewQLabel(this.File, core.Qt__Widget)
	this.LabelOutFile.SetObjectName("LabelOutFile")
	this.LabelOutFile.SetGeometry(core.NewQRect4(40, 160, 531, 21))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.LabelOutFile.SetFont(font)
	this.SelectOutFileBtn = widgets.NewQPushButton(this.File)
	this.SelectOutFileBtn.SetObjectName("SelectOutFileBtn")
	this.SelectOutFileBtn.SetGeometry(core.NewQRect4(30, 120, 221, 41))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.SelectOutFileBtn.SetFont(font)
	this.TabWidget.AddTab(this.File, "")
	this.DecryptBtn = widgets.NewQPushButton(this.Centralwidget)
	this.DecryptBtn.SetObjectName("DecryptBtn")
	this.DecryptBtn.SetGeometry(core.NewQRect4(50, 800, 211, 41))
	this.EncryptBtn = widgets.NewQPushButton(this.Centralwidget)
	this.EncryptBtn.SetObjectName("EncryptBtn")
	this.EncryptBtn.SetGeometry(core.NewQRect4(50, 750, 211, 41))
	this.CipherSwitch = widgets.NewQComboBox(this.Centralwidget)
	this.CipherSwitch.SetObjectName("CipherSwitch")
	this.CipherSwitch.SetGeometry(core.NewQRect4(390, 50, 231, 26))
	this.CipherSwitch.AddItem("", core.NewQVariant())
	this.CipherSwitch.AddItem("", core.NewQVariant())
	this.CipherSwitch.AddItem("", core.NewQVariant())
	this.LabelAlgorithm = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelAlgorithm.SetObjectName("LabelAlgorithm")
	this.LabelAlgorithm.SetGeometry(core.NewQRect4(400, 20, 171, 21))
	font = gui.NewQFont()
	font.SetFamily("Zapfino")
	this.LabelAlgorithm.SetFont(font)
	this.LabelCipherMode = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelCipherMode.SetObjectName("LabelCipherMode")
	this.LabelCipherMode.SetGeometry(core.NewQRect4(400, 100, 171, 21))
	font = gui.NewQFont()
	font.SetFamily("Zapfino")
	this.LabelCipherMode.SetFont(font)
	this.CipherMode = widgets.NewQComboBox(this.Centralwidget)
	this.CipherMode.SetObjectName("CipherMode")
	this.CipherMode.SetGeometry(core.NewQRect4(390, 130, 231, 26))
	this.CipherMode.AddItem("", core.NewQVariant())
	this.CipherMode.AddItem("", core.NewQVariant())
	this.ParamsGroup = widgets.NewQGroupBox(this.Centralwidget)
	this.ParamsGroup.SetObjectName("ParamsGroup")
	this.ParamsGroup.SetGeometry(core.NewQRect4(60, 10, 261, 211))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	font.SetPointSize(14)
	this.ParamsGroup.SetFont(font)
	this.EllipticCurve = widgets.NewQComboBox(this.ParamsGroup)
	this.EllipticCurve.SetObjectName("EllipticCurve")
	this.EllipticCurve.SetGeometry(core.NewQRect4(60, 40, 151, 21))
	this.EllipticCurve.AddItem("", core.NewQVariant())
	this.LabelCurve = widgets.NewQLabel(this.ParamsGroup, core.Qt__Widget)
	this.LabelCurve.SetObjectName("LabelCurve")
	this.LabelCurve.SetGeometry(core.NewQRect4(10, 40, 41, 21))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.LabelCurve.SetFont(font)
	this.LabelX = widgets.NewQLabel(this.ParamsGroup, core.Qt__Widget)
	this.LabelX.SetObjectName("LabelX")
	this.LabelX.SetGeometry(core.NewQRect4(10, 100, 21, 21))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.LabelX.SetFont(font)
	this.ParamXEdit = widgets.NewQLineEdit(this.ParamsGroup)
	this.ParamXEdit.SetObjectName("ParamXEdit")
	this.ParamXEdit.SetGeometry(core.NewQRect4(60, 100, 151, 21))
	this.LabelY = widgets.NewQLabel(this.ParamsGroup, core.Qt__Widget)
	this.LabelY.SetObjectName("LabelY")
	this.LabelY.SetGeometry(core.NewQRect4(10, 160, 31, 21))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.LabelY.SetFont(font)
	this.ParamYEdit = widgets.NewQLineEdit(this.ParamsGroup)
	this.ParamYEdit.SetObjectName("ParamYEdit")
	this.ParamYEdit.SetGeometry(core.NewQRect4(60, 160, 151, 21))
	this.GenerateBtn = widgets.NewQPushButton(this.Centralwidget)
	this.GenerateBtn.SetObjectName("GenerateBtn")
	this.GenerateBtn.SetGeometry(core.NewQRect4(60, 230, 261, 41))
	this.SelectFileBtn3 = widgets.NewQPushButton(this.Centralwidget)
	this.SelectFileBtn3.SetObjectName("SelectFileBtn3")
	this.SelectFileBtn3.SetGeometry(core.NewQRect4(440, 750, 211, 41))
	font = gui.NewQFont()
	font.SetFamily("Ubuntu Mono derivative Powerline")
	this.SelectFileBtn3.SetFont(font)
	this.SecretKey = widgets.NewQLineEdit(this.Centralwidget)
	this.SecretKey.SetObjectName("SecretKey")
	this.SecretKey.SetGeometry(core.NewQRect4(390, 200, 231, 21))
	this.SecretKey.SetEchoMode(widgets.QLineEdit__Normal)
	this.SecretKey.SetReadOnly(false)
	this.LabelSecretKey = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelSecretKey.SetObjectName("LabelSecretKey")
	this.LabelSecretKey.SetGeometry(core.NewQRect4(400, 180, 171, 21))
	font = gui.NewQFont()
	font.SetFamily("Zapfino")
	this.LabelSecretKey.SetFont(font)
	this.Logs = widgets.NewQTextEdit(this.Centralwidget)
	this.Logs.SetObjectName("Logs")
	this.Logs.SetGeometry(core.NewQRect4(50, 850, 601, 71))
	this.Logs.SetReadOnly(true)
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 721, 22))
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)

	this.RetranslateUi(MainWindow)
	this.TabWidget.SetCurrentIndex(1)
}

func (this *UICryptoswitchMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "CryptoSwitch", "", -1))
	this.LabelPlainText.SetText(_translate("MainWindow", "Текст для шифрования/расшифрования", "", -1))
	this.LabelCipherText.SetText(_translate("MainWindow", "Зашифрованный/Расшифрованный текст", "", -1))
	this.LabelCountPlainText.SetText(_translate("MainWindow", "Размер введенного текста: 0 символов", "", -1))
	this.LabelCountCipherText.SetText(_translate("MainWindow", "Размер введенного текста: 0 символов", "", -1))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Text), _translate("MainWindow", "Текст", "", -1))
	this.SelectInFileBtn.SetText(_translate("MainWindow", "Выбрать входной файл", "", -1))
	this.LabelInFile.SetText(_translate("MainWindow", "Файл не выбран", "", -1))
	this.LabelOutFile.SetText(_translate("MainWindow", "Файл не выбран", "", -1))
	this.SelectOutFileBtn.SetText(_translate("MainWindow", "Выбрать выходной файл", "", -1))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.File), _translate("MainWindow", "Файл", "", -1))
	this.DecryptBtn.SetText(_translate("MainWindow", "Расшифровать", "", -1))
	this.EncryptBtn.SetText(_translate("MainWindow", "Зашифровать", "", -1))
	this.CipherSwitch.SetItemText(0, _translate("MainWindow", "AES (128-bit)", "", -1))
	this.CipherSwitch.SetItemText(1, _translate("MainWindow", "Camellia", "", -1))
	this.CipherSwitch.SetItemText(2, _translate("MainWindow", "TwoFish", "", -1))
	this.LabelAlgorithm.SetText(_translate("MainWindow", "Алгоритм", "", -1))
	this.LabelCipherMode.SetText(_translate("MainWindow", "Режим шифрования", "", -1))
	this.CipherMode.SetItemText(0, _translate("MainWindow", "GCM", "", -1))
	this.CipherMode.SetItemText(1, _translate("MainWindow", "CBC", "", -1))
	this.ParamsGroup.SetTitle(_translate("MainWindow", "Параметры кривой", "", -1))
	this.EllipticCurve.SetItemText(0, _translate("MainWindow", "secp256k1", "", -1))
	this.LabelCurve.SetText(_translate("MainWindow", "Кривая", "", -1))
	this.LabelX.SetText(_translate("MainWindow", "X", "", -1))
	this.LabelY.SetText(_translate("MainWindow", "Y", "", -1))
	this.GenerateBtn.SetText(_translate("MainWindow", "Сгенерировать", "", -1))
	this.SelectFileBtn3.SetText(_translate("MainWindow", "Отменить", "", -1))
	this.SecretKey.SetInputMask(_translate("MainWindow", "", "", -1))
	this.LabelSecretKey.SetText(_translate("MainWindow", "Секретный ключ", "", -1))
}

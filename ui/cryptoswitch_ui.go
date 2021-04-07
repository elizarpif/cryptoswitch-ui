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
	ParamsGroup          *widgets.QGroupBox
	EllipticCurve        *widgets.QComboBox
	LabelCurve           *widgets.QLabel
	LabelX               *widgets.QLabel
	ParamXEdit           *widgets.QLineEdit
	LabelY               *widgets.QLabel
	ParamYEdit           *widgets.QLineEdit
	GenerateBtn          *widgets.QPushButton
	SelectFileBtn3       *widgets.QPushButton
	Logs                 *widgets.QTextEdit
	GroupBox2            *widgets.QGroupBox
	GroupBox             *widgets.QGroupBox
	GcmRadio             *widgets.QRadioButton
	CbcRadio             *widgets.QRadioButton
	LabelCipherMode      *widgets.QLabel
	LabelAlgorithm       *widgets.QLabel
	CipherSwitch         *widgets.QComboBox
	Menubar              *widgets.QMenuBar
	Statusbar            *widgets.QStatusBar
}

func (this *UICryptoswitchMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 721, 979))
	MainWindow.SetMinimumSize(core.NewQSize2(721, 979))
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetFamily("Ubuntu")
	font.SetPointSize(14)

	MainWindow.SetFont(font)
	MainWindow.SetStyleSheet("QToolTip\n{\n border: 1px solid black;\n background-color: #ffa02f;\n padding: 1px;\n border-radius: 3px;\n opacity: 100;\n}\n\nQWidget\n{\n color: #b1b1b1;\n background-color: #323232;\n}\n\nQWidget:item:hover\n{\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #ca0619);\n color: #000000;\n}\n\nQWidget:item:selected\n{\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);\n}\n\nQMenuBar::item\n{\n background: transparent;\n}\n\nQMenuBar::item:selected\n{\n background: transparent;\n border: 1px solid #ffaa00;\n}\n\nQMenuBar::item:pressed\n{\n background: #444;\n border: 1px solid #000;\n background-color: QLinearGradient(\n x1:0, y1:0,\n x2:0, y2:1,\n stop:1 #212121,\n stop:0.4 #343434/*,\n stop:0.2 #343434,\n stop:0.1 #ffaa00*/\n );\n margin-bottom:-1px;\n padding-bottom:1px;\n}\n\nQMenu\n{\n border: 1px solid #000;\n}\n\nQMenu::item\n{\n padding: 2px 20px 2px 20px;\n}\n\nQMenu::item:selected\n{\n color: #000000;\n}\n\nQWidget:disabled\n{\n color: #404040;\n background-color: #323232;\n}\n\nQAbstractItemView\n{\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #4d4d4d, stop: 0.1 #646464, stop: 1 #5d5d5d);\n}\n\nQWidget:focus\n{\n /*border: 2px solid QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);*/\n}\n\nQLineEdit\n{\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #4d4d4d, stop: 0 #646464, stop: 1 #5d5d5d);\n padding: 1px;\n border-style: solid;\n border: 1px solid #1e1e1e;\n border-radius: 5;\n}\n\nQPushButton\n{\n color: #b1b1b1;\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #565656, stop: 0.1 #525252, stop: 0.5 #4e4e4e, stop: 0.9 #4a4a4a, stop: 1 #464646);\n border-width: 1px;\n border-color: #1e1e1e;\n border-style: solid;\n border-radius: 6;\n padding: 3px;\n padding-left: 5px;\n padding-right: 5px;\n}\n\nQPushButton:pressed\n{\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #2d2d2d, stop: 0.1 #2b2b2b, stop: 0.5 #292929, stop: 0.9 #282828, stop: 1 #252525);\n}\n\nQComboBox\n{\n selection-background-color: #ffaa00;\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #565656, stop: 0.1 #525252, stop: 0.5 #4e4e4e, stop: 0.9 #4a4a4a, stop: 1 #464646);\n border-style: solid;\n border: 1px solid #1e1e1e;\n border-radius: 5;\n}\n\nQComboBox:hover,QPushButton:hover\n{\n border: 2px solid QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);\n}\n\n\nQComboBox:on\n{\n padding-top: 3px;\n padding-left: 4px;\n background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #2d2d2d, stop: 0.1 #2b2b2b, stop: 0.5 #292929, stop: 0.9 #282828, stop: 1 #252525);\n selection-background-color: #ffaa00;\n}\n\nQComboBox QAbstractItemView\n{\n border: 2px solid darkgray;\n selection-background-color: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);\n}\n\nQComboBox::drop-down\n{\n subcontrol-origin: padding;\n subcontrol-position: top right;\n width: 15px;\n\n border-left-width: 0px;\n border-left-color: darkgray;\n border-left-style: solid; /* just a single line */\n border-top-right-radius: 3px; /* same radius as the QComboBox */\n border-bottom-right-radius: 3px;\n }\n\nQComboBox::down-arrow\n{\n image: url(:/down_arrow.png);\n}\n\nQGroupBox:focus\n{\nborder: 2px solid QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);\n}\n\nQTextEdit:focus\n{\n border: 2px solid QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);\n}\n\nQScrollBar:horizontal {\n border: 1px solid #222222;\n background: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0.0 #121212, stop: 0.2 #282828, stop: 1 #484848);\n height: 7px;\n margin: 0px 16px 0 16px;\n}\n\nQScrollBar::handle:horizontal\n{\n background: QLinearGradient( x1: 0, y1: 0, x2: 1, y2: 0, stop: 0 #ffa02f, stop: 0.5 #d7801a, stop: 1 #ffa02f);\n min-height: 20px;\n border-radius: 2px;\n}\n\nQScrollBar::add-line:horizontal {\n border: 1px solid #1b1b19;\n border-radius: 2px;\n background: QLinearGradient( x1: 0, y1: 0, x2: 1, y2: 0, stop: 0 #ffa02f, stop: 1 #d7801a);\n width: 14px;\n subcontrol-position: right;\n subcontrol-origin: margin;\n}\n\nQScrollBar::sub-line:horizontal {\n border: 1px solid #1b1b19;\n border-radius: 2px;\n background: QLinearGradient( x1: 0, y1: 0, x2: 1, y2: 0, stop: 0 #ffa02f, stop: 1 #d7801a);\n width: 14px;\n subcontrol-position: left;\n subcontrol-origin: margin;\n}\n\nQScrollBar::right-arrow:horizontal, QScrollBar::left-arrow:horizontal\n{\n border: 1px solid black;\n width: 1px;\n height: 1px;\n background: white;\n}\n\nQScrollBar::add-page:horizontal, QScrollBar::sub-page:horizontal\n{\n background: none;\n}\n\nQScrollBar:vertical\n{\n background: QLinearGradient( x1: 0, y1: 0, x2: 1, y2: 0, stop: 0.0 #121212, stop: 0.2 #282828, stop: 1 #484848);\n width: 7px;\n margin: 16px 0 16px 0;\n border: 1px solid #222222;\n}\n\nQScrollBar::handle:vertical\n{\n background: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 0.5 #d7801a, stop: 1 #ffa02f);\n min-height: 20px;\n border-radius: 2px;\n}\n\nQScrollBar::add-line:vertical\n{\n border: 1px solid #1b1b19;\n border-radius: 2px;\n background: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #ffa02f, stop: 1 #d7801a);\n height: 14px;\n subcontrol-position: bottom;\n subcontrol-origin: margin;\n}\n\nQScrollBar::sub-line:vertical\n{\n border: 1px solid #1b1b19;\n border-radius: 2px;\n background: QLinearGradient( x1: 0, y1: 0, x2: 0, y2: 1, stop: 0 #d7801a, stop: 1 #ffa02f);\n height: 14px;\n subcontrol-position: top;\n subcontrol-origin: margin;\n}\n\nQScrollBar::up-arrow:vertical, QScrollBar::down-arrow:vertical\n{\n border: 1px solid black;\n width: 1px;\n height: 1px;\n background: white;\n}\n\n\nQScrollBar::add-page:vertical, QScrollBar::sub-page:vertical\n{\n background: none;\n}\n\nQTextEdit\n{\n background-color: #242424;\n}\n\nQPlainTextEdit\n{\n background-color: #242424;\n}\n\nQHeaderView::section\n{\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:0 #616161, stop: 0.5 #505050, stop: 0.6 #434343, stop:1 #656565);\n color: white;\n padding-left: 4px;\n border: 1px solid #6c6c6c;\n}\n\nQCheckBox:disabled\n{\ncolor: #414141;\n}\n\nQDockWidget::title\n{\n text-align: center;\n spacing: 3px; /* spacing between items in the tool bar */\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:0 #323232, stop: 0.5 #242424, stop:1 #323232);\n}\n\nQDockWidget::close-button, QDockWidget::float-button\n{\n text-align: center;\n spacing: 1px; /* spacing between items in the tool bar */\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:0 #323232, stop: 0.5 #242424, stop:1 #323232);\n}\n\nQDockWidget::close-button:hover, QDockWidget::float-button:hover\n{\n background: #242424;\n}\n\nQDockWidget::close-button:pressed, QDockWidget::float-button:pressed\n{\n padding: 1px -1px -1px 1px;\n}\n\nQMainWindow::separator\n{\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:0 #161616, stop: 0.5 #151515, stop: 0.6 #212121, stop:1 #343434);\n color: white;\n padding-left: 4px;\n border: 1px solid #4c4c4c;\n spacing: 3px; /* spacing between items in the tool bar */\n}\n\nQMainWindow::separator:hover\n{\n\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:0 #d7801a, stop:0.5 #b56c17 stop:1 #ffa02f);\n color: white;\n padding-left: 4px;\n border: 1px solid #6c6c6c;\n spacing: 3px; /* spacing between items in the tool bar */\n}\n\nQToolBar::handle\n{\n spacing: 3px; /* spacing between items in the tool bar */\n background: url(:/images/handle.png);\n}\n\nQMenu::separator\n{\n height: 2px;\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:0 #161616, stop: 0.5 #151515, stop: 0.6 #212121, stop:1 #343434);\n color: white;\n padding-left: 4px;\n margin-left: 10px;\n margin-right: 5px;\n}\n\nQProgressBar\n{\n border: 2px solid grey;\n border-radius: 5px;\n text-align: center;\n}\n\nQProgressBar::chunk\n{\n background-color: #d7801a;\n width: 2.15px;\n margin: 0.5px;\n}\n\nQTabBar::tab {\n color: #b1b1b1;\n border: 1px solid #444;\n border-bottom-style: none;\n background-color: #323232;\n padding-left: 10px;\n padding-right: 10px;\n padding-top: 3px;\n padding-bottom: 2px;\n margin-right: -1px;\n}\n\nQTabWidget::pane {\n border: 1px solid #444;\n top: 1px;\n}\n\nQTabBar::tab:last\n{\n margin-right: 0; /* the last selected tab has nothing to overlap with on the right */\n border-top-right-radius: 3px;\n}\n\nQTabBar::tab:first:!selected\n{\n margin-left: 0px; /* the last selected tab has nothing to overlap with on the right */\n\n\n border-top-left-radius: 3px;\n}\n\nQTabBar::tab:!selected\n{\n color: #b1b1b1;\n border-bottom-style: solid;\n margin-top: 3px;\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:1 #212121, stop:.4 #343434);\n}\n\nQTabBar::tab:selected\n{\n border-top-left-radius: 3px;\n border-top-right-radius: 3px;\n margin-bottom: 0px;\n}\n\nQTabBar::tab:!selected:hover\n{\n /*border-top: 2px solid #ffaa00;\n padding-bottom: 3px;*/\n border-top-left-radius: 3px;\n border-top-right-radius: 3px;\n background-color: QLinearGradient(x1:0, y1:0, x2:0, y2:1, stop:1 #212121, stop:0.4 #343434, stop:0.2 #343434, stop:0.1 #ffaa00);\n}\n\nQRadioButton::indicator:checked, QRadioButton::indicator:unchecked{\n color: #b1b1b1;\n background-color: #323232;\n border: 1px solid #b1b1b1;\n border-radius: 6px;\n}\n\nQRadioButton::indicator:checked\n{\n background-color: qradialgradient(\n cx: 0.5, cy: 0.5,\n fx: 0.5, fy: 0.5,\n radius: 1.0,\n stop: 0.25 #ffaa00,\n stop: 0.3 #323232\n );\n}\n\nQCheckBox::indicator{\n color: #b1b1b1;\n background-color: #323232;\n border: 1px solid #b1b1b1;\n width: 9px;\n height: 9px;\n}\n\nQRadioButton::indicator\n{\n border-radius: 6px;\n}\n\nQRadioButton::indicator:hover, QCheckBox::indicator:hover\n{\n border: 1px solid #ffaa00;\n}\n\nQCheckBox::indicator:checked\n{\n image:url(:/images/checkbox.png);\n}\n\nQCheckBox::indicator:disabled, QRadioButton::indicator:disabled\n{\n border: 1px solid #444;\n}")
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetStyleSheet("")
	this.TabWidget = widgets.NewQTabWidget(this.Centralwidget)
	this.TabWidget.SetObjectName("TabWidget")
	this.TabWidget.SetGeometry(core.NewQRect4(50, 310, 601, 421))

	this.TabWidget.SetFont(font)
	this.Text = widgets.NewQWidget(this.TabWidget, core.Qt__Widget)
	this.Text.SetObjectName("Text")

	this.Text.SetFont(font)
	this.PlainText = widgets.NewQTextEdit(this.Text)
	this.PlainText.SetObjectName("PlainText")
	this.PlainText.SetGeometry(core.NewQRect4(20, 50, 541, 101))

	this.PlainText.SetFont(font)
	this.PlainText.SetStyleSheet("")
	this.LabelPlainText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelPlainText.SetObjectName("LabelPlainText")
	this.LabelPlainText.SetGeometry(core.NewQRect4(20, 20, 501, 21))

	this.LabelPlainText.SetFont(font)
	this.CipherText = widgets.NewQTextEdit(this.Text)
	this.CipherText.SetObjectName("CipherText")
	this.CipherText.SetGeometry(core.NewQRect4(20, 230, 541, 101))

	this.CipherText.SetFont(font)
	this.CipherText.SetReadOnly(true)
	this.LabelCipherText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelCipherText.SetObjectName("LabelCipherText")
	this.LabelCipherText.SetGeometry(core.NewQRect4(20, 190, 511, 21))

	this.LabelCipherText.SetFont(font)
	this.LabelCountPlainText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelCountPlainText.SetObjectName("LabelCountPlainText")
	this.LabelCountPlainText.SetGeometry(core.NewQRect4(20, 160, 541, 21))

	font.SetPointSize(14)
	this.LabelCountPlainText.SetFont(font)
	this.LabelCountCipherText = widgets.NewQLabel(this.Text, core.Qt__Widget)
	this.LabelCountCipherText.SetObjectName("LabelCountCipherText")
	this.LabelCountCipherText.SetGeometry(core.NewQRect4(20, 340, 541, 21))

	font.SetPointSize(14)
	this.LabelCountCipherText.SetFont(font)
	this.TabWidget.AddTab(this.Text, "")
	this.File = widgets.NewQWidget(this.TabWidget, core.Qt__Widget)
	this.File.SetObjectName("File")
	this.SelectInFileBtn = widgets.NewQPushButton(this.File)
	this.SelectInFileBtn.SetObjectName("SelectInFileBtn")
	this.SelectInFileBtn.SetGeometry(core.NewQRect4(30, 20, 281, 41))
	this.LabelInFile = widgets.NewQLabel(this.File, core.Qt__Widget)
	this.LabelInFile.SetObjectName("LabelInFile")
	this.LabelInFile.SetGeometry(core.NewQRect4(40, 80, 531, 21))
	font = gui.NewQFont()
	this.LabelInFile.SetFont(font)
	this.LabelOutFile = widgets.NewQLabel(this.File, core.Qt__Widget)
	this.LabelOutFile.SetObjectName("LabelOutFile")
	this.LabelOutFile.SetGeometry(core.NewQRect4(40, 180, 531, 21))
	font = gui.NewQFont()
	this.LabelOutFile.SetFont(font)
	this.SelectOutFileBtn = widgets.NewQPushButton(this.File)
	this.SelectOutFileBtn.SetObjectName("SelectOutFileBtn")
	this.SelectOutFileBtn.SetGeometry(core.NewQRect4(30, 120, 281, 41))
	this.TabWidget.AddTab(this.File, "")
	this.DecryptBtn = widgets.NewQPushButton(this.Centralwidget)
	this.DecryptBtn.SetObjectName("DecryptBtn")
	this.DecryptBtn.SetGeometry(core.NewQRect4(50, 800, 211, 41))

	font.SetPointSize(14)
	this.DecryptBtn.SetFont(font)
	this.EncryptBtn = widgets.NewQPushButton(this.Centralwidget)
	this.EncryptBtn.SetObjectName("EncryptBtn")
	this.EncryptBtn.SetGeometry(core.NewQRect4(50, 750, 211, 41))

	font.SetPointSize(14)
	this.EncryptBtn.SetFont(font)
	this.ParamsGroup = widgets.NewQGroupBox(this.Centralwidget)
	this.ParamsGroup.SetObjectName("ParamsGroup")
	this.ParamsGroup.SetGeometry(core.NewQRect4(50, 20, 271, 201))

	font.SetPointSize(14)
	this.ParamsGroup.SetFont(font)
	this.EllipticCurve = widgets.NewQComboBox(this.ParamsGroup)
	this.EllipticCurve.SetObjectName("EllipticCurve")
	this.EllipticCurve.SetGeometry(core.NewQRect4(100, 40, 151, 21))

	this.EllipticCurve.SetFont(font)
	this.EllipticCurve.AddItem("", core.NewQVariant())
	this.LabelCurve = widgets.NewQLabel(this.ParamsGroup, core.Qt__Widget)
	this.LabelCurve.SetObjectName("LabelCurve")
	this.LabelCurve.SetGeometry(core.NewQRect4(10, 40, 81, 21))

	this.LabelCurve.SetFont(font)
	this.LabelX = widgets.NewQLabel(this.ParamsGroup, core.Qt__Widget)
	this.LabelX.SetObjectName("LabelX")
	this.LabelX.SetGeometry(core.NewQRect4(10, 100, 21, 21))

	this.LabelX.SetFont(font)
	this.ParamXEdit = widgets.NewQLineEdit(this.ParamsGroup)
	this.ParamXEdit.SetObjectName("ParamXEdit")
	this.ParamXEdit.SetGeometry(core.NewQRect4(100, 100, 151, 21))

	this.ParamXEdit.SetFont(font)
	this.ParamXEdit.SetReadOnly(true)
	this.LabelY = widgets.NewQLabel(this.ParamsGroup, core.Qt__Widget)
	this.LabelY.SetObjectName("LabelY")
	this.LabelY.SetGeometry(core.NewQRect4(10, 160, 31, 21))

	this.LabelY.SetFont(font)
	this.ParamYEdit = widgets.NewQLineEdit(this.ParamsGroup)
	this.ParamYEdit.SetObjectName("ParamYEdit")
	this.ParamYEdit.SetGeometry(core.NewQRect4(100, 160, 151, 21))

	this.ParamYEdit.SetFont(font)
	this.ParamYEdit.SetReadOnly(true)
	this.GenerateBtn = widgets.NewQPushButton(this.Centralwidget)
	this.GenerateBtn.SetObjectName("GenerateBtn")
	this.GenerateBtn.SetGeometry(core.NewQRect4(50, 240, 271, 41))

	this.GenerateBtn.SetFont(font)
	this.SelectFileBtn3 = widgets.NewQPushButton(this.Centralwidget)
	this.SelectFileBtn3.SetObjectName("SelectFileBtn3")
	this.SelectFileBtn3.SetGeometry(core.NewQRect4(440, 750, 211, 41))

	font.SetPointSize(14)
	this.SelectFileBtn3.SetFont(font)
	this.Logs = widgets.NewQTextEdit(this.Centralwidget)
	this.Logs.SetObjectName("Logs")
	this.Logs.SetGeometry(core.NewQRect4(50, 850, 601, 71))

	this.Logs.SetFont(font)
	this.Logs.SetReadOnly(true)
	this.GroupBox2 = widgets.NewQGroupBox(this.Centralwidget)
	this.GroupBox2.SetObjectName("GroupBox2")
	this.GroupBox2.SetGeometry(core.NewQRect4(370, 20, 281, 261))

	this.GroupBox2.SetFont(font)
	this.GroupBox = widgets.NewQGroupBox(this.GroupBox2)
	this.GroupBox.SetObjectName("GroupBox")
	this.GroupBox.SetGeometry(core.NewQRect4(10, 130, 241, 111))

	this.GroupBox.SetFont(font)
	this.GcmRadio = widgets.NewQRadioButton(this.GroupBox)
	this.GcmRadio.SetObjectName("GcmRadio")
	this.GcmRadio.SetGeometry(core.NewQRect4(10, 30, 100, 20))

	font.SetPointSize(14)
	this.GcmRadio.SetFont(font)
	this.CbcRadio = widgets.NewQRadioButton(this.GroupBox)
	this.CbcRadio.SetObjectName("CbcRadio")
	this.CbcRadio.SetGeometry(core.NewQRect4(10, 80, 100, 20))

	font.SetPointSize(14)
	this.CbcRadio.SetFont(font)
	this.CbcRadio.SetChecked(true)
	this.LabelCipherMode = widgets.NewQLabel(this.GroupBox2, core.Qt__Widget)
	this.LabelCipherMode.SetObjectName("LabelCipherMode")
	this.LabelCipherMode.SetGeometry(core.NewQRect4(10, 111, 231, 21))

	font.SetPointSize(14)
	this.LabelCipherMode.SetFont(font)
	this.LabelAlgorithm = widgets.NewQLabel(this.GroupBox2, core.Qt__Widget)
	this.LabelAlgorithm.SetObjectName("LabelAlgorithm")
	this.LabelAlgorithm.SetGeometry(core.NewQRect4(10, 40, 171, 21))

	font.SetPointSize(14)
	this.LabelAlgorithm.SetFont(font)
	this.CipherSwitch = widgets.NewQComboBox(this.GroupBox2)
	this.CipherSwitch.SetObjectName("CipherSwitch")
	this.CipherSwitch.SetGeometry(core.NewQRect4(10, 70, 241, 26))

	this.CipherSwitch.SetFont(font)
	this.CipherSwitch.AddItem("", core.NewQVariant())
	this.CipherSwitch.AddItem("", core.NewQVariant())
	this.CipherSwitch.AddItem("", core.NewQVariant())
	this.CipherSwitch.AddItem("", core.NewQVariant())
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 721, 28))
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)

	this.RetranslateUi(MainWindow)
	this.TabWidget.SetCurrentIndex(0)
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
	this.ParamsGroup.SetTitle(_translate("MainWindow", "Параметры кривой", "", -1))
	this.EllipticCurve.SetItemText(0, _translate("MainWindow", "secp256k1", "", -1))
	this.LabelCurve.SetText(_translate("MainWindow", "Кривая", "", -1))
	this.LabelX.SetText(_translate("MainWindow", "X", "", -1))
	this.LabelY.SetText(_translate("MainWindow", "Y", "", -1))
	this.GenerateBtn.SetText(_translate("MainWindow", "Сгенерировать", "", -1))
	this.SelectFileBtn3.SetText(_translate("MainWindow", "Отменить", "", -1))
	this.GroupBox2.SetTitle(_translate("MainWindow", "", "", -1))
	this.GroupBox.SetTitle(_translate("MainWindow", "", "", -1))
	this.GcmRadio.SetText(_translate("MainWindow", "GCM", "", -1))
	this.CbcRadio.SetText(_translate("MainWindow", "CBC", "", -1))
	this.LabelCipherMode.SetText(_translate("MainWindow", "Режим шифрования", "", -1))
	this.LabelAlgorithm.SetText(_translate("MainWindow", "Алгоритм", "", -1))
	this.CipherSwitch.SetItemText(0, _translate("MainWindow", "AES (128-bit)", "", -1))
	this.CipherSwitch.SetItemText(1, _translate("MainWindow", "DES", "", -1))
	this.CipherSwitch.SetItemText(2, _translate("MainWindow", "Camellia", "", -1))
	this.CipherSwitch.SetItemText(3, _translate("MainWindow", "TwoFish", "", -1))
}

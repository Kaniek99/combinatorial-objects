package windows

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	Width float32
	Hight float32

	// darktheme bool
	// Application fyne.App
	// Window      *fyne.Window
}

func (mainWindow *MainWindow) Run() {
	app := app.New()
	// app.Settings().SetTheme(theme.DarkTheme())

	headertxt := canvas.NewText("Generating Combinatorial Objects", color.Black)
	header := container.New(layout.NewCenterLayout(), headertxt)

	button1 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button2 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button3 := widget.NewButton("Generate combinations of n-set", func() { log.Println("not implemented yet") })
	button4 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button5 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button6 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })

	menu := container.NewGridWithRows(7, header, button1, button2, button3, button4, button5, button6)

	// output.Move(fyne.NewPos(10, 10))

	window := app.NewWindow("Generating Combinatorial Objects")

	window.Resize(fyne.NewSize(mainWindow.Width, mainWindow.Hight))
	window.SetContent(menu)
	window.ShowAndRun()
}

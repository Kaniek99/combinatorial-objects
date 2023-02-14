package windows

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type WindowConfig struct {
	// Width float32
	// Hight float32

	// Application *fyne.App
	// Window      *fyne.Window
}

func (window WindowConfig) Run() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Hello World")

	top := canvas.NewText("top bar", color.Black)
	left := canvas.NewText("left", color.Black)
	middle := canvas.NewText("content", color.Black)
	content := container.NewBorder(top, nil, left, nil, middle)
	mainWindow.Resize(fyne.NewSize(1280, 720))
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

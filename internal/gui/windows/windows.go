package windows

import "fyne.io/fyne/v2"

// "image/color"

// "fyne.io/fyne/v2"
// "fyne.io/fyne/v2/app"
// "fyne.io/fyne/v2/canvas"
// "fyne.io/fyne/v2/container"

type WindowConfig struct {
	Width float32
	Hight float32

	Application fyne.App
	// Window      *fyne.Window
}

func (wc *WindowConfig) Run(name string, content fyne.CanvasObject) {
	window := wc.Application.NewWindow(name)
	window.Resize(fyne.NewSize(wc.Width, wc.Hight))
	window.SetContent(content)
	window.ShowAndRun()
}

// func (window WindowConfig) Run() {
// 	app := app.New()

// 	main := WindowConfig{1280, 720, &app}
// 	mainWindow := app.NewWindow("Hello World")

// 	top := canvas.NewText("Generating Combinatorial Objects", color.Black)
//  left := canvas.NewText("left", color.Black)
// 	buttons := container.NewWithoutLayout()
// 	middle := canvas.NewText("content", color.Black)
// 	content := container.NewBorder(top, nil, buttons, nil, middle)
// 	mainWindow.Resize(fyne.NewSize(main.Width, main.Hight))
// 	mainWindow.SetContent(content)
// 	mainWindow.ShowAndRun()
// }

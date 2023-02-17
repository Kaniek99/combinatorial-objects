package main

import (
	"image/color"
	"kaniek99/combinatorial-objects/internal/gui/windows"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()
	main := windows.WindowConfig{1280, 720, app}

	top := canvas.NewText("Generating Combinatorial Objects", color.Black)
	buttons := container.NewWithoutLayout()
	middle := canvas.NewText("content", color.Black)
	content := container.NewBorder(top, nil, buttons, nil, middle)

	main.Run("Generating Combinatorial Objects", content)

	// mainWindow := windows.WindowConfig{}
	// mainWindow.Run()
}

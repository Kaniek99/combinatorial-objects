package main

import (
	"kaniek99/combinatorial-objects/internal/gui/windows"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	main := windows.MainWindow{Width: 1120, Height: 630, Application: &app}
	main.Run()
}

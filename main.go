package main

import (
	"fyne.io/fyne/v2/app"

	"kaniek99/combinatorial-objects/internal/gui/windows"
)

func main() {
	app := app.New()
	main := windows.MainWindow{Width: 320, Hight: 720, Application: &app}
	main.Run()
}

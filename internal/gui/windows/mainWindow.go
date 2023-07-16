package windows

import (
	"fmt"
	"image/color"
	"kaniek99/combinatorial-objects/internal/logic/set"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	Width float32
	Hight float32

	Application *fyne.App
}

func (mainWindow *MainWindow) Run() {
	headertxt := canvas.NewText("Generating Combinatorial Objects", color.Black)
	header := container.New(layout.NewCenterLayout(), headertxt)

	button1 := widget.NewButton("Generate all permutations on n-set", func() { mainWindow.PermutationsButton() })
	button2 := widget.NewButton("Generate permutation from inversion sequence", func() { mainWindow.PermutationFromInversionSequenceButton() })
	button3 := widget.NewButton("Generate all combinations of n-set", func() { mainWindow.CombinationsButton() })

	menu := container.NewGridWithRows(4, header, button1, button2, button3)

	window := (*mainWindow.Application).NewWindow("Generating Combinatorial Objects")

	window.Resize(fyne.NewSize(mainWindow.Width, mainWindow.Hight))
	window.SetContent(menu)
	window.SetMaster()
	window.ShowAndRun()
}

func (window *MainWindow) RunErrorWindow(errormessage string) {
	errorWindow := (*window.Application).NewWindow("Error")
	content := container.NewVBox(canvas.NewText(errormessage, color.Black), widget.NewButton("Ok", func() { errorWindow.Close() }))
	errorWindow.Resize(fyne.NewSize(640, 100))
	errorWindow.SetContent(content)
	errorWindow.Show()
}

func (window *MainWindow) CombinationsButton() {
	entryWindow := (*window.Application).NewWindow("EntryWidget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Insert elements of the set here. Numbers should be separated with a comma and a space e.g. 3, 2, 1, 0")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		set, err := set.GenerateSet(input.Text)
		if err != nil {
			window.RunErrorWindow(fmt.Sprintf("%v", err)) // why did I do it this way instead of passing error? Fix it in the future
			return
		}
		set.GenerateCombinations()
		fmt.Println("Combinations of inserted set: ")
		for _, combination := range set.Combinations {
			fmt.Println(combination)
		}
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

func (window *MainWindow) PermutationFromInversionSequenceButton() {
	entryWindow := (*window.Application).NewWindow("EntryWidget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Insert elements of the set here. Numbers should be separated with a comma and a space e.g. 3, 2, 1, 0")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		invSeq, err := set.CreateInversionSequence(input.Text)
		if err != nil {
			window.RunErrorWindow(fmt.Sprintf("%v", err))
			return
		}
		perm := invSeq.GeneratePermutation()
		fmt.Printf("The permutation generated from ("+input.Text+") is "+"%v\n", perm)
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

func (window *MainWindow) PermutationsButton() {
	entryWindow := (*window.Application).NewWindow("EntryWidget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Insert elements of the set here. Numbers should be separated with a coma and a space e.g. 3, 2, 1, 0")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		set, err := set.GenerateSet(input.Text)
		if err != nil {
			window.RunErrorWindow(fmt.Sprintf("%v", err))
			return
		}
		set.GenerateAllPermutations()
		for _, permutation := range set.Permutations {
			perm := []int{}
			for _, elem := range permutation {
				perm = append(perm, elem.Number)
			}
			fmt.Println(perm)
		}
		fmt.Printf("%v permutations generated\n", len(set.Permutations))
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

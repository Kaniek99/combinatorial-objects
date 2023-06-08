package windows

import (
	"fmt"
	"image/color"
	"kaniek99/combinatorial-objects/internal/logic/combinations"
	"kaniek99/combinatorial-objects/internal/logic/permutations"
	"log"

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

	button1 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button2 := widget.NewButton("Generate permutation with inversion sequence", func() { mainWindow.PermutationFromInversionSequenceButton() })
	button3 := widget.NewButton("Generate combinations of n-set", func() { mainWindow.CombinationsButton() })
	button4 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button5 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })
	button6 := widget.NewButton("not implemented yet", func() { log.Println("not implemented yet") })

	menu := container.NewGridWithRows(7, header, button1, button2, button3, button4, button5, button6)

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
	input.SetPlaceHolder("Enter cardinality of set")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		Set, err := combinations.GenerateSet(input.Text)
		if err != nil {
			window.RunErrorWindow(fmt.Sprintf("%v", err)) // why did I do it this way instead of passing error? Fix it in the future
			return
		}
		Set.GenerateSubsets()
		usedSet := Set.GetSet()
		combinations := Set.GetCombinations()
		fmt.Println("Combinations of: " + usedSet)
		fmt.Println(combinations)
		fmt.Println(len(Set.Subsets))
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

func (window *MainWindow) PermutationFromInversionSequenceButton() {
	entryWindow := (*window.Application).NewWindow("EntryWidget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter inversion sequence here. Numbers should be separated with a coma and a space e.g. 3, 2, 1, 0")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		invSeq, err := permutations.CreateInversionSequence(input.Text)
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

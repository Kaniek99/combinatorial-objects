// FIXME: HUGE PERFORMANCE DROP
// For generating all permutations of a set of at least 8 elements there is huge
// performance drop. Waiting for all permutations of {1, 2, 3, 4, 5, 6, 7, 8, 9}
// takes more than one minute a few seconds.
// Same problem for generating combinations of set of at least 13 elements.
// I'm not sure if I'm even able to fix it somehow. Printing 3628800 elems takes
// too long (more than 10 minutes). This value is equal to the value of all
// permutations of the ten-element set.

package windows

import (
	"fmt"
	"image/color"
	"kaniek99/combinatorial-objects/internal/logic/set"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	fyne.Window
	Width         float32
	Height        float32
	DisplayedText *widget.Label
	Application   *fyne.App
}

func (window *MainWindow) Run() {
	button1 := widget.NewButton("Generate all permutations on n-set", func() { window.PermutationsButton() })
	button2 := widget.NewButton("Generate permutation from inversion sequence", func() { window.PermutationFromInversionSequenceButton() })
	button3 := widget.NewButton("Generate all combinations of n-set", func() { window.CombinationsButton() })

	menu := container.NewVBox(button1, button2, button3)
	window.DisplayedText = widget.NewLabel("Hello World!")
	scrollContainer := container.NewScroll(window.DisplayedText)
	split := container.NewHSplit(menu, scrollContainer)
	split.SetOffset(0.2)

	window.Window = (*window.Application).NewWindow("Main Window")
	window.Window.Resize(fyne.NewSize(window.Width, window.Height))
	window.Window.SetContent(split)
	window.Window.SetMaster()
	window.Window.ShowAndRun()
}

func (window *MainWindow) RunErrorWindow(err error) {
	errorWindow := (*window.Application).NewWindow("Error")
	errorMessage := fmt.Sprintf("%v", err)
	content := container.NewVBox(canvas.NewText(errorMessage, color.Black), widget.NewButton("Ok", func() { errorWindow.Close() }))
	errorWindow.Resize(fyne.NewSize(640, 100))
	errorWindow.SetContent(content)
	errorWindow.Show()
}

func (window *MainWindow) ShowOutputWindow(windowName, data string, size fyne.Size) {
	outputWindow := (*window.Application).NewWindow(windowName)
	outputWindow.Resize(size)

	label := widget.NewLabel(data)
	label.Wrapping = fyne.TextWrapWord
	label.Resize(fyne.NewSize(300, 300))
	scrollContainer := container.NewScroll(label)
	outputWindow.SetContent(scrollContainer)
	outputWindow.Show()
}

// FIXME: Problem of performance drop reappeared, look comments lines 1-8
func (window *MainWindow) CombinationsButton() {
	entryWindow := (*window.Application).NewWindow("EntryWidget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Insert elements of the set here. Numbers should be separated with a comma and a space e.g. 3, 2, 1, 0")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		set, err := set.GenerateSet(input.Text)
		if err != nil {
			window.RunErrorWindow(err)
			return
		}
		set.GenerateCombinations()
		combinations := "Combinations of inserted set:\n"
		for _, combination := range set.Combinations {
			combinations += fmt.Sprintln(combination)
		}
		// window.ShowOutputWindow("All combinations of inserted set", combinations, fyne.NewSize(400, 300))
		window.DisplayOutput(combinations)
		// fmt.Println(combinations)
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
			window.RunErrorWindow(err)
			return
		}
		perm := invSeq.GeneratePermutation()
		combination := fmt.Sprintf("The permutation generated from ("+input.Text+") is: "+"%v\n", perm)
		// window.ShowOutputWindow("Combination genereted from inversion sequence", combination, fyne.NewSize(600, 100))
		window.DisplayOutput(combination)
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

// FIXME: Problem of performance drop reappeared, look comments lines 1-8
func (window *MainWindow) PermutationsButton() {
	entryWindow := (*window.Application).NewWindow("EntryWidget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Insert elements of the set here. Numbers should be separated with a coma and a space e.g. 3, 2, 1, 0")

	content := container.NewVBox(input, widget.NewButton("Confirm", func() {
		set, err := set.GenerateSet(input.Text)
		if err != nil {
			window.RunErrorWindow(err)
			return
		}
		set.GenerateAllPermutations()
		permutations := ""
		for _, permutation := range set.Permutations {
			perm := []int{}
			for _, elem := range permutation {
				perm = append(perm, elem.Number)
			}
			// fmt.Println(perm) // this line is for performance comparison
			permutations += fmt.Sprintln(perm)
		}
		// window.ShowOutputWindow("All permutations of inserted set", permutations, fyne.NewSize(400, 300))
		window.DisplayOutput(permutations)
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

func (window *MainWindow) DisplayOutput(output string) {
	window.DisplayedText.SetText(output)
}

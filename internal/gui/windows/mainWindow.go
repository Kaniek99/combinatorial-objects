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
	button2 := widget.NewButton("Generate permutation with inversion sequence", func() { mainWindow.PermutationFromInversionSequenceButton() })
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
		Set, err := set.GenerateSet(input.Text)
		if err != nil {
			window.RunErrorWindow(fmt.Sprintf("%v", err)) // why did I do it this way instead of passing error? Fix it in the future
			return
		}
		Set.GenerateCombinations()
		fmt.Println("Combinations of inserted set: ")
		fmt.Println(Set.Combinations)
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

		// set := permutations.NumericSet{Elems: []int{1, 2, 3}}
		// set.CreateFirstDirectedPermutation()
		// fmt.Println(len(set.Permutations))
		// perm := set.Permutations[0] // 1, 2, 3
		// index, err := perm.FindIndexOfGreatestMobileNumber()
		// if err != nil {
		// 	fmt.Println("err")
		// }
		// fmt.Println(index) // 2
		// perm = perm.Swap(index, perm[index]) // 1, 3, 2
		// set.Permutations = append(set.Permutations, perm)
		// fmt.Println(len(set.Permutations))
		// index, err = perm.FindIndexOfGreatestMobileNumber()
		// if err != nil {
		// 	fmt.Println("err")
		// }
		// fmt.Println(index) // 1
		// perm = perm.Swap(index, perm[index]) // 3, 1, 2
		// set.Permutations = append(set.Permutations, perm)
		// fmt.Println(len(set.Permutations))
		// index, err = perm.FindIndexOfGreatestMobileNumber()
		// if err != nil {
		// 	fmt.Println("err")
		// }
		// fmt.Println(index) // 2
		// fmt.Println(perm[0].Direction)
		// perm = perm.Swap(index, perm[index]) // 3, 2, 1
		// set.Permutations = append(set.Permutations, perm)
		// fmt.Println(len(set.Permutations))
		// index, err = perm.FindIndexOfGreatestMobileNumber()
		// if err != nil {
		// 	fmt.Println("err")
		// }
		// fmt.Println(index)
		// perm = perm.Swap(index, perm[index]) // 2, 3, 1
		// set.Permutations = append(set.Permutations, perm)
		// fmt.Println(len(set.Permutations))
		// index, err = perm.FindIndexOfGreatestMobileNumber()
		// if err != nil {
		// 	fmt.Println("err")
		// }
		// fmt.Println(index)
		// perm = perm.Swap(index, perm[index]) // 2, 1, 3
		// set.Permutations = append(set.Permutations, perm)
		// fmt.Println(len(set.Permutations))
		// index, err = perm.FindIndexOfGreatestMobileNumber()
		// if err != nil {
		// 	fmt.Println("err")
		// }
		// fmt.Println(index)

		fmt.Println("test")
		set.GenerateAllPermutations()
		fmt.Println("test2")
		// fmt.Println(set.GetPermutationsAsString())
		fmt.Println(set.Permutations)
		fmt.Println(len(set.Permutations))
	}))
	entryWindow.Resize(fyne.NewSize(640, 100))
	entryWindow.SetContent(content)
	entryWindow.Show()
}

package main

import (
	"tic-tac-toe/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("tic-tac-toe")

	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.CenterOnScreen()

	scenes.NewMenuScene(myWindow)
	myWindow.ShowAndRun()
}

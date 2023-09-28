package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MenuScene struct {
	window fyne.Window
}

func NewMenuScene(fyneWindow fyne.Window) *MenuScene {
	scene := &MenuScene{window: fyneWindow}
	scene.Render()
	return scene
}

func (s*MenuScene) Render() {
		
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.png"))
	background.Resize(fyne.NewSize(500,500))
	background.Move(fyne.NewPos(0,0))

	tit := canvas.NewImageFromURI(storage.NewFileURI("./assets/tic-tac-toe.png"))
	tit.Resize(fyne.NewSize(200,40))
	tit.Move(fyne.NewPos(140,50))
	
	btnStartGame := widget.NewButton("1 Jugador", s.StartGameOnePlayer)
	btnStartGame.Resize(fyne.NewSize(130,30))
	btnStartGame.Move(fyne.NewPos(175,210))

	btnStartGameTwo := widget.NewButton("2 Jugadores", s.StartGame)
	btnStartGameTwo.Resize(fyne.NewSize(130,30))
	btnStartGameTwo.Move(fyne.NewPos(175,250))

	
	s.window.SetContent(container.NewWithoutLayout(background ,tit, btnStartGame, btnStartGameTwo))
}

func (s *MenuScene) StartGame() {
	NewGameScene(s.window)
}

func (s *MenuScene) StartGameOnePlayer() {
	NewGameOnePlayerScene(s.window)
}




package scenes

import (
	
	"image/color"
	"math/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOnePlayerScene struct {
	window fyne.Window
}

var currentPlayerGameOnePlayer string = "X"

func NewGameOnePlayerScene(window fyne.Window) *GameOnePlayerScene {
	sceneTwo := &GameOnePlayerScene{window: window}
	sceneTwo.Render()
	return sceneTwo
}

func (s *GameOnePlayerScene) Render() {

    board := [3][3]string{
        {"", "", ""},
        {"", "", ""},
        {"", "", ""},
    }

    statusLabel := canvas.NewText("", color.Black)

    var cellButtons [3][3]*widget.Button

    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            cellButtons[x][y] = widget.NewButton("", nil)
        }
    }

	cellClicked := func(x, y int) func() {
		return func() {
			// Verifica si la celda ya está ocupada
			if board[x][y] != "" {
				return
			}
			// Verifica que el botón no sea nulo antes de establecer su texto
			if cellButtons[x][y] != nil {
				// Actualiza la celda con la marca del jugador actual (que siempre será "X")
				board[x][y] = "X"
				cellButtons[x][y].SetText("X")
	
				// Verifica si hay un ganador o empate después del movimiento del jugador "X"
				if checkWinOnePlayer(board, "X") {
					statusLabel.Text = "¡Jugador X gana!"
					disableAllCells(cellButtons)
				} else if checkDrawOnePlayer(board) {
					statusLabel.Text = "¡Empate!"
					disableAllCells(cellButtons)
				} else {
					// Turno de la máquina
					playComputerTurn(board, cellButtons)
	
					// Verifica si hay un ganador o empate después del movimiento de la máquina
					if checkWinOnePlayer(board, "O") {
						statusLabel.Text = "¡Jugador O gana!"
						disableAllCells(cellButtons)
					} else if checkDrawOnePlayer(board) {
						statusLabel.Text = "¡Empate!"
						disableAllCells(cellButtons)
					}
				}
			}
		}
	}
	
    grid := container.NewGridWithRows(3)
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            cellButton := widget.NewButton("", cellClicked(x, y))
            cellButtons[x][y] = cellButton
            grid.Add(cellButton)
        }
    }

    resetButton := widget.NewButton("Reiniciar", func() {
        for x := 0; x < 3; x++ {
            for y := 0; y < 3; y++ {
                board[x][y] = ""
                if cellButtons[x][y] != nil {
                    cellButtons[x][y].SetText("") 
                }
            }
        }
        statusLabel.Text = ""
        enableAllCells(cellButtons) 
        currentPlayerGameOnePlayer = "X"
    })

    btnSalir := widget.NewButton("Salir", func() {
        NewMenuScene(s.window)
    })

    content := container.NewVBox(
        grid,
        statusLabel,
        resetButton,
        btnSalir,
    )

	ImageSet := loadImagesConcurrently()

    background := canvas.NewImageFromURI(storage.NewFileURI("./assets/GameBackground.png"))
	background.Resize(fyne.NewSize(500,500))
	background.Move(fyne.NewPos(0,0))

    header := canvas.NewImageFromURI(storage.NewFileURI("./assets/tic-tac-toe.png"))
	header.Resize(fyne.NewSize(200,30))
	header.Move(fyne.NewPos(150,20))

	content.Resize(fyne.NewSize(300,500))
	content.Move(fyne.NewPos(100,190))

    text := widget.NewLabel("Tiempo Jugando: ")
    text.Move(fyne.NewPos(155,430))

	timerLabel := widget.NewLabel("00:00:00")
    timerLabel.Resize(fyne.NewSize(400, 400))
    timerLabel.Move(fyne.NewPos(275, 430))

    s.window.SetContent(container.NewWithoutLayout(background, header, content, timerLabel, text, ImageSet.OImage, ImageSet.XImage))

	go func() {
        for {
            if currentPlayerGameOnePlayer == "X" {
                // Turno de la máquina
                playComputerTurn(board, cellButtons)
    
                // Verifica si hay un ganador o empate después del movimiento de la máquina
                if checkWinOnePlayer(board, "O") {
                    statusLabel.Text = "¡Jugador O gana!"
                    disableAllCells(cellButtons)
                } else if checkDrawOnePlayer(board) {
                    statusLabel.Text = "¡Empate!"
                    disableAllCells(cellButtons)
                }
                currentPlayerGameOnePlayer = "" 
            }
        }
    }()

    go startTimer(timerLabel)
}


func playComputerTurn(board [3][3]string, cellButtons [3][3]*widget.Button) {
    for {
        x := rand.Intn(3)
        y := rand.Intn(3)
        if board[x][y] == "" {
            board[x][y] = "O"
            cellButtons[x][y].SetText("O")
            return
        }
    }
}


func checkWinOnePlayer(board [3][3]string, player string) bool {
    // Verificar filas
    for i := 0; i < 3; i++ {
        if board[i][0] == player && board[i][1] == player && board[i][2] == player {
            return true
        }
    }
    // Verificar columnas
    for i := 0; i < 3; i++ {
        if board[0][i] == player && board[1][i] == player && board[2][i] == player {
            return true
        }
    }
    // Verificar diagonales
    if board[0][0] == player && board[1][1] == player && board[2][2] == player {
        return true
    }
    if board[0][2] == player && board[1][1] == player && board[2][0] == player {
        return true
    }
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                return false
            }
        }
    }
    return true
}

func checkDrawOnePlayer(board [3][3]string) bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                return false
            }
        }
    }
    return true
}
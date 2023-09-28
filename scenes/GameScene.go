package scenes

import (
	"fmt"
	"image/color"
	"tic-tac-toe/models"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameScene struct {
	window fyne.Window
}

var currentPlayer string = "X" 

func NewGameScene(window fyne.Window) *GameScene{
	scene := &GameScene{window: window}
	scene.RenderTwo()
	return scene
}

func (s *GameScene) RenderTwo() {
	
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
				// Actualiza la celda con la marca del jugador actual
				board[x][y] = currentPlayer
				cellButtons[x][y].SetText(currentPlayer)
				// Cambia al siguiente jugador antes de verificar si hay un ganador
				if currentPlayer == "X" {
					currentPlayer = "O"
				} else {
					currentPlayer = "X"
				}
	
				// Verifica si hay un ganador o si el juego termina en empate
				if checkWin(board, x, y) {
					statusLabel.Text = fmt.Sprintf("¡Jugador %s gana!", board[x][y])
					disableAllCells(cellButtons)
				} else if checkDraw(board) {
					statusLabel.Text = "¡Empate!"
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
                    cellButtons[x][y].SetText("") // Restablece el texto de las celdas
                }
            }
        }
        statusLabel.Text = ""
        enableAllCells(cellButtons) // Habilita todas las celdas nuevamente
        currentPlayer = "X"
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

	go startTimer(timerLabel)

}

func loadImagesConcurrently() *models.ImageSet {
    var ImageSet *models.ImageSet
    done := make(chan struct{})

    go func() {
        ImageSet = models.NewImageSet()
        close(done)
    }()

    // Puedes realizar otras operaciones aquí mientras se cargan las imágenes en segundo plano

    <-done
    return ImageSet
}

var startTime time.Time

func startTimer(label *widget.Label) {
    startTime = time.Now()
    for {
        elapsed := time.Since(startTime)
        hours := int(elapsed.Hours())
        minutes := int(elapsed.Minutes()) % 60
        seconds := int(elapsed.Seconds()) % 60
        label.SetText(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
        time.Sleep(100 * time.Millisecond)
    }
}

func checkWin(board [3][3]string, x, y int) bool {
    player := board[x][y]

    // Verificar fila
    if board[x][0] == player && board[x][1] == player && board[x][2] == player {
        return true
    }

    // Verificar columna
    if board[0][y] == player && board[1][y] == player && board[2][y] == player {
        return true
    }

    // Verificar diagonal principal
    if x == y && board[0][0] == player && board[1][1] == player && board[2][2] == player {
        return true
    }

    // Verificar diagonal secundaria
    if x+y == 2 && board[0][2] == player && board[1][1] == player && board[2][0] == player {
        return true
    }

    return false
}

func checkDraw(board [3][3]string) bool {
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            if board[x][y] == "" {
                return false
            }
        }
    }
    return true
}

func disableAllCells(cellButtons [3][3]*widget.Button) {
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            if cellButtons[x][y] != nil {
                cellButtons[x][y].Disable()
            }
        }
    }
}

func enableAllCells(cellButtons [3][3]*widget.Button) {
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            if cellButtons[x][y] != nil {
                cellButtons[x][y].Enable()
            }
        }
    }
}


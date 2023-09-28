// models/image.go
package models

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
)

// ImageSet es una estructura que contiene las imágenes X y O.
type ImageSet struct {
    XImage *canvas.Image
    OImage *canvas.Image
}

// NewImageSet crea y devuelve una instancia de ImageSet con las imágenes X y O configuradas.
func NewImageSet() *ImageSet {
    xImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/X.png"))
    xImage.Resize(fyne.NewSize(30, 30))
    xImage.Move(fyne.NewPos(170, 100))

    oImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/O.png"))
    oImage.Resize(fyne.NewSize(34, 34))
    oImage.Move(fyne.NewPos(300, 100))

    return &ImageSet{
        XImage: xImage,
        OImage: oImage,
    }
}


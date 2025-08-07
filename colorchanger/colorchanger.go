package main

import (
	"image/color"
	"log"

	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	myApp := app.New()

	myWindow := myApp.NewWindow("ColorChanger")

	colorBox := canvas.NewRectangle(color.Black)

	colorBox.SetMinSize(fyne.NewSize(150, 150))

	topText := canvas.NewText("Click the Button to change the color!", color.Black)
	topText.Alignment = fyne.TextAlignTrailing
	topText.TextStyle = fyne.TextStyle{Italic: true}

	changeColorButton := widget.NewButton("Randomize Color", func() {

		r_val := uint8(rand.Intn(256))
		g_val := uint8(rand.Intn(256))
		b_val := uint8(rand.Intn(256))
		a_val := uint8(rand.Intn(256))

		newColor := color.NRGBA{R: r_val, G: g_val, B: b_val, A: a_val}
		colorBox.FillColor = newColor
		colorBox.Refresh()
	})

	userInput := widget.NewEntry()
	userInput.SetPlaceHolder("What color do you want to see?")
	userInputBox := widget.NewButton("Save", func() {
		log.Println("Content was:", userInput.Text)
	})

	content := container.NewVBox(
		topText,
		colorBox,
		userInput,
		userInputBox,
		changeColorButton,
	)

	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(200, 200))

	myWindow.ShowAndRun()
}

package hangman

import (
	"time"

	"github.com/rivo/tview"
)

func Victory() {

	app := tview.NewApplication()

	hangmanGame := tview.NewTextView().
		SetText("Victoire").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	hangmanGame.SetBorder(true).
		SetTitle("Hangman Game🎉")

	go func() {
		time.Sleep(5 * time.Second)
		app.Stop()
	}()

	if err := app.SetRoot(hangmanGame, true).Run(); err != nil {
		panic(err)
	}
}

func Defaite() {
	app := tview.NewApplication()

	hangmanGame := tview.NewTextView().
		SetText("Defaite").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	hangmanGame.SetBorder(true).
		SetTitle("Hangman Game🎉")

	go func() {
		time.Sleep(5 * time.Second)
		app.Stop()
	}()

	if err := app.SetRoot(hangmanGame, true).Run(); err != nil {
		panic(err)
	}
}

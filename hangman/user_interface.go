package hangman

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Victory(H *HangManData) {

	app := tview.NewApplication()

	win := tview.NewTextView().
		SetText("\n\n\n\n\n\n\n\n.##....##..#######..##.....##....##......##.####.##....##\n..##..##..##.....##.##.....##....##..##..##..##..###...##\n...####...##.....##.##.....##....##..##..##..##..####..##\n....##....##.....##.##.....##....##..##..##..##..##.##.##\n....##....##.....##.##.....##....##..##..##..##..##..####\n....##....##.....##.##.....##....##..##..##..##..##...###\n....##.....#######...#######......###..###..####.##....##\n\n\n" + HangmanState(H) + "\n\n\n The word was : \n" + H.ToFind).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetTextColor(tcell.ColorGreen)

	win.SetBorder(true).
		SetTitle(" Victory ").
		SetTitleColor(tcell.ColorDarkGreen).
		SetBorderColor(tcell.ColorDarkGreen)

	go func() {
		time.Sleep(5 * time.Second)
		app.Stop()
	}()

	if err := app.SetRoot(win, true).Run(); err != nil {
		panic(err)
	}
}

func Defaite(H *HangManData) {
	app := tview.NewApplication()

	hangmanGame := tview.NewTextView().
		SetText("\n\n\n\n\n\n.##....##..#######..##.....##.......##........#######...######..########\n..##..##..##.....##.##.....##.......##.......##.....##.##....##....##...\n...####...##.....##.##.....##.......##.......##.....##.##..........##...\n....##....##.....##.##.....##.......##.......##.....##..######.....##...\n....##....##.....##.##.....##.......##.......##.....##.......##....##...\n....##....##.....##.##.....##.......##.......##.....##.##....##....##...\n....##.....#######...#######........########..#######...######.....##...\n\n" + HangmanState(H) + "\n\n\n The word was : \n" + H.ToFind).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetTextColor(tcell.ColorRed)
	hangmanGame.SetBorder(true).
		SetTitle(" Defeat ").
		SetTitleColor(tcell.ColorDarkRed).
		SetBorderColor(tcell.ColorDarkRed)

	go func() {
		time.Sleep(5 * time.Second)
		app.Stop()
	}()

	if err := app.SetRoot(hangmanGame, true).Run(); err != nil {
		panic(err)
	}
}

package hangman

import (
	"strconv"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Box(H *HangManData) {

	app := tview.NewApplication()

	titleColor := tcell.ColorBlue
	titleBorderColor := tcell.ColorRed

	text1 := (" \n\n+---+\n|   |\nO   |\n/|\\  |\n/ \\  |\n    |\n=========")

	hangmanGame := tview.NewTextView().
		SetText(text1).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	hangmanGame.SetBorder(true).
		SetTitle("Hangman Game🎉").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	lettresTrouvees := tview.NewTextView().
		SetText(H.ToFind).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	lettresTrouvees.SetBorder(true).
		SetTitle("🔍Lettres Trouvées🔎").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	choixLettre := tview.NewTextView().
		SetText(H.Word).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	choixLettre.SetBorder(true).
		SetTitle("Choix Lettre🔠 : ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	attempts := tview.NewTextView().
		SetText(strconv.Itoa(H.Attempts)).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	attempts.SetBorder(true).
		SetTitle("Attempts").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	flex := tview.NewFlex().
		AddItem(hangmanGame, 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(lettresTrouvees, 0, 2, false).
			AddItem(choixLettre, 0, 2, false).
			AddItem(attempts, 5, 2, false), 0, 3, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

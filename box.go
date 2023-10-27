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

	text1 := (" \n\n+---+\n|   |\nO   |\n/|\\  |\n/ \\  |\n    |\n=========") + "\n"

	hangmanGame := tview.NewTextView().
		SetText(text1).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	hangmanGame.SetBorder(true).
		SetTitle("Hangman Game🎉").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	lettresTrouvees := tview.NewTextView().
		SetText("\n\n" + H.Letters).
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

	input := tview.NewInputField().
		SetLabel("Enter a letter: ").
		SetFieldWidth(15).
		SetAcceptanceFunc(tview.InputFieldMaxLength(20))

	input.SetBorder(true).
		SetTitle("Letter : ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			H.LetterInput = input.GetText()
			// if len(H.LetterInput) == 1 {
			if !VerifIfAlreadyUse(H) {
				H.Letters += " | " + H.LetterInput
				Verification(H)
				attempts.SetText(strconv.Itoa(H.Attempts))
				choixLettre.SetText(H.Word)
				lettresTrouvees.SetText(H.Letters) // Mettre à jour le widget lettresTrouvees
				input.SetText("")                  // Effacer le champ d'entrée
				input.SetLabel("Enter a letter: ")
			} else {
				input.SetText("")
				input.SetLabel("Enter a letter not already used :")
			}
			// } else if len(H.LetterInput) > 1{
			// 	print("ok")
			// }
		}
	})

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(input, 0, 1, true).
			AddItem(hangmanGame, 0, 2, false), 0, 3, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(lettresTrouvees, 0, 2, false).
			AddItem(choixLettre, 0, 2, false).
			AddItem(attempts, 5, 2, false), 0, 3, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

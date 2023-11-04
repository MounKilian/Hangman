package hangman

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Box(H *HangManData) {

	app := tview.NewApplication()

	titleColor := tcell.ColorBlue
	titleBorderColor := tcell.ColorRed

	// The box of the Hangman state draw
	hangmanDraw := tview.NewTextView().
		SetText(HangmanState(H)).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	hangmanDraw.SetBorder(true).
		SetTitle(" Hangman State ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	// The box of Letters use
	lettersUse := tview.NewTextView().
		SetText("\n\n" + H.Letters).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	lettersUse.SetBorder(true).
		SetTitle(" Letters Use ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	// The box of Word state
	wordState := tview.NewTextView().
		SetText(H.Word).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	wordState.SetBorder(true).
		SetTitle(" Word State ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	// The box of Attempts remaining
	attempts := tview.NewTextView().
		SetText(strconv.Itoa(H.Attempts)).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	attempts.SetBorder(true).
		SetTitle(" Attempts ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	// The box user input
	input := tview.NewInputField().
		SetLabel(" Enter a letter: ").
		SetFieldWidth(15).
		SetAcceptanceFunc(tview.InputFieldMaxLength(20))

	input.SetBorder(true).
		SetTitle(" Letter Choice ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			H.LetterInput = input.GetText()
			//If the user enter "STOP" that save the current game
			if H.LetterInput == "STOP" {
				app.Stop()
				Save(H)
			}
			//If the user enter a letter
			if !VerifIfAlreadyUse(H) && (H.LetterInput >= "a" && H.LetterInput <= "z") {
				if len(H.LetterInput) == 1 {
					Verification(H)
					NewText(H, hangmanDraw, wordState, lettersUse, attempts, input)
					if WordFind(H) {
						time.Sleep(1 * time.Second)
						app.Stop()
						Victory(H)
					}
					//If the user enter a word
				} else if len(H.LetterInput) > 1 {
					win := EnterWord(H)
					NewText(H, hangmanDraw, wordState, lettersUse, attempts, input)
					if win {
						time.Sleep(1 * time.Second)
						app.Stop()
						Victory(H)
					}
				}
				//If the user enter an invalid or already use letter or word
			} else {
				input.SetText("")
				input.SetLabel("Enter a valid letter not already used :")
			}
			if H.Attempts <= 0 {
				time.Sleep(1 * time.Second)
				app.Stop()
				Defaite(H)
			}
		}
	})

	//Regroup all the item for print
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(input, 0, 1, true).
			AddItem(hangmanDraw, 0, 2, false), 0, 3, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(lettersUse, 0, 2, false).
			AddItem(wordState, 0, 2, false).
			AddItem(attempts, 5, 2, false), 0, 3, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

// Refresh the data after all the input of the user
func NewText(H *HangManData, hangmanDraw, wordState, lettersUse, attempts *tview.TextView, input *tview.InputField) {
	H.Letters += H.LetterInput + " | "
	attempts.SetText(strconv.Itoa(H.Attempts))
	hangmanDraw.SetText(HangmanState(H))
	wordState.SetText(H.Word)
	lettersUse.SetText(H.Letters)
	input.SetText("")
	input.SetLabel("Enter a letter or a word : ")
}

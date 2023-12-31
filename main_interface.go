package hangman

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Box(H *HangManData) {

	app := tview.NewApplication()
	app.EnableMouse(true)

	titleColor := tcell.ColorBlue
	titleBorderColor := tcell.ColorRed

	// The box of the Hangman state draw
	hangmanDraw := tview.NewTextView().
		SetText("\n\n\n\n" + HangmanState(H)).
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
		SetLabelWidth(0).
		SetDynamicColors(true)

	wordState.SetBorder(true).
		SetTitle(" Word State ").
		SetBorderColor(titleBorderColor).
		SetTitleColor(titleColor)

	// The box of Attempts remaining
	attempts := tview.NewTextView().
		SetText("Good Luck, " + strconv.Itoa(H.Attempts) + " attempts remaining").
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

	wordState.SetText(H.Word)
	if H.File != "default.txt" {
		wordState.SetText(ConvertToASCII("ASCII/"+H.File, H))
		wordState.SetTextAlign(tview.AlignLeft)
	} else {
		wordState.SetText("\n\n\n\n\n" + H.Word)
		wordState.SetTextAlign(tview.AlignCenter)
	}

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
				Defeat(H)
			}
		}
	})

	//Regroup all the item for print
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(hangmanDraw, 0, 1, false).AddItem(attempts, 5, 2, false), 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(wordState, 0, 2, false).AddItem(lettersUse, 0, 1, false).AddItem(input, 0, 1, true), 0, 3, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

// Refresh the data after all the input of the user
func NewText(H *HangManData, hangmanDraw, wordState, lettersUse, attempts *tview.TextView, input *tview.InputField) {
	H.Letters += H.LetterInput + " | "
	attempts.SetText("Good Luck, " + strconv.Itoa(H.Attempts) + " attempts remaining")
	hangmanDraw.SetText("\n\n\n\n" + HangmanState(H))
	if H.File != "default.txt" {
		wordState.SetText(ConvertToASCII("ASCII/"+H.File, H))
	} else {
		wordState.SetText("\n\n\n\n\n" +H.Word)
	}
	lettersUse.SetText(H.Letters)
	input.SetText("")
	input.SetLabel("Enter a letter or a word : ")
}

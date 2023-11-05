package hangman

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Print new box when the player win
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

// Print new box when the player lost
func Defeat(H *HangManData) {
	app := tview.NewApplication()

	lost := tview.NewTextView().
		SetText("\n\n\n\n\n\n.##....##..#######..##.....##.......##........#######...######..########\n..##..##..##.....##.##.....##.......##.......##.....##.##....##....##...\n...####...##.....##.##.....##.......##.......##.....##.##..........##...\n....##....##.....##.##.....##.......##.......##.....##..######.....##...\n....##....##.....##.##.....##.......##.......##.....##.......##....##...\n....##....##.....##.##.....##.......##.......##.....##.##....##....##...\n....##.....#######...#######........########..#######...######.....##...\n\n" + HangmanState(H) + "\n\n\n The word was : \n" + H.ToFind).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetTextColor(tcell.ColorRed)
	lost.SetBorder(true).
		SetTitle(" Defeat ").
		SetTitleColor(tcell.ColorDarkRed).
		SetBorderColor(tcell.ColorDarkRed)

	go func() {
		time.Sleep(5 * time.Second)
		app.Stop()
	}()

	if err := app.SetRoot(lost, true).Run(); err != nil {
		panic(err)
	}
}

func Menu() {
	app := tview.NewApplication()

	form := tview.NewForm().
		AddButton("Play", func() {
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
			os.Exit(9)
		})

	form.SetBorder(true).SetTitle(" Welcome to Hangman Game ").SetTitleAlign(tview.AlignCenter).SetTitleColor(tcell.ColorDarkBlue).SetBorderColor(tcell.ColorBlue)

	form.SetButtonsAlign(tview.AlignCenter)

	form.AddTextView("", "\n\n\n               .##......##....########....##...........######......#######.....##.....##....########\n               .##..##..##....##..........##..........##....##....##.....##....###...###....##......\n               .##..##..##....##..........##..........##..........##.....##....####.####....##......\n               .##..##..##....######......##..........##..........##.....##....##.###.##....######..\n               .##..##..##....##..........##..........##..........##.....##....##.....##....##......\n               .##..##..##....##..........##..........##....##....##.....##....##.....##....##......\n               ..###..###.....########....########.....######......#######.....##.....##....########\n", 0, 16, true, true)
	form.SetFieldTextColor(tcell.ColorDarkBlue)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func ConvertToASCII(H *HangManData) string {
	ASCII_word := ""
	letter := ""
	for _, i := range H.Word {
		if i == '_' {
			file, err := os.Open("ASCII/standard.txt")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			line := 0
			for scanner.Scan() {
				line++
				if line == 122 {
					letter += scanner.Text()
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			ASCII_word += letter + " "
			letter = ""
		} else if i >= 'a' && i<= 'z'{
			print(rune(i))
		}
	}
	return ASCII_word
}

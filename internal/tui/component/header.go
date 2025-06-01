package component

import (
	"github.com/agnivade/levenshtein"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

type Header struct {
	App   *tview.Application
	Pages *tview.Pages
}

func (h Header) Build(backFocus tview.Primitive) *tview.Flex {
	// Configure top of screen
	topFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Generate Info Table
	// TODO: Compute these variables
	contextInfo := tview.NewTable().SetSelectable(false, false)
	contextInfo.SetCell(0, 0, tview.NewTableCell("Cloud:"))
	contextInfo.SetCell(0, 1, tview.NewTableCell("ik-vinetos"))
	contextInfo.SetCell(1, 0, tview.NewTableCell("Region:"))
	contextInfo.SetCell(1, 1, tview.NewTableCell("dc4-a"))
	contextInfo.SetCell(2, 0, tview.NewTableCell("User:"))
	contextInfo.SetCell(2, 1, tview.NewTableCell("PCU-V2T9XL4"))

	topFlex.AddItem(contextInfo, 0, 1, false)

	// Configure hideable prompt
	cmdPrompt := h.selectionInputField(topFlex, backFocus)
	// Open prompt to view data
	topFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			// Allow only one prompt to be open
			if !cmdPrompt.HasFocus() {
				topFlex.AddItem(cmdPrompt, 0, 1, true)
				h.App.SetFocus(cmdPrompt)
				return nil // Do not forward key event to cmdPrompt
			}
		}
		return event
	})

	return topFlex
}

var commandList = [...]string{"server", "image"}

func (h Header) selectionInputField(parent *tview.Flex, backFocus tview.Primitive) *tview.InputField {
	// Configure input field
	inputField := tview.NewInputField()
	inputField.SetLabel("> ")
	inputField.SetBorder(true)

	// When existing, give back the focus to data table
	inputField.SetDoneFunc(func(key tcell.Key) {
		// Test if entry is a real page
		isValid := false
		pageName := ""
		for _, word := range commandList {
			wordLower := strings.ToLower(word)
			currentTextLower := strings.ToLower(inputField.GetText())
			if wordLower == currentTextLower {
				isValid = true
				pageName = currentTextLower
				break
			}

		}
		// Reset the input
		inputField.SetText("")
		// Hide the input field
		parent.RemoveItem(inputField)

		if !isValid {
			// Give back the focus
			h.App.SetFocus(backFocus)
		} else {
			// Switch to the asked page
			h.Pages.SwitchToPage(pageName)
		}
	})

	// Configure autocompletion
	inputField.SetAutocompleteFunc(func(currentText string) (entries []string) {
		if len(currentText) == 0 {
			return
		}
		for _, word := range commandList {
			wordLower := strings.ToLower(word)
			currentTextLower := strings.ToLower(currentText)
			if strings.HasPrefix(wordLower, currentTextLower) || levenshtein.ComputeDistance(wordLower, currentTextLower) <= 3 {
				entries = append(entries, word)
			}
		}
		if len(entries) <= 0 {
			entries = nil
		}
		return
	})

	// When applying suggestion, we just replace the input by the suggestion
	inputField.SetAutocompletedFunc(func(text string, index, source int) bool {
		if source != tview.AutocompletedNavigate {
			inputField.SetText(text)
		}
		return source == tview.AutocompletedEnter || source == tview.AutocompletedClick
	})

	return inputField
}

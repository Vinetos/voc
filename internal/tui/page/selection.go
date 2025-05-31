package page

import (
	"github.com/agnivade/levenshtein"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"openstack-tui/internal/openstack"
	"strings"
)

type SelectionPage struct {
}

const SelectionListPage = "selection"

var commandList = [...]string{ServerListPage, ImageListPageName}

func (s SelectionPage) Description() Description {
	return Description{
		Name:    SelectionListPage,
		Resize:  true,
		Visible: true,
	}
}

func (s SelectionPage) Content(app *tview.Application, pages *tview.Pages, client *openstack.Client) tview.Primitive {
	inputField := tview.NewInputField().
		SetLabel("> ")

	inputField.SetDoneFunc(func(key tcell.Key) {
		pages.SwitchToPage(inputField.GetText())

		// Reset input
		inputField.SetText("")
	})

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

	inputField.SetAutocompletedFunc(func(text string, index, source int) bool {
		if source != tview.AutocompletedNavigate {
			inputField.SetText(text)
		}
		return source == tview.AutocompletedEnter || source == tview.AutocompletedClick
	})

	return inputField
}

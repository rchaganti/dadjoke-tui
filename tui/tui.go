package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rchaganti/dadjoke-tui/tui/model"
)

func NewApp(term string, limit int) *tea.Program {
	m, err := model.InitialModel(term, limit)
	if err != nil {
		panic(err)
	}

	prog := tea.NewProgram(m, tea.WithAltScreen())
	return prog
}

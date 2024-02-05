package main

import (
	"github.com/rchaganti/dadjoke-tui/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}

/*
import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	page1 := strings.Builder{}

	i1 := lipgloss.JoinVertical(lipgloss.Center, lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Center).
		Render("Are you sure you want to eat marmalade?"),
	)

	d1 := lipgloss.Place(width, 0, 0, 0,
		dialogBoxStyle.Render(i1),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	page1.WriteString(d1 + "\n")

	i2 := lipgloss.JoinVertical(lipgloss.Center, lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Center).
		Render("Are you sure you want to eat marmalade?"),
	)

	d2 := lipgloss.Place(width, 0, 0, 0,
		dialogBoxStyle.Render(i2),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	page1.WriteString(d2 + "\n")

	p := tea.NewProgram(
		model{
			content:     string(page1.String()),
			currentPage: 1,
			totalPages:  2,
		},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
*/

package joke

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	Width  = 96
)

var DialogBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#874BFD")).
	Padding(1, 0).
	BorderTop(true).
	BorderLeft(true).
	BorderRight(true).
	BorderBottom(true)

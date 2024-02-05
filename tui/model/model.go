package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	dj "github.com/rchaganti/dadjoke-go"
	"github.com/rchaganti/dadjoke-tui/styles/joke"
	"github.com/rchaganti/dadjoke-tui/styles/page"
	"github.com/rchaganti/dadjoke-tui/tui/keybinding"
)

const useHighPerformanceRenderer = false

type Result struct {
	ID   string
	Joke string
}

type Model struct {
	SearchTerm   string
	CurrentPage  int
	TotalPages   int
	ItemsPerPage int
	Viewport     viewport.Model
	Content      []Result
	Help         help.Model
	Keys         keybinding.KeyMap
	Ready        bool
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Left):
			{
				if m.CurrentPage == 1 {
					return m, nil
				}
				m.CurrentPage--
				sr, err := getSearchResult(m.SearchTerm, m.ItemsPerPage, m.CurrentPage)
				if err != nil {
					return nil, nil
				}
				content := getJokes(sr.Results)
				m.Content = content
				m.Viewport.SetContent(updateContent(m.Content))
				return m, nil
			}
		case key.Matches(msg, m.Keys.Right):
			{
				if m.CurrentPage == m.TotalPages {
					return m, nil
				}
				m.CurrentPage++
				sr, err := getSearchResult(m.SearchTerm, m.ItemsPerPage, m.CurrentPage)
				if err != nil {
					return nil, nil
				}
				content := getJokes(sr.Results)
				m.Content = content
				m.Viewport.SetContent(updateContent(m.Content))
				return m, nil
			}
		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.HeaderView())
		footerHeight := lipgloss.Height(m.FooterView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.Viewport.YPosition = headerHeight
			m.Viewport.HighPerformanceRendering = useHighPerformanceRenderer

			content := updateContent(m.Content)
			m.Viewport.SetContent(content)
			m.Ready = true

			m.Viewport.YPosition = headerHeight + 1
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.Viewport))
		}
	}

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.Ready {
		return "\n  Initializing..."
	}
	helpView := m.Help.View(m.Keys)

	return fmt.Sprintf("%s\n%s\n%s\n%s", m.HeaderView(), m.Viewport.View(), m.FooterView(), helpView)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) HeaderView() string {
	title := page.TitleStyle.Render("Dad Jokes")
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) FooterView() string {
	info := page.InfoStyle.Render(fmt.Sprintf("Page %d of %d", m.CurrentPage, m.TotalPages))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateContent(content []Result) string {
	var sb strings.Builder
	var d string
	for _, c := range content {
		item := lipgloss.JoinVertical(lipgloss.Center, lipgloss.NewStyle().
			Width(50).
			Align(lipgloss.Center).
			Render(c.Joke),
		)

		d = lipgloss.Place(joke.Width, 0, 0, 0,
			joke.DialogBoxStyle.Render(item),
			lipgloss.WithWhitespaceForeground(joke.Subtle),
		)

		sb.WriteString(d + "\n")
	}
	return sb.String()
}

func InitialModel(term string, limit int) (Model, error) {
	sr, err := getSearchResult(term, limit, 1)
	if err != nil {
		return Model{}, err
	}

	if sr.TotalJokes != 0 {
		jokes := getJokes(sr.Results)

		return Model{
			SearchTerm:   term,
			Content:      jokes,
			CurrentPage:  1,
			TotalPages:   sr.TotalPages,
			ItemsPerPage: limit,
			Viewport:     viewport.Model{},
			Help:         help.New(),
			Keys:         keybinding.Keys,
		}, nil
	}

	return Model{}, err
}

func getSearchResult(term string, limit int, page int) (dj.Search, error) {
	client := dj.NewClient()

	sr, err := client.SearchDadJokes(term, page, limit)
	if err != nil {
		return dj.Search{}, err
	}

	return sr, nil
}

func getJokes(j []dj.Joke) []Result {
	jokes := make([]Result, len(j))
	for i, joke := range j {
		jokes[i] = Result{
			ID:   joke.ID,
			Joke: joke.Joke,
		}
	}
	return jokes
}

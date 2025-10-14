package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
	terminal terminal
}

type terminal struct {
	height int
	width  int
}

func NewModel() *model {
	return &model{
		cursor:   0,
		choices:  []string{"I am gay", "First Choice", "Second Choice", "Third Choice", "Fourth Choice", "Fifth Choice", "Sixth Choice", "Seventh Choice", "Eighth Choice", "Ninth Choice", "Tenth Choice"},
		selected: make(map[int]struct{}),
		terminal: terminal{height: 0, width: 0},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch messages := msg.(type) {
	case tea.KeyMsg:
		switch messages.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ", "enter":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected = map[int]struct{}{m.cursor: {}}
			}

		}
	case tea.WindowSizeMsg:
		m.terminal.width = messages.Width
		m.terminal.height = messages.Height
	}
	return m, nil
}

func (m *model) View() string {
	maxLen := 0
	for _, item := range m.choices {
		if len(item) > maxLen {
			maxLen = len(item)
		}
	}
	content := "What do you want to do?\n\n"
	for i, item := range m.choices {
		cursor := "  "
		selected := "[ ]"
		if m.cursor == i {
			cursor = "->"
		}
		if _, ok := m.selected[i]; ok {
			selected = "[x]"
		}

		paddedItem := fmt.Sprintf("%-*s", maxLen, item) // left pad to same width
		line := fmt.Sprintf("%v %v %v", cursor, selected, paddedItem)

		if m.cursor == i {
			line = pointedLineStyle.Render(line)
		} else if _, ok := m.selected[i]; ok {
			line = selectedLineStyle.Render(line)
		} else {
			line = normalLineStyle.Render(line)
		}

		content = fmt.Sprintf("%v\n%v", content, line)
	}
	outerWidth := int(float64(m.terminal.width) * 0.8)
	outerHeight := int(float64(m.terminal.height) * 0.8)
	innerWidth := outerWidth - 4
	innerHeight := 5
	return CenterSquareWithContent(m.terminal.width, m.terminal.height, outerWidth, outerHeight, innerWidth, innerHeight, content)
}

func CenterSquareWithContent(
	appWidth, appHeight, outerWidth, outerHeight, innerWidth, innerHeight int,
	content string,
) string {
	outerStyle := lipgloss.NewStyle().
		Width(outerWidth).
		Height(outerHeight).
		Margin((appHeight-outerHeight)/2, (appWidth-outerWidth)/2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#E95420"))

	innerRendered := lipgloss.Place(outerWidth, outerHeight, lipgloss.Center, lipgloss.Center, content)
	return outerStyle.Render(innerRendered)
}

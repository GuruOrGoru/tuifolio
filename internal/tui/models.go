package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

func NewModel() *model {
	return &model{
		count: 0,
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
		case "i", "+":
			m.count++
		case "d", "-":
			m.count--
		}
	}
	return m, nil
}

func (m *model) View() string {
	return fmt.Sprintf(
		"Counter: %d\n\nPress 'i' to increase, 'd' to decrease, 'q' to quit.\n",
		m.count,
	)
}

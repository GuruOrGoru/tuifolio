package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	m := NewModel()
	if m.Cursor != 0 {
		t.Errorf("expected cursor 0, got %d", m.Cursor)
	}
	if len(m.Choices) <= 10 {
		t.Errorf("expected more than 10 choices, got %d", len(m.Choices))
	}
	if m.Terminal.Width != 0 || m.Terminal.Height != 0 {
		t.Errorf("expected terminal size 0x0, got %dx%d", m.Terminal.Width, m.Terminal.Height)
	}
	if m.State != StateLogo {
		t.Errorf("expected state StateLogo, got %v", m.State)
	}
}

func TestUpdateWindowSize(t *testing.T) {
	m := NewModel()
	msg := tea.WindowSizeMsg{Width: 80, Height: 24}
	_, _ = m.Update(msg)
	if m.Terminal.Width != 80 || m.Terminal.Height != 24 {
		t.Errorf("expected terminal size 80x24, got %dx%d", m.Terminal.Width, m.Terminal.Height)
	}
}

func TestViewCentering(t *testing.T) {
	m := NewModel()
	m.State = StateMain
	msg := tea.WindowSizeMsg{Width: 50, Height: 30}
	_, _ = m.Update(msg)

	view := m.View()
	lines := strings.Split(view, "\n")

	if len(lines) != 35 {
		t.Errorf("expected 35 lines, got %d", len(lines))
	}

	contentLines := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			contentLines++
		}
	}
	if contentLines < 10 {
		t.Errorf("expected at least 10 content lines, got %d", contentLines)
	}
}

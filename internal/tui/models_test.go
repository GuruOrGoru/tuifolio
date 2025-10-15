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
	if len(m.Choices) != 11 {
		t.Errorf("expected 11 choices, got %d", len(m.Choices))
	}
	if m.Terminal.Width != 0 || m.Terminal.Height != 0 {
		t.Errorf("expected terminal size 0x0, got %dx%d", m.Terminal.Width, m.Terminal.Height)
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
	m.ShowLogo = false
	msg := tea.WindowSizeMsg{Width: 50, Height: 30}
	_, _ = m.Update(msg)

	view := m.View()
	lines := strings.Split(view, "\n")

	if len(lines) != 30 {
		t.Errorf("expected 30 lines, got %d", len(lines))
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

func TestViewWithSelection(t *testing.T) {
	m := NewModel()
	m.ShowLogo = false
	msg := tea.WindowSizeMsg{Width: 50, Height: 30}
	_, _ = m.Update(msg)

	keyMsg := tea.KeyMsg{Type: tea.KeyEnter}
	_, _ = m.Update(keyMsg)

	view := m.View()
	lines := strings.Split(view, "\n")

	if len(lines) != 30 {
		t.Errorf("expected 30 lines, got %d", len(lines))
	}
	found := false
	for _, line := range lines {
		if strings.Contains(line, "[X]") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected selected item with [X], not found")
	}
}

func TestViewWithCursor(t *testing.T) {
	m := NewModel()
	m.ShowLogo = false
	msg := tea.WindowSizeMsg{Width: 50, Height: 30}
	_, _ = m.Update(msg)

	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}
	_, _ = m.Update(keyMsg)

	view := m.View()
	lines := strings.Split(view, "\n")

	if len(lines) != 30 {
		t.Errorf("expected 30 lines, got %d", len(lines))
	}
	found := false
	for _, line := range lines {
		if strings.Contains(line, "->") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected cursor with ->, not found")
	}
}

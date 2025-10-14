package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	m := NewModel()
	if m.cursor != 0 {
		t.Errorf("expected cursor 0, got %d", m.cursor)
	}
	if len(m.choices) != 3 {
		t.Errorf("expected 3 choices, got %d", len(m.choices))
	}
	if m.terminal.width != 0 || m.terminal.height != 0 {
		t.Errorf("expected terminal size 0x0, got %dx%d", m.terminal.width, m.terminal.height)
	}
}

func TestUpdateWindowSize(t *testing.T) {
	m := NewModel()
	msg := tea.WindowSizeMsg{Width: 80, Height: 24}
	_, _ = m.Update(msg)
	if m.terminal.width != 80 || m.terminal.height != 24 {
		t.Errorf("expected terminal size 80x24, got %dx%d", m.terminal.width, m.terminal.height)
	}
}

func TestViewCentering(t *testing.T) {
	m := NewModel()
	// Set terminal size
	msg := tea.WindowSizeMsg{Width: 50, Height: 10}
	_, _ = m.Update(msg)

	view := m.View()
	lines := strings.Split(view, "\n")

	// Should have 10 lines total (height)
	if len(lines) != 10 {
		t.Errorf("expected 10 lines, got %d", len(lines))
	}

	// Content lines: 3 choices
	// Padding top: (10-3)/2 = 3
	// Padding bottom: 10-3-3 = 4
	// So lines 3-5 should have content, others empty or spaces

	// Check that content is centered horizontally
	// Each line should be padded to 50 width
	for i, line := range lines {
		if len(line) != 50 {
			t.Errorf("line %d length %d, expected 50", i, len(line))
		}
	}

	// Check vertical centering: content in lines 3,4,5 (0-based: 3,4,5)
	// Assuming the content lines are the ones with text
	contentLines := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			contentLines++
		}
	}
	if contentLines != 3 {
		t.Errorf("expected 3 content lines, got %d", contentLines)
	}
}

func TestViewWithSelection(t *testing.T) {
	m := NewModel()
	// Set terminal size
	msg := tea.WindowSizeMsg{Width: 50, Height: 10}
	_, _ = m.Update(msg)

	// Select first item
	keyMsg := tea.KeyMsg{Type: tea.KeyEnter}
	_, _ = m.Update(keyMsg)

	view := m.View()
	lines := strings.Split(view, "\n")

	// Should still be centered
	if len(lines) != 10 {
		t.Errorf("expected 10 lines, got %d", len(lines))
	}
	// Check that selection is rendered (contains [x])
	found := false
	for _, line := range lines {
		if strings.Contains(line, "[x]") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected selected item with [x], not found")
	}
}

func TestViewWithCursor(t *testing.T) {
	m := NewModel()
	// Set terminal size
	msg := tea.WindowSizeMsg{Width: 50, Height: 10}
	_, _ = m.Update(msg)

	// Move cursor down
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}
	_, _ = m.Update(keyMsg)

	view := m.View()
	lines := strings.Split(view, "\n")

	// Should still be centered
	if len(lines) != 10 {
		t.Errorf("expected 10 lines, got %d", len(lines))
	}
	// Check that cursor is rendered (contains ->)
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

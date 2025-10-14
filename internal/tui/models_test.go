package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	model := NewModel()
	if model == nil {
		t.Error("expected model to be non-nil")
	}
	if model.count != 0 {
		t.Errorf("expected initial count 0, got %d", model.count)
	}
}

func TestModel_Update(t *testing.T) {
	m := NewModel()

	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'i'}})
	if newModel.(*model).count != 1 {
		t.Errorf("expected count 1 after 'i', got %d", newModel.(*model).count)
	}

	newModel, _ = newModel.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}})
	if newModel.(*model).count != 0 {
		t.Errorf("expected count 0 after 'd', got %d", newModel.(*model).count)
	}

	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Error("expected quit command on 'q'")
	}
}

func TestModel_View(t *testing.T) {
	model := NewModel()
	view := model.View()
	if !strings.Contains(view, "Counter: 0") {
		t.Errorf("expected view to contain 'Counter: 0', got %s", view)
	}
	if !strings.Contains(view, "Press 'i' to increase") {
		t.Error("expected view to contain instructions")
	}
}

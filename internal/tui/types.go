package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type Tab int

const (
	TabHome Tab = iota
	TabSkills
	TabProjects
	TabContact
	TabCount
)

type AppState int

const (
	StateLogo AppState = iota
	StateUsernameInput
	StateMain
)

type (
	LogoDone       struct{}
	CursorBlinkMsg struct{}
)

func (t Tab) String() string {
	switch t {
	case TabHome:
		return "Home"
	case TabSkills:
		return "Skills"
	case TabProjects:
		return "Projects"
	case TabContact:
		return "Contact"
	}
	return ""
}

type KeyMap struct {
	Up          key.Binding
	Down        key.Binding
	Help        key.Binding
	Quit        key.Binding
	Select      key.Binding
	TabForward  key.Binding
	TabBackward key.Binding
	Tab1        key.Binding
	Tab2        key.Binding
	Tab3        key.Binding
	Tab4        key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select},
		{k.Help, k.Quit, k.TabForward, k.TabBackward},
		{k.Tab1, k.Tab2, k.Tab3, k.Tab4},
	}
}

type Terminal struct {
	Height int
	Width  int
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Select: key.NewBinding(
		key.WithKeys(" ", "enter"),
		key.WithHelp("space/enter", "select"),
	),
	TabForward: key.NewBinding(
		key.WithKeys("tab", "right", "l", "w"),
		key.WithHelp("tab/Vim Motions For right", "next tab"),
	),
	TabBackward: key.NewBinding(
		key.WithKeys("shift+tab", "left", "h", "b"),
		key.WithHelp("shift+tab/Vim Motions For left", "previous tab"),
	),
	Tab1: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "go to tab 1"),
	),
	Tab2: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "go to tab 2"),
	),
	Tab3: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "go to tab 3"),
	),
	Tab4: key.NewBinding(
		key.WithKeys("4"),
		key.WithHelp("4", "go to tab 4"),
	),
}

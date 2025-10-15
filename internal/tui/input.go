package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func RenderUsernameInput(terminal Terminal, input textinput.Model) string {
	content := "Welcome to GuruOrGoru!\n\nPlease enter your username:\n\n" + input.View()
	return lipgloss.Place(terminal.Width, terminal.Height, lipgloss.Center, lipgloss.Center, content)
}

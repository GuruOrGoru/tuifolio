package tui

import "github.com/charmbracelet/lipgloss"

func RenderLogo(terminal Terminal, cursorVisible bool) string {
	logo := "GuruOrGoru"
	cursor := " "
	if cursorVisible {
		cursor = "█"
	}
	logo = LogoStyle.Render(logo) + CursorStyle.Render(cursor)
	return lipgloss.Place(terminal.Width, terminal.Height, lipgloss.Center, lipgloss.Center, logo)
}

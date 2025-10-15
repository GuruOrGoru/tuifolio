package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

func RenderLogo(terminal Terminal, cursorVisible bool) string {
	logo := "GuruOrGoru"
	cursor := " "
	if cursorVisible {
		cursor = "█"
	}
	logo = LogoStyle.Render("GuruOrGoru") + CursorStyle.Render(cursor)
	return lipgloss.Place(terminal.Width, terminal.Height, lipgloss.Center, lipgloss.Center, logo)
}

func RenderHeader() string {
	return HeaderStyle.Render("GuruOrGoru TUI")
}

func RenderList(cursor int, choices []string, selected map[int]struct{}) string {
	maxLen := 0
	for _, item := range choices {
		if len(item) > maxLen {
			maxLen = len(item)
		}
	}

	lines := []string{}
	for i, choice := range choices {
		cursorStr := "  "
		if cursor == i {
			cursorStr = "->"
		}

		checked := " "
		if _, ok := selected[i]; ok {
			checked = "X"
		}

		paddedChoice := fmt.Sprintf("%-*s", maxLen, choice)
		text := fmt.Sprintf("%s [%s] %s", cursorStr, checked, paddedChoice)

		if cursor == i {
			lines = append(lines, PointedLineStyle.Render(text))
		} else if _, ok := selected[i]; ok {
			lines = append(lines, SelectedLineStyle.Render(text))
		} else {
			lines = append(lines, NormalLineStyle.Render(text))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Top, lines...)
}

func RenderHelp(h help.Model, keys KeyMap) string {
	return h.View(keys)
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

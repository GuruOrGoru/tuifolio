package tui

import "github.com/charmbracelet/lipgloss"

func CenterSquareWithContent(
	appWidth, appHeight, outerWidth, outerHeight, innerWidth, innerHeight int,
	content string,
) string {
	innerRendered := lipgloss.Place(outerWidth, outerHeight, lipgloss.Center, lipgloss.Center, content)
	outerRendered := CenterOuterStyle.Width(outerWidth).Height(outerHeight).Render(innerRendered)
	return lipgloss.Place(appWidth, appHeight, lipgloss.Center, lipgloss.Center, outerRendered)
}

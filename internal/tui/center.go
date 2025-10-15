package tui

import "github.com/charmbracelet/lipgloss"

func CenterSquareWithContent(
	appWidth, appHeight, outerWidth, outerHeight, innerWidth, innerHeight int,
	content string,
) string {
	outerStyle := CenterOuterStyle.
		Width(outerWidth).
		Height(outerHeight).
		Margin((appHeight-outerHeight)/2, (appWidth-outerWidth)/2)

	innerRendered := lipgloss.Place(outerWidth, outerHeight, lipgloss.Center, lipgloss.Center, content)
	return outerStyle.Render(innerRendered)
}

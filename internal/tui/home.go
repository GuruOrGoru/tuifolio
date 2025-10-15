package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderHomeTab(cursor int, choices []string, selected map[int]struct{}) string {
	content := ""

	if len(choices) == 0 {
		content = "No choices available."
	} else {
		content = "Hi! I'm Siddhartha Dhakall.\n\n"
		content += "I'm a passionate developer with experience in building web applications, mobile apps, and more.\n\n"
		content += "Here are some of my skills and interests:\n\n"

		for i, choice := range choices {
			prefix := " [ ]"
			if _, ok := selected[i]; ok {
				prefix = "[X]"
			}
			if i == cursor {
				prefix = ">" + prefix[1:]
				content += PointedLineStyle.Render(prefix + choice) + "\n"
			} else {
				content += NormalLineStyle.Render(prefix + choice) + "\n"
			}
		}

		content += "\nUse the arrow keys to navigate and space to select/deselect skills."
	}

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top,
			TitleStyle.Render("Welcome to My Portfolio"),
			content,
		),
	)
}

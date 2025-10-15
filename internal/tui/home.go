package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderHomeTab(cursor int, choices []string, selected map[int]struct{}) string {
	content := ""

	if len(choices) == 0 {
		content = "No choices available."
	} else {
		content = "Hi! I'm Siddhartha Dhakal.\n\n"
		content += "Cloud native Go developer with 3+ years programming experience. Hackathon winner, open source contributor, and passionate about building scalable applications.\n\n"
		content += "Here are some of my skills and interests:\n\n"

		for i, choice := range choices {
			prefix := "[ ]"
			if _, ok := selected[i]; ok {
				prefix = "[X]"
			}
			if i == cursor {
				prefix = "->" + prefix[0:]
				content += PointedLineStyle.Render(prefix+choice) + "\n"
			} else if _, ok := selected[i]; ok {
				content += SelectedLineStyle.Render(prefix+choice) + "\n"
			} else {
				content += NormalLineStyle.Render(prefix+choice) + "\n"
			}
		}

		content += "\nHope you'll Love it!"
	}

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top,
			TitleStyle.Render("Welcome to My Portfolio"),
			content,
		),
	)
}

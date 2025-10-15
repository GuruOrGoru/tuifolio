package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderContactTab(cursor int) string {

	contacts := []struct {
		label string
		value string
	}{
		{"Email", "siddharthadhakall3722@gmail.com"},
		{"LinkedIn", "linkedin.com/in/yourprofile"},
		{"GitHub", "github.com/yourusername"},
		{"Twitter", "@yourhandle"},
		{"Website", "yourwebsite.com"},
	}

	var contentParts []string
	contentParts = append(contentParts, TitleStyle.Render("Get In Touch"))

	contentParts = append(contentParts, DescriptionStyle.Render("I'm always open to discussing new opportunities, collaborations, or just having a chat about technology!"))

	for i, contact := range contacts {
		prefix := " [ ]"
		if i == cursor {
			prefix = ">" + prefix[1:]
		}
		contentParts = append(contentParts, ContactItemStyle.Render(
			prefix+" "+ContactLabelStyle.Render(contact.label+": ")+ContactValueStyle.Render(contact.value),
		))
	}

	contentParts = append(contentParts, "")
	contentParts = append(contentParts, FooterStyle.Render("Use arrow keys to navigate, space to select contact method."))
	contentParts = append(contentParts, FooterStyle.Render("Feel free to reach out anytime!"))

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentParts...),
	)
}

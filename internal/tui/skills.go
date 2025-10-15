package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderSkillsTab(cursor int, expanded map[int]bool) string {

	skills := []struct {
		category string
		items    []string
	}{
		{
			category: "Programming Languages",
			items:    []string{"Go", "Python", "JavaScript", "TypeScript", "Rust"},
		},
		{
			category: "Web Development",
			items:    []string{"React", "Node.js", "HTML5", "CSS3", "Next.js"},
		},
		{
			category: "Databases",
			items:    []string{"PostgreSQL", "MongoDB", "Redis", "MySQL"},
		},
		{
			category: "Tools & Technologies",
			items:    []string{"Docker", "Kubernetes", "Git", "Linux", "AWS"},
		},
		{
			category: "Other Skills",
			items:    []string{"System Design", "API Development", "Testing", "CI/CD"},
		},
	}

	var contentParts []string
	contentParts = append(contentParts, TitleStyle.Render("Skills & Technologies"))

	for i, skill := range skills {
		prefix := " [ ]"
		if i == cursor {
			prefix = ">" + prefix[1:]
		}
		if expanded[i] {
			prefix = prefix[:3] + "+" + prefix[4:]
		} else {
			prefix = prefix[:3] + "-" + prefix[4:]
		}
		contentParts = append(contentParts, SectionStyle.Render(prefix+" "+skill.category))
		if expanded[i] {
			for _, item := range skill.items {
				contentParts = append(contentParts, ItemStyle.Render("  • "+item))
			}
		}
	}

	contentParts = append(contentParts, "")
	contentParts = append(contentParts, FooterStyle.Render("Use arrow keys to navigate, space to toggle category visibility."))
	contentParts = append(contentParts, FooterStyle.Render("Always learning and exploring new technologies!"))

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentParts...),
	)
}

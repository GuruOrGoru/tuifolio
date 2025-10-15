package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderSkillsTab() string {

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

	for _, skill := range skills {
		contentParts = append(contentParts, SectionStyle.Render(skill.category))
		for _, item := range skill.items {
			contentParts = append(contentParts, ItemStyle.Render("• "+item))
		}
	}

	contentParts = append(contentParts, "")
	contentParts = append(contentParts, FooterStyle.Render("Always learning and exploring new technologies!"))

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentParts...),
	)
}

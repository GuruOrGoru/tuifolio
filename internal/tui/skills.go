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
			items:    []string{"Go", "C", "C#", "Ruby", "Python", "Elixir(just started)"},
		},
		{
			category: "Web Development",
			items:    []string{"HTMX", "React", "Gin", "Chi", "Gorm", "Fiber", "Postman"},
		},
		{
			category: "Databases",
			items:    []string{"SQLite", "PostgreSQL", "Supabase", "MongoDB"},
		},
		{
			category: "Tools & Technologies",
			items:    []string{"BubbleTea", "Nvim", "Git", "Bruno", "Wish", "Docker", "Kubernetes", "Linux"},
		},
		{
			category: "Other Skills",
			items:    []string{"CLI Development", "TUI Development", "SSH Servers", "API Security", "Algorithms"},
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
	contentParts = append(contentParts, FooterStyle.Render("Use arrow keys/Vim motions to navigate, space to toggle category visibility."))
	contentParts = append(contentParts, FooterStyle.Render("Always learning and exploring new technologies!"))

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentParts...),
	)
}

package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderProjectsTab() string {

	projects := []struct {
		title        string
		description  string
		technologies []string
	}{
		{
			title:        "Terminal Portfolio",
			description:  "A beautiful terminal-based portfolio application built with Go and Bubble Tea. Features interactive tabs, smooth animations, and responsive design.",
			technologies: []string{"Go", "Bubble Tea", "Lipgloss", "TUI"},
		},
		{
			title:        "E-commerce Platform",
			description:  "Full-stack e-commerce solution with user authentication, payment processing, and admin dashboard. Built with modern web technologies.",
			technologies: []string{"React", "Node.js", "PostgreSQL", "Stripe"},
		},
		{
			title:        "Task Management API",
			description:  "RESTful API for task management with real-time updates, user collaboration, and comprehensive testing suite.",
			technologies: []string{"Go", "Gin", "WebSocket", "PostgreSQL"},
		},
		{
			title:        "Data Visualization Dashboard",
			description:  "Interactive dashboard for data analysis and visualization with multiple chart types and real-time data updates.",
			technologies: []string{"Python", "FastAPI", "D3.js", "MongoDB"},
		},
	}

	var contentParts []string
	contentParts = append(contentParts, TitleStyle.Render("Featured Projects"))

	for _, project := range projects {
		contentParts = append(contentParts, ProjectItemStyle.Render(
			ProjectTitleStyle.Render(project.title),
		))
		contentParts = append(contentParts, ProjectDescStyle.Render(project.description))

		var techs []string
		for _, tech := range project.technologies {
			techs = append(techs, TechStyle.Render(tech))
		}
		contentParts = append(contentParts, lipgloss.JoinHorizontal(lipgloss.Top, techs...))
		contentParts = append(contentParts, "")
	}

	contentParts = append(contentParts, FooterStyle.Render("Check out my GitHub for more projects and contributions."))

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentParts...),
	)
}

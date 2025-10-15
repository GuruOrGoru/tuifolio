package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderProjectsTab(cursor int) string {

	projects := []struct {
		title string
		link  string
	}{
		{
			title: "Terminal Portfolio",
			link:  "https://github.com/user/terminal-portfolio",
		},
		{
			title: "E-commerce Platform",
			link:  "https://github.com/user/ecommerce-platform",
		},
		{
			title: "Task Management API",
			link:  "https://github.com/user/task-api",
		},
		{
			title: "Data Visualization Dashboard",
			link:  "https://github.com/user/data-dashboard",
		},
	}

	var contentParts []string
	contentParts = append(contentParts, TitleStyle.Render("Featured Projects"))

	for i, project := range projects {
		prefix := " [ ]"
		if i == cursor {
			prefix = ">" + prefix[1:]
		}
		contentParts = append(contentParts, ItemStyle.Render(prefix+" "+project.title+" ("+project.link+")"))
	}

	contentParts = append(contentParts, "")
	contentParts = append(contentParts, FooterStyle.Render("Use arrow keys or j/k to navigate projects, space/enter to view details."))
	contentParts = append(contentParts, FooterStyle.Render("Check out my GitHub for more projects and contributions."))

	return BoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentParts...),
	)
}

func RenderProjectModal(projectIndex int, width int, height int) string {
	projects := []struct {
		title        string
		link         string
		description  string
		technologies []string
	}{
		{
			title:        "Terminal Portfolio",
			link:         "https://github.com/user/terminal-portfolio",
			description:  "A beautiful terminal-based portfolio application built with Go and Bubble Tea. Features interactive tabs, smooth animations, and responsive design.",
			technologies: []string{"Go", "Bubble Tea", "Lipgloss", "TUI"},
		},
		{
			title:        "E-commerce Platform",
			link:         "https://github.com/user/ecommerce-platform",
			description:  "Full-stack e-commerce solution with user authentication, payment processing, and admin dashboard. Built with modern web technologies.",
			technologies: []string{"React", "Node.js", "PostgreSQL", "Stripe"},
		},
		{
			title:        "Task Management API",
			link:         "https://github.com/user/task-api",
			description:  "RESTful API for task management with real-time updates, user collaboration, and comprehensive testing suite.",
			technologies: []string{"Go", "Gin", "WebSocket", "PostgreSQL"},
		},
		{
			title:        "Data Visualization Dashboard",
			link:         "https://github.com/user/data-dashboard",
			description:  "Interactive dashboard for data analysis and visualization with multiple chart types and real-time data updates.",
			technologies: []string{"Python", "FastAPI", "D3.js", "MongoDB"},
		},
	}

	project := projects[projectIndex]

	var contentParts []string
	contentParts = append(contentParts, TitleStyle.Render(project.title))
	contentParts = append(contentParts, lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render("Link: "+project.link))
	contentParts = append(contentParts, "")
	contentParts = append(contentParts, SectionStyle.Render("Description"))
	contentParts = append(contentParts, DescriptionStyle.Render(project.description))
	contentParts = append(contentParts, "")
	contentParts = append(contentParts, SectionStyle.Render("Tech Stack"))
	var techs []string
	for _, tech := range project.technologies {
		techs = append(techs, TechStyle.Render("• "+tech))
	}
	contentParts = append(contentParts, lipgloss.JoinVertical(lipgloss.Top, techs...))
	contentParts = append(contentParts, "")
	contentParts = append(contentParts, FooterStyle.Render("Press esc or q to close modal."))

	modalContent := lipgloss.JoinVertical(lipgloss.Top, contentParts...)

	modalStyle := lipgloss.NewStyle().
		Width(width - 4).
		Height(height - 4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1).
		Align(lipgloss.Left)

	return modalStyle.Render(modalContent)
}

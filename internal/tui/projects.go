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
			title: "Pokedex CLI",
			link:  "https://github.com/GuruOrGoru/pokedex",
		},
		{
			title: "Inclusiv",
			link:  "https://github.com/GuruOrGoru/inclusiv",
		},
		{
			title: "Terminal Portfolio",
			link:  "https://github.com/GuruOrGoru/tuifolio",
		},
		{
			title: "Password Manager",
			link:  "https://github.com/GuruOrGoru/passguru",
		},
		{
			title: "Personal Website",
			link:  "https://github.com/GuruOrGoru/GuruOrGoru.github.io",
		},
		{
			title: "Adarsha-School-Server",
			link:  "https://github.com/GuruOrGoru/adarsha-server",
		},
		{
			title: "Fragiment(Code Snippet Sharing Platform)",
			link:  "https://github.com/GuruOrGoru/Fragiment",
		},
		{
			title: "GuruVim",
			link:  "github.com/GuruOrGoru/guruvim",
		},
		{
			title: "interpreter-for-guruverbal",
			link:  "https://github.com/GuruOrGoru/Making-An-Interpreter-in-Go",
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
			title:        "Pokedex CLI",
			link:         "https://github.com/GuruOrGoru/pokedex",
			description:  "A command-line REPL for a Pokemon encyclopedia built with Go. Includes utility functions for input cleaning and command handling.",
			technologies: []string{"Go"},
		},
		{
			title:        "Inclusiv",
			link:         "https://github.com/GuruOrGoru/inclusiv",
			description:  "AI-powered job portal focused on inclusivity, filtering job listings for accessibility, diversity, and flexibility. Built with Go backend and modern web technologies.",
			technologies: []string{"Go", "HTMX", "Tailwind", "Supabase"},
		},
		{
			title:        "Terminal Portfolio",
			link:         "https://github.com/GuruOrGoru/tuifolio",
			description:  "A basic terminal application built with Bubble Tea. Displays 'Hello, World!' and handles user input for quitting.",
			technologies: []string{"Go", "Bubble Tea"},
		},
		{
			title:        "Password Manager",
			link:         "https://github.com/GuruOrGoru/passguru",
			description:  "A secure password manager built with Go and Chi router. Includes environment configuration and vendor dependencies for robust functionality.",
			technologies: []string{"Go", "Chi", "Security"},
		},
		{
			title:        "Personal Website",
			link:         "https://github.com/GuruOrGoru/GuruOrGoru.github.io",
			description:  "Horror-themed personal website with interactive animations, sound effects, and dynamic content. Built with modern web technologies for an immersive experience.",
			technologies: []string{"HTML", "CSS", "JavaScript", "Alpine.js", "GSAP", "Howler", "Tailwind CSS"},
		},
		{
			title:        "Adarsha-School-Server",
			link:         "https://github.com/GuruOrGoru/adarsha-server",
			description:  "Backend for the Adarsha School Website, a web application designed to provide information about the school, manage content, and handle school-related functionalities. Utilizes Go templates for dynamic content serving, modular routing with go-chi/chi, and CORS support.",
			technologies: []string{"Go", "Chi", "Supabase"},
		},
		{
			title:        "Fragiment(Code Snippet Sharing Platform)",
			link:         "https://github.com/GuruOrGoru/Fragiment",
			description:  "A Rails application for sharing code snippets with syntax highlighting, user authentication, comments, likes, and admin panel. Features MacBook-style code preview and responsive UI.",
			technologies: []string{"Ruby on Rails", "Tailwind CSS", "PostgreSQL"},
		},
		{
			title:        "GuruVim",
			link:         "github.com/GuruOrGoru/guruvim",
			description:  "A personal Neovim setup using lazy.nvim for plugin management, with plugins for autopair, completion, LSP, telescope, and more.",
			technologies: []string{"Vim", "Shell"},
		},
		{
			title:        "interpreter-for-guruverbal",
			link:         "https://github.com/GuruOrGoru/Making-An-Interpreter-in-Go",
			description:  "An interpreter for a custom programming language implemented in Go, following the book 'Writing An Interpreter In Go'.",
			technologies: []string{"Go"},
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

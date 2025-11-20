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
			title: "The GURU MMO",
			link:  "https://github.com/GuruOrGoru/mmo",
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
			title: "Interpreter",
			link:  "https://github.com/GuruOrGoru/making-an-interpreter-in-go",
		},
		{
			title: "Lsp-Server",
			link:  "https://github.com/GuruOrGoru/lsp-server",
		},
		{
			title: "HTTP Server IN go",
			link:  "https://github.com/GuruOrGoru/http-server-in-go",
		},
		{
			title: "Fragiment(Code Snippet Sharing Platform)",
			link:  "https://github.com/GuruOrGoru/Fragiment",
		},
		{
			title: "GuruVim",
			link:  "github.com/GuruOrGoru/guruvim",
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
			title:        "The GURU MMO",
			link:         "https://github.com/GuruOrGoru/mmo",
			description:  "A go based MMO game designed to support many users concurrently, game mechanics is kinda default with collecting spores and competing like agar.io",
			technologies: []string{"Go", "godot"},
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
			title:        "Interpreter",
			link:         "https://github.com/GuruOrGoru/making-an-interpreter-in-go",
			description:  "An interpreter purely based on nepali language, more commanly roman scripted nepali language with keywords like manau x=5; yadi satya {}, etc",
			technologies: []string{"Go", "language-architecture"},
		},
		{
			title:        "Lsp-Server",
			link:         "https://github.com/GuruOrGoru/lsp-server",
			description:  "A lsp server written in go under a 4 hr coding challenge, has basic lsp stuffs like completions, code actions and diagonostics.",
			technologies: []string{"Go", "VIM", "HTTP"},
		},
		{
			title:        "HTTP Server IN go",
			link:         "https://github.com/GuruOrGoru/http-server-in-go",
			description:  "A http server writtn in go, solely for learning purposes. Consists of chunked encoding, proper request handling in a modular approach",
			technologies: []string{"Go", "HTTP"},
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

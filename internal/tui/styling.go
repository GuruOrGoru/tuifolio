package tui

import "github.com/charmbracelet/lipgloss"

var (
	PointedLineStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#73f5e6")).Bold(true)
	NormalLineStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#c6c6c6"))
	SelectedLineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9d00")).Bold(true)

	LogoStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff"))
	CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#E95420"))

	HeaderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#E95420"))
)

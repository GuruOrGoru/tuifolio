package tui

import "github.com/charmbracelet/lipgloss"

var (
	pointedLineStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#73f5e6")).Bold(true)
	normalLineStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#c6c6c6"))
	selectedLineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9d00")).Bold(true)
)

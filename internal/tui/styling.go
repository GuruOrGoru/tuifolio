package tui

import "github.com/charmbracelet/lipgloss"

var (
	PointedLineStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#73f5e6")).Bold(true)
	NormalLineStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#c6c6c6"))
	SelectedLineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9d00")).Bold(true)

	LogoStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff"))
	CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#E95420"))

	HeaderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#E95420"))

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Margin(1, 0).
			Width(80)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Align(lipgloss.Center).
			MarginBottom(1)

	CenterOuterStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#E95420"))

	ContactItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252")).
				MarginBottom(1)

	ContactLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("11")).
				Bold(true)

	ContactValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("39"))

	DescriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("246")).
				Align(lipgloss.Center).
				MarginBottom(2)

	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246")).
			Italic(true).
			Align(lipgloss.Center)

	ProjectItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252")).
				MarginBottom(1)

	ProjectTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("11")).
				Bold(true).
				MarginBottom(1)

	ProjectDescStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("246")).
				MarginLeft(2).
				MarginBottom(1)

	TechStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			MarginLeft(2)

	SectionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")).
			Bold(true).
			MarginTop(1).
			MarginBottom(1)

	ItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			MarginLeft(2)
)

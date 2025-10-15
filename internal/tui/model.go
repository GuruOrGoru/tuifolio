package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Cursor        int
	Choices       []string
	Selected      map[int]struct{}
	Terminal      Terminal
	ShowLogo      bool
	CursorVisible bool
	activeTab     Tab
	viewport      viewport.Model

	Keys KeyMap
	Help help.Model
}

func NewModel() *Model {
	vp := viewport.New(0, 0)
	vp.MouseWheelEnabled = true
	vp.MouseWheelDelta = 3
	return &Model{
		Cursor:        0,
		Choices:       []string{"I am gay", "First Choice", "Second Choice", "Third Choice", "Fourth Choice", "Fifth Choice", "Sixth Choice", "Seventh Choice", "Eighth Choice", "Ninth Choice", "Tenth Choice", "Eleventh Choice", "Twelfth Choice", "Thirteenth Choice", "Fourteenth Choice", "Fifteenth Choice", "Sixteenth Choice", "Seventeenth Choice", "Eighteenth Choice", "Nineteenth Choice", "Twentieth Choice", "Twenty-first Choice", "Twenty-second Choice", "Twenty-third Choice", "Twenty-fourth Choice", "Twenty-fifth Choice", "Twenty-sixth Choice", "Twenty-seventh Choice", "Twenty-eighth Choice", "Twenty-ninth Choice", "Thirtieth Choice", "Thirty-first Choice", "Thirty-second Choice", "Thirty-third Choice", "Thirty-fourth Choice", "Thirty-fifth Choice", "Thirty-sixth Choice", "Thirty-seventh Choice", "Thirty-eighth Choice", "Thirty-ninth Choice", "Fortieth Choice", "Forty-first Choice", "Forty-second Choice", "Forty-third Choice", "Forty-fourth Choice", "Forty-fifth Choice", "Forty-sixth Choice", "Forty-seventh Choice", "Forty-eighth Choice", "Forty-ninth Choice", "Fiftieth Choice", "Fifty-first Choice", "Fifty-second Choice", "Fifty-third Choice", "Fifty-fourth Choice", "Fifty-fifth Choice", "Fifty-sixth Choice", "Fifty-seventh Choice", "Fifty-eighth Choice", "Fifty-ninth Choice", "Sixtieth Choice", "Sixty-first Choice", "Sixty-second Choice", "Sixty-third Choice", "Sixty-fourth Choice", "Sixty-fifth Choice", "Sixty-sixth Choice", "Sixty-seventh Choice", "Sixty-eighth Choice", "Sixty-ninth Choice", "Seventieth Choice", "Seventy-first Choice", "Seventy-second Choice", "Seventy-third Choice", "Seventy-fourth Choice", "Seventy-fifth Choice", "Seventy-sixth Choice", "Seventy-seventh Choice", "Seventy-eighth Choice", "Seventy-ninth Choice", "Eightieth Choice", "Eighty-first Choice", "Eighty-second Choice", "Eighty-third Choice", "Eighty-fourth Choice", "Eighty-fifth Choice", "Eighty-sixth Choice", "Eighty-seventh Choice", "Eighty-eighth Choice", "Eighty-ninth Choice", "Ninetieth Choice", "Ninety-first Choice", "Ninety-second Choice", "Ninety-third Choice", "Ninety-fourth Choice", "Ninety-fifth Choice", "Ninety-sixth Choice", "Ninety-seventh Choice", "Ninety-eighth Choice", "Ninety-ninth Choice", "One Hundredth Choice"},
		Selected:      make(map[int]struct{}),
		Terminal:      Terminal{Height: 0, Width: 0},
		ShowLogo:      true,
		CursorVisible: true,
		activeTab:     TabHome,
		viewport:      vp,
		Keys:          Keys,
		Help:          help.New(),
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(time.Millisecond*1500, func(time.Time) tea.Msg { return LogoDone{} }),
		tea.Tick(time.Millisecond*500, func(time.Time) tea.Msg { return CursorBlinkMsg{} }),
	)
}

func (m *Model) View() string {
	header := RenderHeader()

	tabStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(0, 1).
		Align(lipgloss.Center)

	activeTabStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("63")).
		Padding(0, 1).
		Align(lipgloss.Center).
		Bold(true)

	var sb strings.Builder
	sb.WriteString(header + "\n\n")
	for t := range TabCount {
		label := t.String()
		if t == m.activeTab {
			sb.WriteString(activeTabStyle.Render(label))
		} else {
			sb.WriteString(tabStyle.Render(label))
		}
		if t < TabCount-1 {
			sb.WriteString(" | ")
		}
	}
	sb.WriteString("\n\n")

	content := ""
	helpView := RenderHelp(m.Help, m.Keys)
	outerWidth := int(float64(m.Terminal.Width) * 0.8)
	outerHeight := int(float64(m.Terminal.Height) * 0.8)
	innerWidth := outerWidth - 4

	switch m.activeTab {
	case TabHome:
		content = RenderList(m.Cursor, m.Choices, m.Selected)
		content += "\n\n" + helpView
	case TabSettings:
		content = "Settings: adjust your preferences.\n\nThis is a longer text to demonstrate scrolling.\n" + strings.Repeat("Setting option number X.\n", 50)
	case TabHelp:
		content = "Help: press q to quit, ←/→ or 1/2/3 to switch tabs.\n\n" + strings.Repeat("Helpful information line Y.\n", 50)
	}

	// Set viewport content and size
	m.viewport.Width = innerWidth
	m.viewport.Height = outerHeight - 4 // account for header and tabs
	m.viewport.SetContent(content)

	if m.ShowLogo {
		return RenderLogo(m.Terminal, m.CursorVisible)
	}

	sb.WriteString(m.viewport.View())

	return CenterSquareWithContent(m.Terminal.Width, m.Terminal.Height, outerWidth, outerHeight, innerWidth, outerHeight, sb.String())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch messages := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(messages, m.Keys.Quit):
			return m, tea.Quit
		case key.Matches(messages, m.Keys.Up):
			if m.activeTab == TabHome {
				if m.Cursor > 0 {
					m.Cursor--
					cursorLineInView := m.Cursor - m.viewport.YOffset
					if cursorLineInView < 0 {
						m.viewport.ScrollUp(1)
					}
				}
			} else {
				m.viewport.ScrollUp(1)
			}
		case key.Matches(messages, m.Keys.Down):
			if m.activeTab == TabHome {
				if m.Cursor < len(m.Choices)-1 {
					m.Cursor++
					cursorLineInView := m.Cursor - m.viewport.YOffset
					if cursorLineInView >= m.viewport.Height {
						m.viewport.ScrollDown(1)
					}
				}
			} else {
				m.viewport.ScrollDown(1)
			}
		case key.Matches(messages, m.Keys.Select):
			if m.activeTab == TabHome {
				if _, ok := m.Selected[m.Cursor]; ok {
					delete(m.Selected, m.Cursor)
				} else {
					m.Selected = map[int]struct{}{m.Cursor: {}}
				}
			}
		case key.Matches(messages, m.Keys.TabBackward):
			if m.activeTab > 0 {
				m.activeTab--
			}
		case key.Matches(messages, m.Keys.TabForward):
			if m.activeTab < TabCount-1 {
				m.activeTab++
			}
		case key.Matches(messages, m.Keys.Tab1):
			m.activeTab = TabHome
		case key.Matches(messages, m.Keys.Tab2):
			m.activeTab = TabSettings
		case key.Matches(messages, m.Keys.Tab3):
			m.activeTab = TabHelp
		case key.Matches(messages, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		case messages.String() == "ctrl+d":
			m.viewport.PageDown()
		case messages.String() == "ctrl+u":
			m.viewport.PageUp()
		}
	case tea.MouseMsg:
		switch messages.Type {
		case tea.MouseWheelUp:
			m.viewport.ScrollUp(m.viewport.MouseWheelDelta)
		case tea.MouseWheelDown:
			m.viewport.ScrollDown(m.viewport.MouseWheelDelta)
		}
	case tea.WindowSizeMsg:
		m.Terminal.Width = messages.Width
		m.Terminal.Height = messages.Height
		m.Help.Width = messages.Width
		// Update viewport size
		outerWidth := int(float64(m.Terminal.Width) * 0.8)
		outerHeight := int(float64(m.Terminal.Height) * 0.8)
		innerWidth := outerWidth - 4
		m.viewport.Width = innerWidth
		m.viewport.Height = outerHeight - 4
	case LogoDone:
		m.ShowLogo = false
		return m, nil
	case CursorBlinkMsg:
		m.CursorVisible = !m.CursorVisible
		return m, tea.Tick(time.Millisecond*500, func(time.Time) tea.Msg { return CursorBlinkMsg{} })
	}
	return m, cmd
}

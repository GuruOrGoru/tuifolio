package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
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

	Keys     KeyMap
	Help     help.Model
	Username string
	Input    textinput.Model
	State    AppState

	// Interactive states for other tabs
	SkillsCursor    int
	SkillsExpanded  map[int]bool
	ShowFullDetails bool
	ContactCursor   int

	// Scroll states for each tab
	HomeScroll     int
	SkillsScroll   int
	ProjectsScroll int
	ContactScroll  int
}

func NewModel() *Model {
	vp := viewport.New(0, 0)
	vp.MouseWheelEnabled = true
	vp.MouseWheelDelta = 3
	ti := textinput.New()
	ti.Placeholder = "Enter your username"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 20
	return &Model{
		Cursor:          0,
		Choices:         []string{"Web Development", "Mobile Apps", "Backend Systems", "DevOps & Cloud", "Open Source", "Machine Learning", "UI/UX Design", "Database Design", "API Development", "Testing & QA"},
		Selected:        make(map[int]struct{}),
		Terminal:        Terminal{Height: 0, Width: 0},
		ShowLogo:        true,
		CursorVisible:   true,
		activeTab:       TabHome,
		viewport:        vp,
		Keys:            Keys,
		Help:            help.New(),
		Username:        "",
		Input:           ti,
		State:           StateLogo,
		SkillsCursor:    0,
		SkillsExpanded:  make(map[int]bool),
		ShowFullDetails: false,
		ContactCursor:   0,
		HomeScroll:      0,
		SkillsScroll:    0,
		ProjectsScroll:  0,
		ContactScroll:   0,
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(time.Millisecond*1500, func(time.Time) tea.Msg { return LogoDone{} }),
		tea.Tick(time.Millisecond*500, func(time.Time) tea.Msg { return CursorBlinkMsg{} }),
	)
}

func (m *Model) View() string {
	header := RenderHeader(m.Username)

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
	outerWidth := int(float64(m.Terminal.Width) * 0.8)
	outerHeight := int(float64(m.Terminal.Height) * 0.8)
	innerWidth := outerWidth - 4

	switch m.activeTab {
	case TabHome:
		content = RenderHomeTab(m.Cursor, m.Choices, m.Selected)
	case TabSkills:
		content = RenderSkillsTab(m.SkillsCursor, m.SkillsExpanded)
	case TabProjects:
		content = RenderProjectsTab(m.ShowFullDetails)
	case TabContact:
		content = RenderContactTab(m.ContactCursor)
	}

	m.viewport.Width = innerWidth
	m.viewport.Height = outerHeight - 4

	// Set scroll position for the active tab
	switch m.activeTab {
	case TabHome:
		m.viewport.YOffset = m.HomeScroll
	case TabSkills:
		m.viewport.YOffset = m.SkillsScroll
	case TabProjects:
		m.viewport.YOffset = m.ProjectsScroll
	case TabContact:
		m.viewport.YOffset = m.ContactScroll
	}

	m.viewport.SetContent(content)

	switch m.State {
	case StateLogo:
		return RenderLogo(m.Terminal, m.CursorVisible)
	case StateUsernameInput:
		return RenderUsernameInput(m.Terminal, m.Input)
	case StateMain:
	}

	sb.WriteString(m.viewport.View())
	sb.WriteString("\n\n")
	sb.WriteString(RenderHelp(m.Help, m.Keys))

	return CenterSquareWithContent(m.Terminal.Width, m.Terminal.Height, outerWidth, outerHeight, innerWidth, outerHeight, sb.String())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch messages := msg.(type) {
	case tea.KeyMsg:
		if m.State == StateUsernameInput {
			if messages.String() == "enter" {
				m.Username = m.Input.Value()
				m.State = StateMain
				return m, nil
			}
			var inputCmd tea.Cmd
			m.Input, inputCmd = m.Input.Update(messages)
			return m, inputCmd
		}

		switch {
		case key.Matches(messages, m.Keys.Quit):
			return m, tea.Quit
		case key.Matches(messages, m.Keys.Up):
			switch m.activeTab {
			case TabHome:
				if m.Cursor > 0 {
					m.Cursor--
					cursorLineInView := m.Cursor - m.viewport.YOffset
					if cursorLineInView < 0 {
						m.viewport.ScrollUp(1)
					}
				}
			case TabSkills:
				if m.SkillsCursor > 0 {
					m.SkillsCursor--
				}
			case TabContact:
				if m.ContactCursor > 0 {
					m.ContactCursor--
				}
			default:
				m.viewport.ScrollUp(1)
			}
		case key.Matches(messages, m.Keys.Down):
			switch m.activeTab {
			case TabHome:
				if m.Cursor < len(m.Choices)-1 {
					m.Cursor++
					cursorLineInView := m.Cursor - m.viewport.YOffset
					if cursorLineInView >= m.viewport.Height {
						m.viewport.ScrollDown(1)
					}
				}
			case TabSkills:
				if m.SkillsCursor < 4 { // 5 categories
					m.SkillsCursor++
				}
			case TabContact:
				if m.ContactCursor < 4 { // 5 contacts
					m.ContactCursor++
				}
			default:
				m.viewport.ScrollDown(1)
			}
		case key.Matches(messages, m.Keys.Select):
			switch m.activeTab {
			case TabHome:
				if _, ok := m.Selected[m.Cursor]; ok {
					delete(m.Selected, m.Cursor)
				} else {
					m.Selected = map[int]struct{}{m.Cursor: {}}
				}
			case TabSkills:
				if _, ok := m.SkillsExpanded[m.SkillsCursor]; ok {
					delete(m.SkillsExpanded, m.SkillsCursor)
				} else {
					m.SkillsExpanded[m.SkillsCursor] = true
				}
			case TabProjects:
				m.ShowFullDetails = !m.ShowFullDetails
			case TabContact:
				// Maybe do nothing or show a message, but for now, just highlight
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
			m.activeTab = TabSkills
		case key.Matches(messages, m.Keys.Tab3):
			m.activeTab = TabProjects
		case key.Matches(messages, m.Keys.Tab4):
			m.activeTab = TabContact
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
		outerWidth := int(float64(m.Terminal.Width) * 0.8)
		outerHeight := int(float64(m.Terminal.Height) * 0.8)
		innerWidth := outerWidth - 4
		m.viewport.Width = innerWidth
		m.viewport.Height = outerHeight - 4
	case LogoDone:
		m.State = StateUsernameInput
		return m, nil
	case CursorBlinkMsg:
		m.CursorVisible = !m.CursorVisible
		return m, tea.Tick(time.Millisecond*500, func(time.Time) tea.Msg { return CursorBlinkMsg{} })
	}

	// Update scroll position for the active tab
	switch m.activeTab {
	case TabHome:
		m.HomeScroll = m.viewport.YOffset
	case TabSkills:
		m.SkillsScroll = m.viewport.YOffset
	case TabProjects:
		m.ProjectsScroll = m.viewport.YOffset
	case TabContact:
		m.ContactScroll = m.viewport.YOffset
	}

	return m, cmd
}

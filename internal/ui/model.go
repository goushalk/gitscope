package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/goushalk/gitscope/internal/logic"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// Model implements the Bubble Tea model used for the activity table screen.
type Model struct {
	Table    table.Model
	Username string
}

// NewModel wires the table view with the username shown in the banner.
func NewModel(t table.Model, username string) Model {
	return Model{
		Table:    t,
		Username: username,
	}
}

// Init satisfies the Bubble Tea interface. No startup command is required.
func (m Model) Init() tea.Cmd { return nil }

// Update handles input and delegates table navigation/state updates.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

// View renders the banner and the styled table.
func (m Model) View() string {
	banner := logic.Banner(m.Username)
	return banner + "\n" + baseStyle.Render(m.Table.View()) + "\n"
}

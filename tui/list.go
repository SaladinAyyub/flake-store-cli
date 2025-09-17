package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SaladinAyyub/flake-store-cli/internal/models"
)

// flakeItem implements list.Item interface for Bubble Tea
type flakeItem struct {
	Name, Desc string
}

func (f flakeItem) Title() string { return f.Name }

func (f flakeItem) Description() string { return f.Desc }

func (f flakeItem) FilterValue() string { return f.Name }

// model holds the Bubble Tea list
type model struct {
	list list.Model
}

// Init is part of Bubble Tea interface
func (m model) Init() tea.Cmd {
	return nil
}

// Update is part of Bubble Tea interface
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			item := m.list.SelectedItem().(flakeItem)
			fmt.Printf("\nSelected flake: %s\n", item.Name)
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the Bubble Tea list
func (m model) View() string {
	return m.list.View()
}

func List(flakes []models.Flake) error {
	items := []list.Item{}
	for _, f := range flakes {
		items = append(items, flakeItem{Name: f.Name, Desc: f.Description})
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = true

	// Provide explicit width and height for proper rendering
	l := list.New(items, delegate, 50, 15)
	l.Title = "Select a Flake (Enter to install, q to quit)"

	p := tea.NewProgram(model{list: l})
	_, err := p.Run() // Run returns (Model, error)
	return err
}

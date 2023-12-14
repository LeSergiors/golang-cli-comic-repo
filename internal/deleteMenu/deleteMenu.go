package deletemenu

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type DelMenu struct{}

func (d DelMenu) Load() tea.Model {
	return menu{
		options: []menuItem{
			{
				text:    "Delete Publishers",
				onPress: func() tea.Msg { return struct{}{} },
			},
			{
				text:    "Delete Comics",
				onPress: func() tea.Msg { return struct{}{} },
			},
		},
	}
}

type menuItem struct {
	text    string
	onPress func() tea.Msg
}

type menu struct {
	options       []menuItem
	selectedIndex int
}

func (m menu) Init() tea.Cmd {
	return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down", "up", "right", "left":
			return m.moveCursor(msg.(tea.KeyMsg)), nil
		case "enter", "return":
			return m, m.options[m.selectedIndex].onPress
		}
	}

	return m, nil
}

func (m menu) moveCursor(msg tea.KeyMsg) menu {
	switch msg.String() {
	case "up", "left":
		m.selectedIndex--
	case "down", "right":
		m.selectedIndex++
	default:

	}

	optCount := len(m.options)
	m.selectedIndex = (m.selectedIndex + optCount) % optCount
	return m
}

func (m menu) View() string {
	var options []string
	for i, o := range m.options {
		if i == m.selectedIndex {
			options = append(options, fmt.Sprintf("-> %s", o.text))
		} else {
			options = append(options, fmt.Sprintf("   %s", o.text))
		}
	}

	return fmt.Sprintf(`%s 
	
	Press enter/return to select a list item, arrow kwys to move, or Ctrl+c to exit.`, strings.Join(options, "\n"))
}

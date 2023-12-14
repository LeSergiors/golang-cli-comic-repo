package deletemenu

import (
	"comic-cli/internal/load"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type toggleBackMenu struct{}

type DelMenu struct{}

var backMenu load.LoadMenu

func (d DelMenu) Load(l load.LoadMenu) tea.Model {
	me := menu{
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

	s := "Exit"
	if l != nil {
		s = "Go Back."
		backMenu = l
	}

	me.options = append(me.options, menuItem{text: s, onPress: func() tea.Msg { return toggleBackMenu{} }})

	return me

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
	case toggleBackMenu:
		if backMenu != nil {
			return m.goBack(), nil
		} else {
			return m, tea.Quit
		}
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

func (m menu) goBack() tea.Model {
	d := backMenu.Load(DelMenu{})
	return d
}

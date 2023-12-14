package load

import tea "github.com/charmbracelet/bubbletea"

type LoadMenu interface {
	Load() tea.Model
}

func LoadNewMenu(l LoadMenu) tea.Model {
	return l.Load()
}

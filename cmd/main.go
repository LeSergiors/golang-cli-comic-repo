package main

import (
	"fmt"
	"os"

	"comic-cli/internal/initMenu"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	initModel := initmenu.InitMenu{}

	p := tea.NewProgram(initModel.Load(nil))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas! There's been an error %v", err)
		os.Exit(1)
	}
}

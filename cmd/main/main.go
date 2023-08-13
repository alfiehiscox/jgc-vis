package main

import (
	"fmt"
	"os"

	"github.com/alfiehiscox/jgc-vis/pkg/view"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := view.Model{}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

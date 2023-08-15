package main

import (
	"fmt"
	"os"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
	"github.com/alfiehiscox/jgc-vis/pkg/view"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := view.Model{
		Bus:   make(chan []parser.GCLog),
		Table: view.NewTable(),
	}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

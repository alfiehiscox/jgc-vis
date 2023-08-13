package view

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Table table.Model
}

// Base Style
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// Column Names
var cols []table.Column = []table.Column{
	{Title: "Timestamp", Width: 25},
	{Title: "Type", Width: 7},
	{Title: "Reason", Width: 20},
	{Title: "Before Size", Width: 12},
	{Title: "After Size", Width: 10},
	{Title: "Total Size", Width: 10},
}

// Fetch Messages
type logMsg []parser.GCLog
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

// Fetch Command
func fetchLogs() tea.Msg {
	var logs []parser.GCLog

	bs, err := os.ReadFile("/Users/alfiehiscox/Code/alfiehiscox/go/jgc-vis/resource/java8-test-timestamp.log")
	if err != nil {
		return errMsg{err}
	}

	fs := string(bs)
	ss := strings.Split(fs, "\n")

	for _, l := range ss {
		parser := parser.NewParser(l)
		log, err := parser.Parse()
		if err != nil {
			return errMsg{err}
		}
		logs = append(logs, *log)
	}

	return logMsg(logs)
}

func (m Model) Init() tea.Cmd {
	return fetchLogs
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case logMsg:
		// Here 'msg' is all of the GCLog objects from the file
		logs := []parser.GCLog(msg)
		// Create a new table with all the logs
		m.Table = table.New(
			table.WithColumns(cols),
			table.WithRows(logsToRows(logs)),
			table.WithFocused(true),
			table.WithHeight(20),
		)
		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		m.Table.SetStyles(s)
	case errMsg:
		fmt.Println(msg)
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			// Ctrl+c will quit
			return m, tea.Quit
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.Table.View()) + "\n"
}

func logsToRows(logs []parser.GCLog) []table.Row {
	var rows []table.Row
	for _, log := range logs {
		rows = append(rows, []string{
			log.Timestamp.Format(time.ANSIC),
			log.Type,
			log.Reason,
			fmt.Sprint(log.MainEvent.BeforeSize),
			fmt.Sprint(log.MainEvent.AfterSize),
			fmt.Sprint(log.MainEvent.TotalSize),
		})
	}
	return rows
}

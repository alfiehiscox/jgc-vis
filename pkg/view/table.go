package view

import (
	"fmt"
	"time"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Styles
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
	tableHeight = 15

	// Timers
	wait = time.Second * 5
)

// Column Names
var cols []table.Column = []table.Column{
	{Title: "Index", Width: 0},
	{Title: "Timestamp", Width: 25},
	{Title: "Type", Width: 7},
	{Title: "Reason", Width: 20},
	{Title: "Before Size", Width: 12},
	{Title: "After Size", Width: 10},
	{Title: "Total Size", Width: 10},
}

// Using https://github.com/charmbracelet/bubbletea/blob/master/examples/realtime/main.go as
// an example.

// Fetch Messages
type logMsg []parser.GCLog
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

// Sends the Logs through on a given channel
func PollLogsForData(c chan []parser.GCLog) tea.Cmd {
	return func() tea.Msg {
		for {
			// Fetch
			logs, err := parser.FetchLogs(FileName)
			if err != nil {
				return errMsg{err}
			}
			c <- logs

			// Sleep
			time.Sleep(wait)
		}
	}
}

// Wait on the channel and return a logMsg when we do.
func WaitForLogs(c chan []parser.GCLog) tea.Cmd {
	return func() tea.Msg {
		return logMsg(<-c) // Waits for logs
	}
}

func logsToRows(logs []parser.GCLog) []table.Row {
	var rows []table.Row
	for i, log := range logs {
		rows = append(rows, []string{
			fmt.Sprint(i),
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

func NewTable() table.Model {
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

	table := table.New(
		table.WithColumns(cols),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(tableHeight),
	)

	table.SetStyles(s)

	return table
}

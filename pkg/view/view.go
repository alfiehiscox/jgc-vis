package view

import (
	"strconv"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	FileName string

	Bus      chan []parser.GCLog
	Table    table.Model
	Logs     []parser.GCLog
	Selected parser.GCLog

	Width  int
	Height int

	Error error
}

// Start the timers off going
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		PollLogsForData(m.Bus, m.FileName),
		WaitForLogs(m.Bus),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case logMsg:
		// Here 'msg' is all of the GCLog objects from the file
		logs := []parser.GCLog(msg)
		m.Table.SetRows(logsToRows(logs))
		m.Logs = logs
	case errMsg:
		tea.Println(msg)
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, err := strconv.Atoi(m.Table.SelectedRow()[0])
			if err != nil {
				m.Error = err
			}
			m.Selected = m.Logs[i]
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, tea.Batch(cmd, WaitForLogs(m.Bus))
}

func (m Model) View() string {
	table := baseStyle.Render(m.Table.View())
	details := logToDetails(m.Selected)
	timeGraph := NewTimeGraph(m.Logs)
	sizeGraph := NewSizeGraph(m.Logs)

	var tableAndDetails string
	if m.Width > tableWidth+detailsWidth {
		tableAndDetails = lipgloss.JoinHorizontal(lipgloss.Top, table, details)
	} else {
		tableAndDetails = lipgloss.JoinVertical(lipgloss.Top, table, details)
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		tableAndDetails,
		lipgloss.JoinHorizontal(lipgloss.Top, sizeGraph, timeGraph),
	) + "\n"
}

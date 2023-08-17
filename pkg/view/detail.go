package view

import (
	"fmt"
	"time"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
	"github.com/charmbracelet/lipgloss"
)

var (
	detailsWidth = 36
	detailsStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Height(17).
			Width(detailsWidth)
	titleStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Width(detailsWidth)
)

func logToDetails(log parser.GCLog) string {
	return detailsStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		titleStyle.Render("Details:"),
		"Timestamp: "+log.Timestamp.Format(time.ANSIC),
		"Type: "+log.Type,
		"Reason: "+log.Reason,
		"Time: "+log.Time+" secs",
		"Old Generation:",
		"  Pre-GC: "+fmt.Sprint(log.MainEvent.BeforeSize)+"K",
		"  Post-GC: "+fmt.Sprint(log.MainEvent.AfterSize)+"K",
		"  Total: "+fmt.Sprint(log.MainEvent.TotalSize)+"K",
		"Young Generation:",
		getYoungGenerationDetails(log),
	))
}

func getYoungGenerationDetails(log parser.GCLog) string {
	var event parser.GCEvent
	for _, g := range log.GenEvents {
		if g.Type == "PSYoungGen" {
			event = g.Event
		}
	}
	strings := "  Pre-GC: " + fmt.Sprint(event.BeforeSize) + "K" + "\n"
	strings += "  Post-GC: " + fmt.Sprint(event.AfterSize) + "K" + "\n"
	strings += "  Total: " + fmt.Sprint(event.TotalSize) + "K" + "\n"
	return strings
}

// func logToGenerations(log parser.GCLog) string {
// 	strings := ""
// 	for _, g := range log.GenEvents {
// 		strings += "  " + g.Type + "\n"
// 		strings += "  - Pre-GC: " + fmt.Sprint(g.Event.BeforeSize) + "\n"
// 		strings += "  - Post-GC: " + fmt.Sprint(g.Event.AfterSize) + "\n"
// 		strings += "  - Total: " + fmt.Sprint(g.Event.TotalSize) + "\n"
// 	}
// 	return strings
// }

package view

import (
	"strconv"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
	"github.com/guptarohit/asciigraph"
)

var (
	// Styles
	graphBox = baseStyle

	// Config
	maxRecord = 40
)

func NewTimeGraph(logs []parser.GCLog) string {
	if len(logs) <= 0 {
		return ""
	}

	plots := logsToTimePlot(logs)
	graph := asciigraph.Plot(
		plots,
		asciigraph.Height(10),
		asciigraph.Caption("GC Time"),
	)
	styled := graphBox.Render(graph)
	return styled
}

func NewSizeGraph(logs []parser.GCLog) string {
	if len(logs) <= 0 {
		return ""
	}

	plots := logsToTotalSizePlot(logs)
	graph := asciigraph.Plot(
		plots,
		asciigraph.Height(10),
		asciigraph.Caption("Total GC Size"),
	)
	styled := graphBox.Render(graph)
	return styled
}

func logsToTotalSizePlot(logs []parser.GCLog) []float64 {
	var plots []float64
	for i, log := range logs {
		if i > maxRecord-1 {
			break
		}
		plots = append(plots, float64(log.MainEvent.TotalSize))
	}
	return plots
}

func logsToTimePlot(logs []parser.GCLog) []float64 {
	var plots []float64
	for i, log := range logs {
		if i > maxRecord-1 {
			break
		}
		time, _ := strconv.ParseFloat(log.Time, 64)
		plots = append(plots, time)
	}
	return plots
}

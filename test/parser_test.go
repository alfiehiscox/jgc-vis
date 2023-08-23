package main

import (
	"testing"
	"time"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
)

func TestJava8MinorGCLogToGCLog(t *testing.T) {
	log := "[GC (Allocation Failure) [PSYoungGen: 8704K->992K(9728K)] 8704K->3776K(31744K), 0.0073015 secs] [Times: user=0.02 sys=0.00, real=0.01 secs]"

	expected := parser.GCLog{
		Type:   "GC",
		Reason: "Allocation Failure",
		MainEvent: parser.GCEvent{
			BeforeSize: 8704,
			AfterSize:  3776,
			TotalSize:  31744,
		},
		Time: "0.0073015",
		GenEvents: []struct {
			Type  string
			Event parser.GCEvent
		}{
			{
				Type: "PSYoungGen",
				Event: parser.GCEvent{
					BeforeSize: 8704,
					AfterSize:  992,
					TotalSize:  9728,
				},
			},
		},
	}

	parser := parser.NewParser(log)
	actual, err := parser.Parse()
	if err != nil {
		t.Fatalf("error from parsing: %s", err)
	}

	ok := checkGCLog(*actual, expected)

	if !ok {
		t.Errorf("expected: %v", expected)
		t.Errorf("actual  : %v", *actual)
		t.Fatalf("actual and exepcted did not match.")
	}
}

func TestJava8MajorGCLogToGCLog(t *testing.T) {
	log := "[Full GC (Ergonomics) [PSYoungGen: 57344K->0K(113664K)] [ParOldGen: 337435K->261246K(339968K)] 394779K->261246K(453632K), [Metaspace: 2866K->2866K(1056768K)], 0.3415608 secs] [Times: user=1.26 sys=0.00, real=0.35 secs]"

	expected := parser.GCLog{
		Type:   "Full GC",
		Reason: "Ergonomics",
		MainEvent: parser.GCEvent{
			BeforeSize: 394779,
			AfterSize:  261246,
			TotalSize:  453632,
		},
		Time: "0.3415608",
		GenEvents: []struct {
			Type  string
			Event parser.GCEvent
		}{
			{
				Type: "PSYoungGen",
				Event: parser.GCEvent{
					BeforeSize: 57344,
					AfterSize:  0,
					TotalSize:  113664,
				},
			},
			{
				Type: "ParOldGen",
				Event: parser.GCEvent{
					BeforeSize: 337435,
					AfterSize:  261246,
					TotalSize:  339968,
				},
			},
			{
				Type: "Metaspace",
				Event: parser.GCEvent{
					BeforeSize: 2866,
					AfterSize:  2866,
					TotalSize:  1056768,
				},
			},
		},
	}

	parser := parser.NewParser(log)
	actual, err := parser.Parse()
	if err != nil {
		t.Fatalf("error from parsing: %s", err)
	}

	ok := checkGCLog(*actual, expected)

	if !ok {
		t.Errorf("expected: %v", expected)
		t.Errorf("actual  : %v", *actual)
		t.Fatalf("actual and exepcted did not match.")
	}
}

func TestJava8SystemGCLogToGCLog(t *testing.T) {
	log := "[GC (System.gc()) [PSYoungGen: 96578K->40640K(113664K)] 369011K->368584K(453632K), 0.2221062 secs] [Times: user=0.32 sys=0.28, real=0.22 secs]"

	expected := parser.GCLog{
		Type:   "GC",
		Reason: "System.gc",
		MainEvent: parser.GCEvent{
			BeforeSize: 369011,
			AfterSize:  368584,
			TotalSize:  453632,
		},
		Time: "0.2221062",
		GenEvents: []struct {
			Type  string
			Event parser.GCEvent
		}{
			{
				Type: "PSYoungGen",
				Event: parser.GCEvent{
					BeforeSize: 96578,
					AfterSize:  40640,
					TotalSize:  113664,
				},
			},
		},
	}

	parser := parser.NewParser(log)
	actual, err := parser.Parse()

	if err != nil {
		t.Fatalf("error from parsing: %s", err)
	}

	ok := checkGCLog(*actual, expected)

	if !ok {
		t.Errorf("expected: %v", expected)
		t.Errorf("actual  : %v", *actual)
		t.Fatalf("actual and exepcted did not match.")
	}
}

func TestJava8TimestampToGCLog(t *testing.T) {
	log := "2023-08-10T11:09:31.795+0000: [Full GC (Ergonomics) [PSYoungGen: 57344K->0K(113664K)] [ParOldGen: 337435K->261246K(339968K)] 394779K->261246K(453632K), [Metaspace: 2866K->2866K(1056768K)], 0.3415608 secs] [Times: user=1.26 sys=0.00, real=0.35 secs]"

	timestamp, err := time.Parse("2006-01-02T15:04:05-0700", "2023-08-10T11:09:31.795+0000")
	if err != nil {
		t.Fatal("Error parsing date format for test")
	}

	expected := parser.GCLog{
		Timestamp: timestamp,
		Type:      "Full GC",
		Reason:    "Ergonomics",
		MainEvent: parser.GCEvent{
			BeforeSize: 394779,
			AfterSize:  261246,
			TotalSize:  453632,
		},
		Time: "0.3415608",
		GenEvents: []struct {
			Type  string
			Event parser.GCEvent
		}{
			{
				Type: "PSYoungGen",
				Event: parser.GCEvent{
					BeforeSize: 57344,
					AfterSize:  0,
					TotalSize:  113664,
				},
			},
			{
				Type: "ParOldGen",
				Event: parser.GCEvent{
					BeforeSize: 337435,
					AfterSize:  261246,
					TotalSize:  339968,
				},
			},
			{
				Type: "Metaspace",
				Event: parser.GCEvent{
					BeforeSize: 2866,
					AfterSize:  2866,
					TotalSize:  1056768,
				},
			},
		},
	}

	parser := parser.NewParser(log)
	actual, err := parser.Parse()

	if err != nil {
		t.Fatalf("error from parsing: %s", err)
	}

	ok := checkGCLogWithTimestamp(*actual, expected)

	if !ok {
		t.Errorf("expected: %v", expected)
		t.Errorf("actual  : %v", *actual)
		t.Fatalf("actual and exepcted did not match.")
	}
}

func TestJava8TimeSinceStart(t *testing.T) {
	log := "2023-08-10T11:09:31.795+0000: [Full GC (Ergonomics) [PSYoungGen: 57344K->0K(113664K)] [ParOldGen: 337435K->261246K(339968K)] 394779K->261246K(453632K), [Metaspace: 2866K->2866K(1056768K)], 0.3415608 secs] [Times: user=1.26 sys=0.00, real=0.35 secs]"
}

func checkGCLogWithTimestamp(actual, expected parser.GCLog) bool {
	if actual.Timestamp != expected.Timestamp {
		return false
	}

	return checkGCLog(actual, expected)
}

func checkGCLog(actual, expected parser.GCLog) bool {
	if actual.Type != expected.Type {
		return false
	}

	if actual.Reason != expected.Reason {
		return false
	}

	if actual.Time != expected.Time {
		return false
	}

	if len(actual.GenEvents) != len(expected.GenEvents) {
		return false
	}

	// The events should be ordered correctly in the slice
	for i, a := range actual.GenEvents {
		e := expected.GenEvents[i]

		if a.Type != e.Type {
			return false
		}

		ok := checkGCEvent(a.Event, e.Event)
		if !ok {
			return false
		}
	}

	return true
}

func checkGCEvent(actual, expected parser.GCEvent) bool {
	if actual.BeforeSize != expected.BeforeSize {
		return false
	}
	if actual.AfterSize != expected.AfterSize {
		return false
	}
	if actual.TotalSize != expected.TotalSize {
		return false
	}
	return true
}

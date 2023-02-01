package aggregator

import (
	"fmt"
	"time"
)

// From https://cs.opensource.google/go/go/+/refs/tags/go1.19.5:src/cmd/test2json/main.go
type TestEvent struct {
	Time    time.Time // encodes as an RFC3339-format string
	Action  string
	Package string
	Test    string
	Elapsed float64 // seconds
	Output  string
}

// From https://cs.opensource.google/go/go/+/refs/tags/go1.19.5:src/cmd/test2json/main.go
const (
	TestEventAction_run    = "run"
	TestEventAction_pause  = "pause"
	TestEventAction_cont   = "cont"
	TestEventAction_pass   = "pass"
	TestEventAction_bench  = "bench"
	TestEventAction_fail   = "fail"
	TestEventAction_output = "output"
	TestEventAction_skip   = "skip"
)

type TestSummary struct {
	Package        string
	Test           string
	Events         []*TestEvent
	Failed         bool
	Skipped        bool
	OutputLines    []string
	BenchmarkLines []string
}

type TestEventAggregator struct {
	TestEvents []*TestEvent
	TestMap    map[string]*TestSummary
	Failures   int32
	Skips      int32
	Passes     int32
}

func (t *TestEventAggregator) Add(te *TestEvent) {
	t.TestEvents = append(t.TestEvents, te)
	ts, ok := t.TestMap[fmt.Sprintf("%s.%s", te.Package, te.Test)]
	if !ok {
		ts = &TestSummary{Package: te.Package, Test: te.Test, Events: make([]*TestEvent, 0), OutputLines: make([]string, 0), BenchmarkLines: make([]string, 0)}
		if t.TestMap == nil {
			t.TestMap = make(map[string]*TestSummary)
		}
		t.TestMap[fmt.Sprintf("%s.%s", te.Package, te.Test)] = ts
	}
	ts.Add(te)
	switch te.Action {
	case TestEventAction_fail:
		t.Failures++
	case TestEventAction_skip:
		t.Skips++
	case TestEventAction_pass:
		t.Passes++
	}
}

func (t *TestSummary) Add(te *TestEvent) {
	t.Events = append(t.Events, te)
	switch te.Action {
	case TestEventAction_skip:
		t.Skipped = true
	case TestEventAction_fail:
		t.Failed = true
	case TestEventAction_output:
		t.OutputLines = append(t.OutputLines, te.Output)
	case TestEventAction_bench:
		t.BenchmarkLines = append(t.BenchmarkLines, te.Output)
	}
}

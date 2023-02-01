package aggregator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TestSummary_Add_Output(t *testing.T) {
	ts := &TestSummary{}
	ts.Add(&TestEvent{Action: TestEventAction_output, Output: "A"})
	ts.Add(&TestEvent{Action: TestEventAction_output, Output: "B"})

	assert.Len(t, ts.OutputLines, 2)
	assert.Equal(t, "A", ts.OutputLines[0])
	assert.Equal(t, "B", ts.OutputLines[1])
}

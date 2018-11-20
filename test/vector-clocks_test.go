package test

import (
	"concurrency-assignment2/vectorclocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorClock(t *testing.T) {
	p := vectorclocks.MakeProcesses(3)
	vectorclocks.IncrEvent(p[0], "A")
	vectorclocks.SendMessage(p[0], p[1], "B", "F")
	vectorclocks.IncrEvent(p[0], "C")
	vectorclocks.SendMessage(p[2], p[1], "H", "K")
	vectorclocks.SendMessage(p[0], p[2], "E", "J")
	vectorclocks.SendMessage(p[1], p[0], "G", "D")
	vectorclocks.IncrEvent(p[2], "I")
	assert.True(t, vectorclocks.IsConcurrent("C", "F"))
	assert.True(t, vectorclocks.IsConcurrent("C", "H"))
	assert.True(t, vectorclocks.IsBefore("A", "B"))
	assert.True(t, vectorclocks.IsBefore("B", "F"))
	assert.True(t, vectorclocks.IsBefore("A", "F"))
	assert.True(t, vectorclocks.IsBefore("H", "K"))
	assert.True(t, vectorclocks.IsBefore("E", "J"))
	assert.True(t, vectorclocks.IsBefore("G", "D"))
}

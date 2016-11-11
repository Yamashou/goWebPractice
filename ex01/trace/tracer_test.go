package trace

import (
	"bytes"
	"testing"
)

// TestNew tests the tracing behaviour.
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer1 := New(&buf)
	if tracer1 == nil {
		t.Error("Return from New should not be nil")
	}
	tracer1.Trace("Hello trace package.")
	if buf.String() != "Hello trace package.\n" {
		t.Errorf("Trace should not write '%s'.", buf.String())
	}

}
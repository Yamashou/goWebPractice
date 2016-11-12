package trace

import (
	"fmt"
	"io"
)

//Tracer <1>,<2>を分けるために使うインターフェース
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *tracer) Trace(a ...interface{}) { //<1>
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
func (t *nilTracer) Trace(a ...interface{}) {} //<2>

func Off() Tracer {
	return &nilTracer{}
}

package main

import (
	"fmt"
	"github.com/metacosm/odo-event-api/odo/api/events"
	"github.com/metacosm/odo-event-api/odo/api/events/plugin"
)

type tracer struct{}

func (t tracer) OnEvent(event events.Event) error {
	fmt.Printf("got event %v\n", event)
	return nil
}

func (t tracer) OnAbort(abortError events.EventCausedAbortError) {
	fmt.Printf("abort: %v\n", abortError)
}

func (t tracer) Name() string {
	return "tracer"
}

func (t tracer) String() string {
	return t.Name()
}

var Listener tracer

func main() {
	plug := plugin.NewPlugin(Listener)
	plug.Serve()
}

package main

import (
	"fmt"
	"github.com/redhat-developer/odo/pkg/odo/events"
)

type tracer struct{}

func (t tracer) OnEvent(event events.Event) error {
	fmt.Printf("got event %v\n", event)
	return nil
}

func (t tracer) OnAbort(abortError *events.EventCausedAbortError) {
	fmt.Printf("abort: %v\n", abortError)
}

func (t tracer) Name() string {
	return "tracer"
}

var Listener tracer

package main

import (
	"github.com/metacosm/odo-event-api/odo/api/events"
	"github.com/metacosm/odo-event-api/odo/api/events/plugin"
	"log"
	"os"
	"path/filepath"
)

var logger *log.Logger

type tracer struct{}

func (t tracer) OnEvent(event events.Event) error {
	logger.Printf("got event %v\n", event)
	return nil
}

func (t tracer) OnAbort(abortError events.EventCausedAbortError) {
	logger.Printf("abort: %v\n", abortError)
}

func (t tracer) Name() string {
	return "tracer"
}

func (t tracer) String() string {
	return t.Name()
}

var Listener tracer

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := filepath.Join(dir, "tracer.log")
	logFile, _ := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()
	logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)

	plug := plugin.NewPlugin(Listener)
	plug.Serve()
}

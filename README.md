# odo-tracer-listener

Build using:
```bash
go build -buildmode=plugin -o tracer.so listener.go
```

Then:
```bash
mkdir -p ~/.odo/plugins
cp tracer.so ~/.odo/plugins
```

Note that for the plugin to work it needs to be run with the exact `odo` version that was used to build it. Since this is 
obviously an issue, we would need to separate the API required by the plugin into a different project so that it can be reused
so that plugins can be built against a given API instead of requiring them to be built against the `odo` binary.
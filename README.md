## Requirements

* [Bazel](https://docs.bazel.build/versions/master/install.html)
* [go](https://golang.org/doc/install)

## Running stockbuddy
This is a 2-step process involving 2 separate binaries.

### 1. Start the quote service:

```sh
bazel run //quote_service
```

This starts an RPC server on localhost:50051.

### 2. Run the analyzer.

```sh
bazel run //analysis
```

This binary connects to the RPC service started in step 1, runs the analyzers
and then dies.

## Compiling stockbuddy
Stockbuddy currently runs on raspberry Pi's. However, bazel is currently
unsupported on Arm, so we must compile for the Pi and scp the binary,
which can be done via:

```sh
bazel build //quote_service --platforms=//:rpi_linux-arm
```

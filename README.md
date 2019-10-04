## Requirements

* [Bazel](https://docs.bazel.build/versions/master/install.html).
* [go](https://golang.org/doc/install?download=go1.13.linux-amd64.tar.gz)

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

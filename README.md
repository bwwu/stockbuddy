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
STOCKBUDDY_PASSWORD=<myPassword> bazel run //analysis \
  --action_env=STOCKBUDDY_PASSWORD
  -- \
  --use_watchlist=<relative_path> (OPTIONAL) \
  --mail_to=<comma_separated_list>  \
  --nomail (OPTIONAL) \
```

You must set the env var `$STOCKBUDDY_PASSWORD` to the email user credentials.

This binary connects to the RPC service started in step 1, runs the analyzers
and then dies.

## Compiling stockbuddy
Stockbuddy currently runs on raspberry Pi's. However, bazel is currently
unsupported on Arm, so we must compile for the Pi and scp the binary,
which can be done via:

```sh
bazel build //quote_service --platforms=//:rpi_linux-arm
```

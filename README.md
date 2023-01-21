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

## Using docker (new)

First install [Docker](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository).

```sh
bazel run //quote_service:image
```
To build an image for loading into docker on a raspi:

```
# Might instead need to add 'goarch = "arm"' to the "go_image" rule
bazel build //quote_service:image.tar --platforms=//:rpi_linux-arm
# scp the image.tar to the Pi
docker load -i bazel-bin/quote_service/image.tar
docker run --restart=always -p 50051:50051 <IMAGE_ID>
```


### References

* [bazelbuild/docker](https://github.com/bazelbuild/rules_docker/blob/master/README.md)

* https://github.com/bazelbuild/rules_go/blob/master/go/core.rst#cross-compilation
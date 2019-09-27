## Requirements

* python3. Install with:
* install [go](https://golang.org/doc/install?download=go1.13.linux-amd64.tar.gz)

```sh
sudo apt-get install python3
sudo apt-get install python3-pip
```

## Install notes
Install grpc python deps.
```sh
python3 -m pip install grpcio
python3 -m pip install grpcio-tools
```
Install [Bazel](https://docs.bazel.build/versions/master/install.html).

## Re-compile protos
```sh
 python3 -m grpc_tools.protoc \
	-I $SRC_DIR \
	--python_out=$DST_DIR \
	--grpc_python_out=$DST_DIR $SRC_DIR/quote.proto \
```

## Running stockbuddy
This is a 2-step process.
### 1. Start the quote service:

```sh
python3 quote_service/main.py
```

This starts an RPC server on localhost:50051.

### 2. Run the analyzer.

```sh
bazel run //analysis:main
```

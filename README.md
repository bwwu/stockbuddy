## Requirements

* python3. Install with:

```sh
sudo apt-get install python3
sudo apt-get install python3-pip
```

* virtualenv. Install with:

```sh
sudo apt-get install virtualenv
```



## Notes
```sh
python3 -m pip install grpcio
python3 -m pip install grpcio-tools
```

## Run
```sh
 python3 -m grpc_tools.protoc \
	-I $SRC_DIR \
	--python_out=$DST_DIR \
	--grpc_python_out=$DST_DIR $SRC_DIR/quote.proto \
```

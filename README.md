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



## Install notes
```sh
python3 -m pip install grpcio
python3 -m pip install grpcio-tools
```

## Re-compile protos
```sh
 python3 -m grpc_tools.protoc \
	-I $SRC_DIR \
	--python_out=$DST_DIR \
	--grpc_python_out=$DST_DIR $SRC_DIR/quote.proto \
```

# Notes/Next steps
Rough roadmap of what needs to happen next. Order is not necessarily important
due to loose coupling of micro-services:

### 1. Persist quote data
Figure out DB requirements for persisting data. Determine whether it is 
important to persist quote as protos. If so, probably need to use 
[leveldb](https://github.com/google/leveldb). Otherwise, may use a SQL DB.

If leveldb is chosen, need to write some sort of daily cronjob in C++ to
interact with Python stubby server for fetching quote data, then persist new
data to db. 

### 2. Analysis tools
Design framework for adding analyses on top of data. Should be easily
extensible to adopt new analyses. Need not wait on persistent storage, since
API calls to fetch quote data is cheap. Most of the application code should
exist in this layer, so choice of framework is important.It should have the
following properties:

* Strongly & statically typed.
* Quick & easy to prototype new features. 
* Well supported library which applies numerical methods.
* Simple library for integrating with notifications (email or SMS)
* Interoptibility with protobufs.

Go, C++ and Java seem to meet this criteria.

### 3. Codify infrastructure
Containerize

### 4. Codify dev environment
Remove requirements for a local installation of package managers like
pip, npm, etc and use a proper build system. [Bazel](https://www.bazel.build/), 
since it works with most languages.

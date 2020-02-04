# lncm/lnd-rpc

This repo aims to make _grpc_ communication with (any version of) `lnd` trivial.  

Currently Go, and Python are supported.

### Go

While `lnd` natively provides `.go` files for grpc communication, importing the entirety of `lnd` sometimes causes issues with dependencies (ex. the infamous `btcd`), etc.  We aim to solve it by having zero/minimal dependencies, and providing direct access to each version individually.

**tl;dr:** Just use ex. `github.com/lncm/lnd-rpc/v0.9.0/invicesrpc` in your source


### Python

This repo also holds the source (and scrips necessary to generate) the contents of [`lnd-rpc` PyPI package].  To use it install the version of the version you want to use, and 


[`lnd-rpc` PyPI package]: https://pypi.org/project/lnd-rpc/


## tl;dr

This repo helps with:

1. [download] - download all `lnd` `.proto`'s, and their dependencies
1. [generate-go] - generate `.go` sources for available `.proto` files
1. [generate-python] - generate `.py` sources for available `.proto` files

[download]: https://github.com/lncm/lnd-rpc/blob/58f303abd4b3e3d51c2d27872dc715bd5a0a5bed/scripts/download#L98-L227
[generate-go]: https://github.com/lncm/lnd-rpc/blob/58f303abd4b3e3d51c2d27872dc715bd5a0a5bed/scripts/generate-go#L89-L112
[generate-python]: https://github.com/lncm/lnd-rpc/blob/58f303abd4b3e3d51c2d27872dc715bd5a0a5bed/scripts/generate-python#L90-L131

Each of these can be done in two ways:

1. Run the script directly
1. Via `docker run` 


## Download

`./script/download` downloads _all_ `.proto` files, and all their dependencies unless a specific version is provided. 


### Run directly

```
./scripts/download --help
download v1.0.0

Download all .proto files necessary to build lnd's gRPC client libraries

Usage: ./scripts/download [options] LND_VERSION

Where LND_VERSION is in a form: [v]MAJOR.MINOR.PATCH (ex: v0.9.0), or "all" to download all versions

Options:

  -h, --help, help      Show this help message
  -G, --no-google       Skip download of google/api/* and google/protobuf/*
  -S, --strip-version   Don't include lnd version in the path (only works if LND_VERSION != "all")
  -o, --output          Download to a specified dir (will be created, if doesn't exist)

Examples:

  ./scripts/download  all                          # Download all lnd versions, and all google/* protos
  ./scripts/download  --no-google  v0.4.2          # Only download protos for lnd v0.4.2, and no google/* protos
  ./scripts/download  -G -S -o=~/last-lnd/ v0.9.0  # Only download protos for lnd v0.9.0, and save them to last-lnd/
                                                   #    in user's HOME directory w/o the /LND_VERSION/ segment in path

github: github.com/lncm/lnd-rpc/
```

### Docker

NOTE: this one requires `DOCKER_BUILDKIT=1` due to usage of `--target=`

```shell script
# Build with:
DOCKER_BUILDKIT=1  docker build . \
    --target=protos-downloader
    --tag=lnd-rpc-downloader

# Run with:
docker run --rm -it \
    --volume=$(pwd)/:/protos/ \
    lnd-rpc-downloader # [VERSION|all] 
```


## Generate Go

`./scripts/generate-go` generates `.go` files for _all_ available versions, unless a specific version is provided. 


### Run directly

```
./scripts/generate-go --help
generate-go v1.0.0

Compile all .proto definitions into importable .go files

Usage: generate-go [options] LND_VERSION

Where LND_VERSION is in a form: [v]MAJOR.MINOR.PATCH (ex: v0.9.0), or "all" to generate for all versions

Options:

  -h, --help, help      Show this help message
  -S, --strip-version   Don't include lnd version in the path (only works if LND_VERSION != "all")
  -o, --output          Save generated files to a specified dir (created, if doesn't exist)

Examples:

  ./generate-go all
  ./generate-go  -o /tmp/last/  v0.9.0

github: github.com/lncm/lnd-rpc/
```

### Docker

```shell script
docker build . \
    --build-arg="LANG=go" \
    --tag=lnd-rpc-go

docker run --rm -it \
    --volume=$(pwd):/data/go/ \
    lnd-rpc-go  # [VERSION|all]
```


## Generate Python

`./scripts/generate-python` generates `.py` files for _all_ available versions, unless a specific version is provided. 

> **NOTE:** All generated versions are published to PyPi using [this workflow]

[this workflow]: https://github.com/lncm/lnd-rpc/blob/58f303abd4b3e3d51c2d27872dc715bd5a0a5bed/.github/workflows/release-python.yml


### Run directly

```
./scripts/generate-python --help
generate-python v1.0.0

Compile all .proto definitions into .py files

Usage: generate-python [options] LND_VERSION

Where LND_VERSION is in a form: [v]MAJOR.MINOR.PATCH (ex: v0.9.0), or "all" to generate for all versions

Options:

  -h, --help, help      Show this help message
  -S, --strip-version   Don't include lnd version in the path (only works if LND_VERSION != "all")
  -o, --output          Save generated files to a specified dir (created, if doesn't exist)

Examples:

  ./generate-python all
  ./generate-python  -o /tmp/last/  v0.9.0

github: github.com/lncm/lnd-rpc/
```

### Docker 

```shell script
docker build . \
    --build-arg="LANG=python" \
    --tag=lnd-rpc-python

docker run --rm -it \
    --volume=$(pwd):/data/python/ \
    lnd-rpc-python  # [VERSION|all]
```

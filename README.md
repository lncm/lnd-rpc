# lncm/lnd-rpc

TODO: Describe the point of this repo better

## tl;dr

You can:

1. download
1. generate-go
1. generate-python


Each of these is possible in two ways:

1. Run the script directly
1. Via `docker run` 


## Download all `.proto`s

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

### Run in Docker

NOTE: this one requires `DOCKER_BUILDKIT=1` due to `--target=` being used

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

### Run in Docker

```shell script
docker build . \
    --build-arg="LANG=go" \
    --tag=lnd-rpc-go

docker run --rm -it \
    --volume=$(pwd):/data/go/ \
    lnd-rpc-go  # [VERSION|all]
```

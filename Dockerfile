#
##  This Dockerfile aims to simpolify generation of ready-to-use gRPC interfaces for multiple programming languages.
#
ARG LANG=unknown

ARG VER_ALPINE=3.11

ARG VER_PYTHON=3.8
ARG VER_GO=1.13



#
## Return an error when language not provided
#
FROM alpine:$VER_ALPINE  AS  unknown-builder

RUN printf '\nERR: $LANG has to be specified, try:\n\t%s\n\n' \
    'docker build --build-arg="LANG=python" …'

RUN exit 1



# TODO: Create & use non-root users for each container that's run w/volume



#
##  This stage creates image that can be used to download all protos:
#       DOCKER_BUILDKIT=1 docker build --target protos-downloader --tag protos-downloader .
#       docker run --rm -it -v=$(pwd)/:/protos protos-downloader v0.9.0
#
FROM alpine:$VER_ALPINE  AS  protos-downloader

RUN apk add --no-cache git jq

COPY  scripts/download  /usr/local/bin/

RUN mkdir /protos
WORKDIR /protos
VOLUME /protos

ENTRYPOINT ["download"]
CMD ["all"]



#
## Use the previously defined downloader to fetch all (lnd & deps) protos here
FROM protos-downloader  AS  protos-download
RUN download --output=/protos all



#
##  Create an image able to generate .go source files in a specified volume
#
FROM golang:$VER_GO-alpine$VER_ALPINE  AS  go-builder

# `git` for Go
# `findutils`, because `busybox`'s `find` is stupidly slow
RUN apk add --no-cache  git  findutils  protoc

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

COPY scripts/generate-go  /usr/local/bin/

ENV PROTOS /data/proto/
ENV GO_OUT /data/go/
RUN mkdir -p  "$PROTOS"  "$GO_OUT"

WORKDIR $PROTOS
VOLUME $GO_OUT

COPY --from=protos-download  /protos/  $PROTOS

ENTRYPOINT ["generate-go", "--output=/data/go/"]
CMD  ["all"]



#
## TODO: consider generation of other interfaces (client-only would be great…)
#
FROM go-builder  AS  go-builder-extra

RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway



#
##  Create an image able to generate .py source files in a specified volume
#
FROM python:$VER_PYTHON-alpine$VER_ALPINE  AS  python-builder

RUN apk add --no-cache  gcc g++ findutils protoc

RUN python -m pip install --upgrade pip
RUN python -m pip install grpcio grpcio-tools

COPY scripts/generate-python  /usr/local/bin/

ENV PROTOS /data/proto/
ENV PYTHON_OUT /data/python/
RUN mkdir -p  "$PROTOS"  "$PYTHON_OUT"

WORKDIR $PROTOS
VOLUME $PYTHON_OUT

COPY --from=protos-download  /protos/  $PROTOS

ENTRYPOINT ["generate-python", "--output=/data/python/"]
CMD  ["all"]



#
##  A convergence point for all LANGs
#
FROM $LANG-builder  AS  final

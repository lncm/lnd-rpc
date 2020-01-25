#!/bin/sh

ACTION=${1:-all}
TAG=$2

REPO=lightningnetwork/lnd

fetch_all_protos() {
  tag=$1
  dir=$2

  if [ "${dir}" != "" ]; then
    if [ "${dir}" != "${dir#https}" ]; then
      echo "${dir}"
      return
    fi

      dir="/${dir}"
  fi

  list=$(curl -sL "https://api.github.com/repos/${REPO}/contents/lnrpc${dir}?ref=${tag}" \
    | jq -r '.[] | if (.name | endswith("proto")) then .download_url elif (.type == "dir") then .name else empty end')

  for l in ${list}; do
    fetch_all_protos "${tag}" "${l}"
  done
}

download() {
  # Download all dependencies external to lnd
  DIR=google/api
  mkdir -p "${DIR}"
  curl -s https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > "${DIR}/annotations.proto"
  curl -s https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > "${DIR}/http.proto"

  TAGS=$(curl -s "https://api.github.com/repos/${REPO}/tags" | jq -r '.[].name' | sed 's|-beta||' | grep '^v\d.\d.\d$' | tac)

  for t in ${TAGS}; do
    if [ "${TAG}" != "" ] && [ "${TAG}" != "${t}" ]; then
      continue
    fi

    echo "Processing ${t}…"

    mkdir -p "${t}"

    for proto in $(fetch_all_protos "${t}-beta"); do
      printf "Downloading: %s…\n" "${proto}"
      wget -nH --cut-dirs=4  -xqcP "${t}/"  "${proto}"
    done

    echo
  done
}

generate() {
  for version in ./v**; do
    if [ "${TAG}" != "" ] && [ "./${TAG}" != "${version}" ]; then
      continue
    fi

    echo "${version}"
    for proto in $(find "${version}" -type f -name '*.proto'); do
      printf "Generating %s…\n" "${proto}"

      protoc --go_out=plugins=grpc,paths=source_relative:.  -I.  -I"${version}"  "${proto}"
    done

    echo
  done
}


if [ "${ACTION}" = "all" ] || [ "${ACTION}" = "download" ]; then
  download
fi

if [ "${ACTION}" = "all" ] || [ "${ACTION}" = "generate" ]; then
  generate
fi

#!/bin/sh

TAG=$1

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


ls

  #	go_file=rpc.pb.go

  #	cat "${t}/${proto_file}" | sed 's|github.com/lightningnetwork/lnd/lnrpc|lnd|'  > "lnd/${proto_file}"
  #
  #	protoc --go_out=plugins=grpc,paths=source_relative:./lnd/  -Ilnd/  "lnd/${proto_file}"
  #
  #	mv "lnd/${go_file}"  "${t}/"
  #
  #	if [[ -n "${prev}" ]]; then
  #		diff  "${prev}/${proto_file}"	"${t}/${proto_file}"	> "${t}/proto-diff-to-${prev}.patch"
  #		diff  "${prev}/${go_file}"		"${t}/${go_file}"			> "${t}/go-diff-to-${prev}.patch"
  #	fi
  #


}

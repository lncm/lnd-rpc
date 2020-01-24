#!/bin/sh

REPO=lightningnetwork/lnd

TAGS=$(curl -s "https://api.github.com/repos/${REPO}/tags" | jq -r '.[].name' | sed 's|-beta||' | grep '^v\d.\d.\d$' | tac)

DIR=lnd/google/api
mkdir -p "${DIR}"
curl -s https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > "${DIR}/annotations.proto"
curl -s https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > "${DIR}/http.proto"


prev=""
for t in ${TAGS}; do
	echo "Processing ${t}…"

	mkdir -p "${t}"

	proto_file=rpc.proto
	go_file=rpc.pb.go
	curl -s "https://raw.githubusercontent.com/${REPO}/${t}-beta/lnrpc/${proto_file}" > "${t}/${proto_file}"

	cat "${t}/${proto_file}" | sed 's|github.com/lightningnetwork/lnd/lnrpc|lnd|'  > "lnd/${proto_file}"

	protoc --go_out=plugins=grpc,paths=source_relative:./lnd/  -Ilnd/  "lnd/${proto_file}"

	mv "lnd/${go_file}"  "${t}/"

	if [[ -n "${prev}" ]]; then
		diff  "${prev}/${proto_file}"	"${t}/${proto_file}"	> "${t}/proto-diff-to-${prev}.patch"
		diff  "${prev}/${go_file}"		"${t}/${go_file}"			> "${t}/go-diff-to-${prev}.patch"
	fi

	prev=${t}
done

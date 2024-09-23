#!/usr/bin/env bash

# --------------
# Commands to run locally
# docker run --network host --rm -v $(CURDIR):/workspace --workdir /workspace ghcr.io/cosmos/proto-builder:v0.11.6 sh ./protocgen.sh
#
set -eo pipefail

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  proto_files=$(find "${dir}" -maxdepth 1 -name '*.proto')
  for file in $proto_files; do
    # this regex checks if a proto file has its go_package set to github.com/KiiChain/kiichainV3...
    # gogo proto files SHOULD ONLY be generated if this is false
    # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
    if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*kiichainV3' "$file"; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done
#

echo "Copying files"
cp -r github.com/KiiChain/kiichainV3/x/epoch/types/epoch.pb.go ./
#cp -r kiichainV3/* ./
#rm -rf github.com
#rm -rf kiichainV3
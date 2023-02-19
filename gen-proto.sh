#!/usr/bin/env bash

proto_dirs=$(find ./chain/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
proto_files=()

for dir in $proto_dirs; do
  proto_files=("${proto_files[@]} $(find "${dir}" -maxdepth 1 -name '*.proto')")
done

echo $proto_files

npx pbjs \
  -o ./src/proto.cjs \
  -t static-module \
  --force-long \
  --keep-case \
  --no-create \
  --path=./chain/proto/ \
  --path=./chain/third_party/proto/ \
  --root=crescent-client \
  ${proto_files[@]}

npx pbjs \
  -o ./src/proto.js \
  -t static-module \
  -w es6 \
  --es6 \
  --force-long \
  --keep-case \
  --no-create \
  --path=./chain/proto/ \
  --path=./chain/third_party/proto/ \
  --root=crescent-client \
  ${proto_files[@]}

npx pbts \
  -o ./src/proto.d.ts \
  ./src/proto.js

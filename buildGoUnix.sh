#!/bin/bash

ARCHS=(
  "linux.386.so"
  "linux.amd64.so"
)

for i in "${ARCHS[@]}"; do
  IFS="." read -r -a ARCH <<< "$i"

  echo "building: " "$i"

  env CGO_ENABLED=1 GOOS="${ARCH[0]}" GOARCH="${ARCH[1]}" go build -buildmode=c-shared -o PyChestBuild/"GoChest.""$i" CWrapper.go
done

git add PyChestBuild/*.so


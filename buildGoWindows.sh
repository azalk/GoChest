#!/bin/bash

ARCHS=(
  "windows.386.dll"
  "windows.amd64.dll"
)

for i in "${ARCHS[@]}"; do
  readarray -d . -t ARCH <<<"$i"

  CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 GOOS="${ARCH[0]}" GOARCH="${ARCH[1]}" go build -buildmode=c-shared -o GoBuild/"GoChest.""$i" CWrapper.go
done

git add PyChestBuild/*.so
git add PyChestBuild/*.dll

git commit -m "precompiled libraries"
git push
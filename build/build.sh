#!/bin/bash

set -e

rm -f dist/event-srv*

if [ -z "$VERSION" ]; then
  VERSION="0.0.1-dev"
fi
echo "Building application version $VERSION"

echo "Building default binaries"
CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X main.version=${VERSION}" -o "dist/event-srv" github.com/rafaeljesus/event-srv/cmd/event-srv

OS_PLATFORM_ARG=(linux darwin)
OS_ARCH_ARG=(amd64)
for OS in ${OS_PLATFORM_ARG[@]}; do
  for ARCH in ${OS_ARCH_ARG[@]}; do
    echo "Building binaries for $OS/$ARCH..."
    GOARCH=$ARCH GOOS=$OS CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X main.version=${VERSION}" -o "dist/event-srv_$OS-$ARCH" github.com/rafaeljesus/event-srv/cmd/event-srv
  done
done

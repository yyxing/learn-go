#!/bin/bash
SOURCE_FILE_NAME=main
TARGET_FILE_NAME=reskd
build() {
  echo $GOOS $GOARCH
  if [ "$GOOS" == "windows" ]; then
    EXT=".exe"
  else
    EXT=""
  fi
  tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}
  env GOOS="$GOOS" GOARCH="$GOARCH" go build -o ${tname} \
  -v ${SOURCE_FILE_NAME}.go
  chmod +x ${tname}
}
CGO_ENABLED=0
#mac os 64
GOOS=darwin
GOARCH=amd64
build

#linux 64
GOOS=linux
GOARCH=amd64
build

##windows 64
#GOOS=windows
#GOARCH=amd64
#build



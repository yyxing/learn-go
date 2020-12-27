#!/bin/bash
SOURCE_FILE_NAME=main
TARGET_FILE_NAME="$3"
build(){
   env GOOS="$4" GOARCH="$5"
   echo GOOS GOARCH
   if [ -z "$GOOS" ]; then
      echo "The build command must be in the third, the four parameters specify GOOS and GOARCH"
   fi
}
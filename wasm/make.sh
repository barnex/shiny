#! /bin/bash

rm -f lib.wasm
GOARCH=wasm GOOS=js go build -o lib.wasm

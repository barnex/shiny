#! /bin/bash

rm -f play.wasm edit.wasm
./build_levels.sh
GOARCH=wasm GOOS=js go build -o play.wasm
GOARCH=wasm GOOS=js go build -o edit.wasm github.com/barnex/shiny/edit 

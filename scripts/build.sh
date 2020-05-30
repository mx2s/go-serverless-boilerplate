#!/bin/bash
GOPATH=$PWD go build src/app/main.go
mkdir -p build
mv main build/go-serverless

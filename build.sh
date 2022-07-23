#!/bin/bash

rm -rf out/ vendor/
go fmt ./...
go mod tidy
go mod vendor
go build -o out/wol -mod vendor


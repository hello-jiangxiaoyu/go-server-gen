#!/bin/bash

go build -ldflags="-s -w" -o "${GOPATH}"/bin/gsg .

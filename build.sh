#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o bin/cyj.exe main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o bin/cyj main.go

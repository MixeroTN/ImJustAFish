#!/bin/bash

go build -ldflags "-H=windowsgui" -o fish.exe src/main.go
./fish.exe

#!/bin/bash
echo "Start build macOS 64 bit"
GOOS=darwin GOARCH=amd64 go build -o build/macos/64/intergalactic
echo  "Completed"
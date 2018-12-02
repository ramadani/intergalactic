#!/bin/bash
echo "Start build macOS 64 bit"
GOOS=darwin GOARCH=amd64 go build -o build/macos/64/intergalactic
echo  "Completed"
echo "Start build linux 32 bit"
GOOS=linux GOARCH=386 go build -o build/linux/32/intergalactic
echo "Completed"
echo "Start build linux 64 bit"
GOOS=linux GOARCH=amd64 go build -o build/linux/64/intergalactic
echo "Completed"
echo "Start build windows 64 bit"
GOOS=windows GOARCH=amd64 go build -o build/windows/64/intergalactic
echo "Completed"
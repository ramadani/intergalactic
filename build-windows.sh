#!/bin/bash
echo "Start build windows 64 bit"
GOOS=windows GOARCH=amd64 go build -o build/windows/64/intergalactic
echo "Completed"
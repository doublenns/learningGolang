#!/usr/bin/env bash

echo "Building the binary.."
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating a ZIP file of the built binary.."
zip deployment.zip main

echo "Cleaning up.."
rm main
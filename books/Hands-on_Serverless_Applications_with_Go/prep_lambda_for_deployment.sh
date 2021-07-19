#!/usr/bin/env bash

set -e

echo "Building the binary.."
GOOS=linux GOARCH=amd64 go build -o main ${PWD}/main.go

echo "Creating a ZIP file of the built binary.."
zip deployment.zip main

echo "Cleaning up.."
rm main
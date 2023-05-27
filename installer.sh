#!/bin/bash

# Install Go
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."
    apt-get update
    apt-get install -y golang
fi

# Install required Go packages
echo "Installing required Go packages..."
go install golang.org/x/crypto/ssh@latest
go install github.com/fatih/color@latest

# Compile the Go code
echo "Compiling the Go code..."
go build -o myprogram main.go

echo "Installation completed."
echo "Thank you for using Excalibur"

                                                                              
                                                                              
                                                                              

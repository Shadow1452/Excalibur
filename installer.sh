#!/bin/bash

# Install Go
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."
    apt-get update
    apt-get install -y golang
fi

# Install required Go packages
echo "Installing required Go packages..."
go get golang.org/x/crypto/ssh
go get github.com/fatih/color

# Compile the Go code
echo "Compiling the Go code..."
go build -o myprogram main.go

echo "Installation completed."
echo "Thank you for using Excalibur"
echo " /$$$$$$$$                               /$$ /$$ /$$                          
| $$_____/                              | $$|__/| $$                          
| $$       /$$   /$$  /$$$$$$$  /$$$$$$ | $$ /$$| $$$$$$$  /$$   /$$  /$$$$$$ 
| $$$$$   |  $$ /$$/ /$$_____/ |____  $$| $$| $$| $$__  $$| $$  | $$ /$$__  $$
| $$__/    \  $$$$/ | $$        /$$$$$$$| $$| $$| $$  \ $$| $$  | $$| $$  \__/
| $$        >$$  $$ | $$       /$$__  $$| $$| $$| $$  | $$| $$  | $$| $$      
| $$$$$$$$ /$$/\  $$|  $$$$$$$|  $$$$$$$| $$| $$| $$$$$$$/|  $$$$$$/| $$      
|________/|__/  \__/ \_______/ \_______/|__/|__/|_______/  \______/ |__/ "     
                                                                              
                                                                              
                                                                              
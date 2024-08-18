#!/bin/bash

set -e

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to find Go binary
find_go_binary() {
    if command_exists go; then
        echo "go"
    elif [ -x "/usr/local/go/bin/go" ]; then
        echo "/usr/local/go/bin/go"
    elif [ -x "/opt/homebrew/bin/go" ]; then
        echo "/opt/homebrew/bin/go"
    elif [ -x "$HOME/go/bin/go" ]; then
        echo "$HOME/go/bin/go"
    else
        echo ""
    fi
}

# Function to install Rclone
install_rclone() {
    echo "Installing Rclone..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        brew install rclone
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        curl https://rclone.org/install.sh | sudo bash
    elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        powershell -Command "iwr https://rclone.org/install.ps1 -UseBasicParsing | iex"
    else
        echo "Unsupported OS for Rclone installation"
        exit 1
    fi
}

# Function to install Go
install_go() {
    echo "Installing Go..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        brew install go
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        sudo apt-get update
        sudo apt-get install -y golang-go
    elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        powershell -Command "
            $url = 'https://golang.org/dl/go1.17.windows-amd64.msi'
            $outpath = '$env:TEMP\go1.17.windows-amd64.msi'
            Invoke-WebRequest -Uri $url -OutFile $outpath
            Start-Process msiexec.exe -Wait -ArgumentList '/I $outpath /quiet'
            Remove-Item $outpath
        "
    else
        echo "Unsupported OS for Go installation"
        exit 1
    fi
}

# Function to install Go dependencies
install_go_dependencies() {
    echo "Installing Go dependencies..."
    GO_BIN=$(find_go_binary)
    if [ -z "$GO_BIN" ]; then
        echo "Go binary not found. Please ensure Go is installed correctly."
        exit 1
    fi
    $GO_BIN get github.com/kirsle/configdir
    $GO_BIN get golang.org/x/crypto
}

# Main installation process
main() {
    if ! command_exists rclone; then
        install_rclone
    else
        echo "Rclone is already installed."
    fi

    GO_BIN=$(find_go_binary)
    if [ -z "$GO_BIN" ]; then
        install_go
        GO_BIN=$(find_go_binary)
    else
        echo "Go is already installed at $GO_BIN"
    fi

    # Ensure Go binary directory is in PATH
    GO_DIR=$(dirname "$GO_BIN")
    export PATH=$PATH:$GO_DIR

    install_go_dependencies

    echo "Installation process completed."
}

# Run main function
main
Go Playground
=============

Just testing Go Language. Nothing useful here

## Setup

For Ubuntu 14.04, just install golang metapackage

    sudo apt-get install golang

## Ctrl + C stuff

Just run and exit
    go run hello.go

Make executable
    # Create binary hello (*nix) / hello.exe (Windows)
    go build hello.go

    # Cross compile, from Linux x64 to Windows Win32
    GOARCH=386 GOOS=windows go build -o hello.exe hello.go
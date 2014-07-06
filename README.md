Go Playground
=============

Just testing Go Language. Nothing useful here

## Setup

For Ubuntu 14.04, just install golang metapackage

    sudo apt-get install golang

Any text editor will work. 

## (Simple) Crash Course about how to compile


### hello example

Just run the code (compile, execute and discart binary)

    go run hello.go

Make executable
    # Create binary hello (*nix) / hello.exe (Windows)
    go build hello.go

    # Cross compile, from Linux x64 to Windows Win32
    GOARCH=386 GOOS=windows go build -o hello.exe hello.go


    
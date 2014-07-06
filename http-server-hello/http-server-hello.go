// Run server with command:
//   $ go run http-server-hello
// Visit http://localhost:8080/ [Hi there, I love !]
// Visit http://localhost:8080/test [http://localhost:8080/test]
//
// To build a win32 binary, run 
//   $ GOARCH=386 GOOS=windows go build -o http-server-hello.exe http-server-hello.go

package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
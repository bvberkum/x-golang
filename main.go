package main // import "github.com/bvberkum/x-golang"

import (
	"os"
	"fmt"
	"runtime"
)

var app_id = "x-golang"
var version = "0.0.2-dev" // x-golang

func main() {
  fmt.Println("Hello world!")
  fmt.Printf("This is a %s/main.go, version %s, running as process %d\n", app_id, version, os.Getpid())

  fmt.Printf("Compiled with %s for %s/%s\n", runtime.Compiler, runtime.GOOS, runtime.GOARCH)
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("number of processor: ", runtime.NumCPU())
	fmt.Println("GOROOT: ", runtime.GOROOT())
	fmt.Println("ARCH: ", runtime.GOARCH)
	fmt.Println("OS: ", runtime.GOOS)
}

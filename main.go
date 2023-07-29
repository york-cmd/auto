package main

import (
	"auto/initpkg"
	"os"
)

func main() {
	initpkg.Run()
	os.RemoveAll("tmp/")
}

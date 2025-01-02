package main

import (
	"os"

	"github.com/happymanju/convertwebp/convertwebp"
)

func main() {
	os.Exit(convertwebp.Run(os.Args[1:]))
}

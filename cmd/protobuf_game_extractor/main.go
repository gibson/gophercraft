package main

import (
	"os"

	"github.com/gibson/gophercraft/cmd/protobuf_game_extractor/ex"
	"github.com/gibson/gophercraft/utils/log"
)

func main() {
	program := os.Args[1]
	folder := os.Args[2]

	goPackagePrefix := "github.com/Gophercraft/core/bnet/bnetproto"

	if err := ex.Extract(goPackagePrefix, program, folder); err != nil {
		log.Fatal(err)
	}
}

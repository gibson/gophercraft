package autocode

import (
	"log"

	"github.com/gibson/gophercraft/vsn"
)

func runMain(build vsn.Build, loc string) {
	g := NewGenerator(build, loc)
	if err := g.Generate(); err != nil {
		log.Fatal(err)
	}
}

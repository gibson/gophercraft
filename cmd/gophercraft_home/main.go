package main

import (
	"github.com/gibson/gophercraft/home"
	_ "github.com/gibson/gophercraft/home/dbsupport"
)

func main() {
	home.Run()
}

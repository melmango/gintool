package main

import (
	"os"
)

func main() {
	args := os.Args
	if (len(args) < 1) {
		println("no detail command, try 'gintool help' for command list")
	} else {
		switch args[1] {
		case "help":helpDoc()
		case "new":NewApp(args)
		default: println("Unrecognized command")
		}
	}
}

func helpDoc() {
	println(docTmpl)
}

var docTmpl = `Gintool is a tool for gin

Usage: gintool command [arguments]

The command are:
	new    Create a new gin application

Other functions will be implemented later.
`
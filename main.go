package main

import (
	"os"

	"github.com/arisu-archive/arona-unflatd/cmd/root"
)

var version = "dev"

func main() {
	root.Execute(root.ExecuteOptions{
		Version: version,
		Exit:    os.Exit,
		In:      os.Stdin,
		Out:     os.Stdout,
		Err:     os.Stderr,
	}, os.Args[1:])
}

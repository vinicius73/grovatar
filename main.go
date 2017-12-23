package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Email string `short:"e" long:"email" description:"A email" required:"true"`
}

var opts = options{}

func init() {
	flags.ParseArgs(&opts, os.Args)
}

func main() {
	avatar, _ := MakeAvatar(opts.Email)
	OutputStdout(avatar)
}

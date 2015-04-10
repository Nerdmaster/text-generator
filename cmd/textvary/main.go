package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"go.nerdbucket.com/text/pkg/filter/variation"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var opts struct {
	Seed int64  `short:"s" long:"seed" description:"Seed for PRNG"`
	Text string `short:"t" long:"text" description:"Text to vary"`
}

var parser = flags.NewParser(&opts, flags.PassDoubleDash|flags.HelpFlag)

func usage() {
	parser.WriteHelp(os.Stderr)
	os.Exit(1)
}

func main() {
	parser.Usage = `--text "Something with {{some|a few|several}} variation patterns"

Runs a quick variation filter over the given text.  The filter randomly chooses
a single item within pipe-delimited values inside double-curly-braces.
`
	_, err := parser.Parse()

	if err != nil {
		usage()
	}

	fromStdin := false
	if opts.Text == "" {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			fromStdin = true
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Printf("Error reading from stdin: %s\n\n", err)
				usage()
			}

			opts.Text = string(bytes)
		}

		if opts.Text == "" {
			usage()
		}
	}

	// If no seed was passed in, generate one
	if opts.Seed == 0 {
		opts.Seed = time.Now().UTC().UnixNano()
	}

	rand.Seed(opts.Seed)

	v := variation.New()

	if fromStdin {
		fmt.Print(v.Filter(opts.Text))
	} else {
		fmt.Println(v.Filter(opts.Text))
	}
}

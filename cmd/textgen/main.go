package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/Nerdmaster/text-generator/pkg/filter/iafix"
	"github.com/Nerdmaster/text-generator/pkg/filter/substitution"
	"github.com/Nerdmaster/text-generator/pkg/filter/variation"
	"github.com/Nerdmaster/text-generator/pkg/generator"
	"github.com/Nerdmaster/text-generator/pkg/template"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var opts struct {
	Seed               int64             `short:"s" long:"seed" description:"Seed for PRNG"`
	StringListOverride map[string]string `long:"value" description:"Override a word list with a specific value"`
}

var parser = flags.NewParser(&opts, flags.PassDoubleDash|flags.HelpFlag)

func usage() {
	parser.WriteHelp(os.Stderr)
	os.Exit(1)
}

func main() {
	parser.Usage = `[template file] [wordlist] [OPTIONS]

Reads [template file] and all *.txt files in the given [wordlist] directory,
and produces random text, recursively replacing anything in double-curly-braces
with a random item from a wordlist file of the same name.

e.g., if your template includes {{noun}} somewhere in it, a file called
[wordlist]/noun.txt is expected to exist, and one of the lines will be put in
place of the template's "{{noun}}" text.

If desired, instead of a wordlist, a value can be passed on the command line,
but this only allows a single value for a given wordlist, so wouldn't work as
well for parts of speech as for single-use terms like {{nameofboy}}.

e.g.: --value "nameofboy:Nerd Master"`
	args, err := parser.Parse()
	if err != nil || len(args) != 2 {
		usage()
	}

	t, err := template.FromFile(args[0])
	if err != nil {
		fmt.Printf("Error trying to read the template file '%s': %s", args[0], err)
		os.Exit(1)
	}

	subFilter := substitution.New()
	subFilter.NullGeneratorFactory = func(id string) generator.Generator {
		fmt.Printf("ERROR: Requested generator, '%s', does not exist", id)
		return substitution.MakeNullGenerator(id)
	}

	buildWordlists(subFilter, args[1])
	t.AddFilter(subFilter)
	t.AddFilter(variation.New())
	t.AddFilter(iafix.New())

	// Load overrides
	for name, value := range opts.StringListOverride {
		subFilter.SetValue(name, value)
	}

	// If no seed was passed in, generate one
	if opts.Seed == 0 {
		opts.Seed = time.Now().UTC().UnixNano()
	}

	rand.Seed(opts.Seed)

	fmt.Println(t.Execute())
}

func buildWordlists(subFilter *substitution.Substitution, dirname string) {
	// Read in all *.txt files to populate string lists
	dataFiles, err := filepath.Glob(fmt.Sprintf("%s/*.txt", dirname))
	if err != nil {
		fmt.Println("Error trying to read word lists:", err)
		os.Exit(1)
	}

	// Pull all 'wordlist' files and populate the StringList array
	for _, file := range dataFiles {
		fileBytes, _ := ioutil.ReadFile(file)
		fileData := string(fileBytes)
		listname := strings.Replace(path.Base(file), ".txt", "", -1)
		list := generator.MakeRandom()
		subFilter.SetGenerator(listname, list)

		for _, str := range strings.Split(fileData, "\n") {
			if strings.TrimSpace(str) != "" {
				list.Append(str)
			}
		}

		if list.IsEmpty() {
			fmt.Printf("FATAL: List '%s' exists but has no data!\n", listname)
			os.Exit(1)
		}
	}
}

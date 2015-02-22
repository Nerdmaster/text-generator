package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var opts struct {
	Seed int64 `short:"s" long:"seed" description:"Seed for PRNG"`
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
place of the template's "{{noun}}" text.`
	args, err := parser.Parse()
	if err != nil || len(args) != 2 {
		usage()
	}

	template := readTemplate(args[0])
	lists := readWordlists(args[1])

	// If no seed was passed in, generate one
	if opts.Seed == 0 {
		opts.Seed = time.Now().UTC().UnixNano()
	}

	rand.Seed(opts.Seed)

	// Read the template and populate data
	tvarRegex := regexp.MustCompile(`{{([^}]*)}}`)
	for {
		foundStrings := tvarRegex.FindStringSubmatch(template)
		if foundStrings == nil {
			break
		}

		// Set up a variable to hold the replacement value
		replacementValue := ""

		// Store the full match in an alias for easier replacing later
		fullMatch := foundStrings[0]

		// Handle possible variable assignments
		data := strings.Split(foundStrings[1], "->")
		listname := data[0]
		variable := ""
		if len(data) == 2 {
			variable = data[1]
		}

		// See if the list exists and warn if not
		list := lists[listname]
		if list == nil {
			fmt.Printf("ERROR: List '%s' needed but doesn't exist\n", listname)
		} else {
			replacementValue = list.RandomString()
		}

		if variable != "" {
			lists[variable] = NewStringList()
			lists[variable].AddString(replacementValue)
		}

		template = strings.Replace(template, fullMatch, replacementValue, 1)
	}

	fmt.Println(template)
}

func readTemplate(filename string) string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error trying to read the template file '%s': %s", filename, err)
		os.Exit(1)
	}

	return string(fileBytes)
}

func readWordlists(dirname string) map[string]*StringList {
	// Maps a word type ("noun", etc) to a string list containing possible values
	// for the given word type
	lists := make(map[string]*StringList)

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
		lists[listname] = NewStringList()

		for _, str := range strings.Split(fileData, "\n") {
			if strings.TrimSpace(str) != "" {
				lists[listname].AddString(str)
			}
		}
	}

	// Throw out errors if any lists are empty
	for listname, list := range lists {
		if list.masterList.Len() == 0 {
			fmt.Printf("FATAL: List '%s' exists but has no data!\n", listname)
			os.Exit(1)
		}
	}

	return lists
}

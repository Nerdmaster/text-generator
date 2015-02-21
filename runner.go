package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func usage() {
	fmt.Printf("usage: %s <OPTIONS...> [template file] [wordlist]\n", os.Args[0])
	fmt.Println("")
	fmt.Println("Reads [template file] and all *.txt files in the given [wordlist] directory,")
	fmt.Println("and produces random text, recursively replacing anything in double-curly-")
	fmt.Println("braces with a random item from a file of the same name.")
	fmt.Println("")
	fmt.Println("Unfortunately, options must be provided before the template file and word")
	fmt.Println("list.  This is a limitation of the effective, but simple, Go flag library.")
	fmt.Println("")
	fmt.Println("Options:")
	flag.CommandLine.VisitAll(func (f *flag.Flag) {
		prefix := fmt.Sprintf("  --%s", f.Name)
		fmt.Printf("%-16s %s (defaults to %s)\n", prefix, f.Usage, f.DefValue)
	})
	os.Exit(1)
}

func main() {
	var templateFilename, wordlistDirectoryName string
	var seed int64

	templateFilename, wordlistDirectoryName, seed = parseCLI()
	template := readTemplate(templateFilename)
	lists := readWordlists(wordlistDirectoryName)

	// If no seed was passed in, generate one
	if seed == 0 {
		seed = time.Now().UTC().UnixNano()
	}

	rand.Seed(seed)

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

func parseCLI() (string, string, int64) {
	var s int64

	flag.Usage = usage
	flag.Int64Var(&s, "seed", 0, "Seed for PRNG")
	flag.Parse()

	if len(flag.Args()) < 2 {
		usage()
	}

	t := flag.Arg(0)
	w := flag.Arg(1)

	if t == "" || w == "" {
		usage()
	}

	return t, w, s
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

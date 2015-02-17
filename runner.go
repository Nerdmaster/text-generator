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
	fmt.Printf("usage: %s --template [template file] --wordlist [word list directory]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	var templateFilename, wordlistDirectoryName string
	parseCLI(&templateFilename, &wordlistDirectoryName)
	template := readTemplate(templateFilename)
	lists := readWordlists(wordlistDirectoryName)

	// Seed the PRNG so stuff is unique
	rand.Seed(time.Now().UTC().UnixNano())

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

func parseCLI(t *string, w *string) {
	flag.StringVar(t, "template", "", "Template file for building random text")
	flag.StringVar(w, "wordlist", "", "Directory where word lists are located")
	flag.Parse()

	if *t == "" || *w == "" {
		usage()
	}
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

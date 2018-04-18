// +build ignore
// This is meant to be run via `go run makedeps.go`, not any compilation process

package main

import (
	"fmt"
	"os"
	"strings"
)

type File struct {
	*os.File
}

var namespace = "go.nerdbucket.com"
var project = "text"
var imports = []string{
	"github.com/jessevdk/go-flags",
}
var magic = "# GENERATED FOR GO DEPENDENCIES - DO NOT EDIT!"

func main() {
	varsFile, err := os.Create("go.vars.mk")
	if err != nil {
		fmt.Println("Error trying to write go.vars.mk: ", err)
		os.Exit(1)
	}
	defer varsFile.Close()
	vars := &File{varsFile}

	rulesFile, err := os.Create("go.rules.mk")
	if err != nil {
		fmt.Println("Error trying to write go.rules.mk: ", err)
		os.Exit(1)
	}
	defer rulesFile.Close()
	rules := &File{rulesFile}

	vars.WriteLn(magic)
	vars.WriteLn("")

	vars.WriteLn("NAMESPACE_DIR=%s", pathify(namespace))
	vars.WriteLn("SYMLINK_DIR=$(NAMESPACE_DIR)/%s", project)
	vars.WriteLn("SYMLINK_EXISTS=$(SYMLINK_DIR)/Makefile")
	vars.WriteLn("DEPS=%s", strings.Join(pathifyList(imports), " "))
	vars.WriteLn("")
	vars.WriteLn(`# This is to be used by binaries to verify all dependencies are "built"`)
	vars.WriteLn("ALLDEPS=$(SYMLINK_EXISTS) $(DEPS)")

	rules.WriteLn(magic)
	rules.WriteLn("")

	rules.WriteLn(".PHONY: alldeps")
	rules.WriteLn("alldeps: $(ALLDEPS)")
	rules.WriteLn("")
	rules.WriteLn("# We only depend on a file within the project; if we just depend on our")
	rules.WriteLn("# project's symlink, *any* change to the directory triggers a build")
	rules.WriteLn("$(SYMLINK_EXISTS):")
	rules.WriteLn("\tmkdir -p $(NAMESPACE_DIR)")
	rules.WriteLn("\tln -s $(CURDIR) $(SYMLINK_DIR)")

	for _, imp := range imports {
		rules.WriteLn("")
		rules.WriteLn("%s:", pathify(imp))
		rules.WriteLn("\tgo get %s", imp)
	}
}

func (f File) WriteLn(format string, args ...interface{}) {
	f.WriteString(fmt.Sprintf(format, args...) + "\n")
}

func pathify(s string) string {
	return "$(GOPATH)/src/" + s
}

func pathifyList(list []string) []string {
	newlist := make([]string, len(list))
	for idx, s := range list {
		newlist[idx] = pathify(s)
	}
	return newlist
}

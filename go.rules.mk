# GENERATED FOR GO DEPENDENCIES - DO NOT EDIT!

.PHONY: alldeps
alldeps: $(ALLDEPS)

# We only depend on a file within the project; if we just depend on our
# project's symlink, *any* change to the directory triggers a build
$(SYMLINK_EXISTS):
	mkdir -p $(NAMESPACE_DIR)
	ln -s $(CURDIR) $(SYMLINK_DIR)

$(GOPATH)/src/github.com/jessevdk/go-flags:
	go get github.com/jessevdk/go-flags

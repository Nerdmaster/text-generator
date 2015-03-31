SRCS=cmd/textgen/main.go \
	pkg/stringlist/*.go \
	pkg/template/*.go \
	pkg/generator/*.go \
	pkg/filter/*.go \
	pkg/filter/substitution/*.go

.PHONY: all clean format lint test

all: bin/textgen

bin/textgen: $(SRCS)
	mkdir -p bin
	go build -o ./bin/textgen ./cmd/textgen

clean:
	rm -f ./bin/*

format:
	find . -name "*.go" | xargs gofmt -l -w -s

lint:
	golint ./pkg/...

test:
	go test ./pkg/...

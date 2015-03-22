SRCS=cmd/textgen/main.go \
	lib/stringlist/*.go \
	lib/template/*.go \
	lib/generator/*.go \
	lib/filter/*.go \
	lib/filter/substitution/*.go

.PHONY : all format clean lint

all: bin/textgen

bin/textgen: $(SRCS)
	mkdir -p bin
	go build -o ./bin/textgen ./cmd/textgen

clean:
	rm -f ./bin/*

format:
	find . -name "*.go" | xargs gofmt -l -w -s

lint:
	golint ./lib/...

test:
	go test ./lib/...

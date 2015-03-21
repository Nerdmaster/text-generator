SRCS=cmd/textgen/main.go lib/stringlist/*.go lib/template/*.go

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

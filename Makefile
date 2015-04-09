all: bin/textgen

go.vars.mk: makedeps.go
	go run ./makedeps.go
go.rules.mk: makedeps.go
	go run ./makedeps.go

include go.vars.mk

SRCS=cmd/textgen/main.go \
	pkg/stringlist/*.go \
	pkg/template/*.go \
	pkg/generator/*.go \
	pkg/filter/*.go \
	pkg/filter/substitution/*.go \
	pkg/filter/iafix/*.go

.PHONY: all clean format lint test

bin/textgen: $(ALLDEPS) $(SRCS)
	go build -o ./bin/textgen ./cmd/textgen

clean:
	rm -r ./bin

format:
	find . -name "*.go" | xargs gofmt -l -w -s

lint:
	golint ./pkg/...

test:
	go test ./pkg/...

include go.rules.mk

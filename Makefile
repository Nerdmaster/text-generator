SRCS=cmd/textgen/main.go lib/stringlist/list.go lib/stringlist/randomizer.go

.PHONY : all

all: bin/textgen

bin/textgen: $(SRCS)
	mkdir -p bin
	go build -o ./bin/textgen ./cmd/textgen

clean:
	rm ./bin/*

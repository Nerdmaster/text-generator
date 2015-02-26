SRCS=runner.go stringlist.go

.PHONY : all

all: bin/textgen

bin/textgen: $(SRCS)
	mkdir -p bin
	go build -o ./bin/textgen .

clean:
	rm ./bin/*

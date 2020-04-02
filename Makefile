BIN=./bin
BINCGO=$(BIN)/certs-cgo
BINNOCGO=$(BIN)/certs-nocgo

all: build
build: $(BINCGO) $(BINNOCGO)

$(BIN):
	mkdir -p $@

$(BINCGO): $(BIN)
	CGO_ENABLED=1 go build -o $@ .

$(BINNOCGO): $(BIN)
	CGO_ENABLED=0 go build -o $@ .

clean:
	rm $(BIN)/*

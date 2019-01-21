GCB = ./cmd/gcb
GCBS = ./cmd/gcbs
SRC = $(shell find . -name '*.go')
GO = go
.PHONY: clean all run race

all: gcb gcbs

gcb: $(SRC)
	$(GO) build $(GCB)

gcbs: $(SRC)
	$(GO) build $(GCBS)

run:
	$(GO) run $(GCB)

race:
	$(GO) run -race $(GCB)

clean:
	rm -f ./gcb ./gcbs

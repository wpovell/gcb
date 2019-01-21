GCB = ./cmd/gcb
GCBS = ./cmd/gcbs
SRC = $(shell find . -name '*.go')
GO = go

.PHONY: clean all run race fmt

all: fmt gcb gcbs

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

fmt:
	gofmt -l -w .

CMD = ./cmd/gcb
SRC = $(shell find . -name '*.go')
GO = go
.PHONY: clean all run race

all: gcb

gcb: $(SRC)
	$(GO) build $(CMD)

run:
	$(GO) run $(CMD)

race:
	$(GO) run -race $(CMD)

clean:
	rm ./gcb

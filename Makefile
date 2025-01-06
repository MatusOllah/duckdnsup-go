GO = go

GOOS = $(shell $(GO) env GOOS)
GOARCH = $(shell $(GO) env GOARCH)

GO_FLAGS = -gcflags="-dwarf=false" -ldflags="-s -w"

BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE = $(BINARY)/duckdnsup-go$(shell $(GO) env GOEXE)

.PHONY: all
all: clean $(EXE)

$(EXE):
	mkdir -p $(BINARY)

	$(GO) get

	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE)

.PHONY: clean
clean:
	rm -rf $(BINARY)

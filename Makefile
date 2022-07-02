

BINARY := mm-webpage

ALL: $(BINARY)

$(BINARY): *.go */*.go go.mod
	go build -o $(BINARY)

clean:
	rm -f $(BINARY)

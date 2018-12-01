GOCMD=go

GOTEST=$(GOCMD) test

t:
	$(GOTEST) -v ./test

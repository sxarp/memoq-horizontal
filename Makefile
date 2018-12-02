GOCMD=go

GOTEST=$(GOCMD) test

test:
	$(GOTEST) -v ./...

# usage: $ make unit -d=app/hoge
unit:
	$(GOTEST) -v ./$(d)

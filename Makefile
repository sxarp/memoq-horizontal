GOCMD=go

GOTEST=$(GOCMD) test

test:
	$(GOTEST) -v -cover ./...

# usage: $ make unit -d=app/hoge
unit:
	$(GOTEST) -v -cover ./$(d)

build:
	$(GOCMD) build -o dist/server ./app/server

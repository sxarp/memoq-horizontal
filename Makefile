GOCMD=go

GOTEST=$(GOCMD) test

test:
	$(GOTEST) -v -cover -count=1 ./... # never cache

# usage: $ make unit -d=app/hoge
unit:
	$(GOTEST) -v -cover ./$(d)

build:
	$(GOCMD) build -o dist/server ./app/server

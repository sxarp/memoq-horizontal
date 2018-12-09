GOCMD=go
GOTEST=$(GOCMD) test

# setting for datastore emulator
DATASTORE_EMULATOR_HOST=localhost:8081
DATASTORE_PROJECT_ID=my-project-id

test:
	DATASTORE_EMULATOR_HOST=$(DATASTORE_EMULATOR_HOST) \
	DATASTORE_PROJECT_ID=$(DATASTORE_PROJECT_ID) \
	$(GOTEST) -v -cover -count=1 ./... # run tests without using cache

unit:
	DATASTORE_EMULATOR_HOST=$(DATASTORE_EMULATOR_HOST) \
	DATASTORE_PROJECT_ID=$(DATASTORE_PROJECT_ID) \
	$(GOTEST) -v -cover -count=1 -run $(f) ./$(d) # usage: $ make unit -f=LileFuncName -d=app/hoge

build:
	$(GOCMD) build -o dist/server ./app/server

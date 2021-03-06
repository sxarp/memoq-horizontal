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
	$(GOTEST) -v -cover -count=1 -run $(f) ./$(d) # $ make unit f=LileFuncName d=app/hoge

build:
	$(GOCMD) build -o dest/server ./app/server

local:
	DATASTORE_EMULATOR_HOST=$(DATASTORE_EMULATOR_HOST) \
	DATASTORE_PROJECT_ID=$(DATASTORE_PROJECT_ID) \
	HORIZONTAL_ENV=development $(GOCMD) run $$(ls -1 app/server | grep -v _test | awk '{print "app/server/"$$1}')

# setup

## setup symlink
`$ ln -s $PWD/app vendor/app`

# datastore emulator

## install
`$ gcloud components install cloud-datastore-emulator`

## start
`$ gcloud beta emulators datastore start`

## before running tests
```
$ export DATASTORE_EMULATOR_HOST=localhost:8432
$ export DATASTORE_PROJECT_ID=my-project-id
```

# run tests
`$ go test ./test`

# fmt all files 
`$ go fmt ./...`

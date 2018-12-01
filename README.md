# setup

## setup symlink
`$ ln -s $PWD/app vendor/app`

# datastore emulator

## install
`$ gcloud components install cloud-datastore-emulator`

## start
`$ gcloud beta emulators datastore start`

# run tests
`$ go test ./test`

# fmt all files 
`$ go fmt ./...`

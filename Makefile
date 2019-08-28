all: deps bin/derive bin/keys bin/seed bin/sign

test: deps
	go test -v ./...

deps:
	go get -v -d ./...

bin/derive: ./example/derive.go
	go build -o $@ $^
bin/keys: ./example/keys.go
	go build -o $@ $^
bin/seed: ./example/seed.go
	go build -o $@ $^
bin/sign: ./example/sign.go
	go build -o $@ $^


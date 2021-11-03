# TODO(dape): update to only generate if any proto file changed. (once tests take 'too long')

build: protoc-gen-bson/protoc-gen-bson

test: build
	go generate ./tests/...
	go test -v ./tests/...

protoc-gen-bson/protoc-gen-bson: protoc-gen-bson/main.go
	go build -o protoc-gen-bson/protoc-gen-bson ./protoc-gen-bson/main.go

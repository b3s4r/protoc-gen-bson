Every directly should hold a self contained test; start simple and build up.
Each test is implemented as a go test with //go:generate statements
//go:generate protoc -I=. --go_out=. repeated.proto
//go:generate protoc -I=. --bson_out=. repeated.proto


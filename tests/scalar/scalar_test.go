//go:generate protoc -I=. --go_out=. scalar.proto
//go:generate protoc -I=. --bson_out=. scalar.proto

package repeated

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/b3s4r/protoc-gen-bson/tests/scalar/proto/scalar"
)

// TestScalar checks some scalar types, since they are simple type checked copies
// there is no need to check all types. The ones added are only here as a sanity check./
func TestScalar(t *testing.T) {
	src := &pb.ScalarMessage{
		Double: 123.345,
		Float: 234.456,
		Int32: 345,
		Int64: 456,
	}
	t.Log(src)
	
	got := pb.ScalarMessageToBSON(src)
	t.Log(got)

	want := bson.M{
		"Double": 123.345,
		"Float": float32(234.456),
		"Int32": int32(345),
		"Int64": int64(456),
	}
	t.Log(want)
	
	if !reflect.DeepEqual(got, want) {
		t.Fatal()
	}
}

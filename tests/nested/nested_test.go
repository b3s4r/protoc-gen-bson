//go:generate protoc -I=. --go_out=. nested.proto
//go:generate protoc -I=. --bson_out=. nested.proto

package repeated

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/b3s4r/protoc-gen-bson/tests/nested/proto/nested"
)


func TestNested(t *testing.T) {
	tests := []struct {
		src *pb.NestedMessage
		want bson.M
	} {
		{
			&pb.NestedMessage{},
			bson.M{
				"Foo": bson.M{"Text":""},
				"Bar": bson.M{
					"Foo": bson.M{"Text":""},
				},
			},
		},
		{
			&pb.NestedMessage{
				Foo: &pb.Foo{Text: "hello"},
			},
			bson.M{
				"Foo": bson.M{
					"Text": "hello",
				},
				"Bar": bson.M{
					"Foo": bson.M{
						"Text":"",
					},
				},
			},
		},
		{
			&pb.NestedMessage{
				Foo: &pb.Foo{Text: "hello"},
				Bar: &pb.Bar{
					Foo: &pb.Foo{Text: "world"},
				},		
			},
			bson.M{
				"Foo": bson.M{
					"Text": "hello",
				},
				"Bar": bson.M{
					"Foo": bson.M{
						"Text":"world",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Logf("%+v", tc.src)
		got := pb.NestedMessageToBSON(tc.src)
		
		t.Logf("%+v", got)
		t.Logf("%+v", tc.want)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatal()
		}
	}
}


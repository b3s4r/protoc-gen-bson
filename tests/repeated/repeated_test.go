//go:generate protoc -I=. --go_out=. repeated.proto
//go:generate protoc -I=. --bson_out=. repeated.proto

package repeated

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/b3s4r/protoc-gen-bson/tests/repeated/proto/repeated"
)

func TestRepeated(t *testing.T) {
	tests := []struct {
		src *pb.RepeatedMessage
		want bson.M
	} {
		{
			&pb.RepeatedMessage{Tests: []string{"hello", "world", "!!!"}},
			bson.M{
				"Tests": bson.A{"hello", "world", "!!!"},
				"Foos": bson.A{},
				"Bars": bson.A{},
			},
		},
 		{
			&pb.RepeatedMessage{
				Tests: []string{"hello", "world", "!!!"},
				Foos: []*pb.Foo{
					&pb.Foo{Text:"hello"},
					&pb.Foo{Text:"world"},
				},
			},
			bson.M{
				"Tests": bson.A{"hello", "world", "!!!"},
				"Foos": bson.A{
					bson.M{"Text":"hello"},
					bson.M{"Text":"world"},
				},
				"Bars": bson.A{},
			},
		},
		{
			&pb.RepeatedMessage{
				Tests: []string{"hello", "world", "!!!"},
				Foos: []*pb.Foo{
					&pb.Foo{Text:"hello"},
					&pb.Foo{Text:"world"},
				},
				Bars: []*pb.Bar{
					&pb.Bar{Foo:&pb.Foo{Text: "olleh"}},
					&pb.Bar{Foo:&pb.Foo{Text: "dlrow"}},
				},
			},
			bson.M{
				"Tests": bson.A{"hello", "world", "!!!"},
				"Foos": bson.A{
					bson.M{"Text":"hello"},
					bson.M{"Text":"world"},
				},
				"Bars": bson.A{
					bson.M{"Foo":bson.M{"Text":"olleh"}},
					bson.M{"Foo":bson.M{"Text":"dlrow"}},
				},
			},
		},
	}
	
	for _, tc := range tests {
		t.Log(tc.src)
		got := pb.RepeatedMessageToBSON(tc.src)
		t.Log(got)

		t.Log(tc.want)
	
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatal()
		}
	}

	

}

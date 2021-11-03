//go:generate protoc -I=. --go_out=. map.proto
//go:generate protoc -I=. --bson_out=. map.proto

package tests

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/b3s4r/protoc-gen-bson/tests/map/proto/map"
)

func TestMap(t *testing.T) {
	tests := []struct {
		src *pb.MapMessage
		want bson.M
	} {
		{
			&pb.MapMessage{},
			bson.M{
				"Strings": bson.M{},
				"Strings2": bson.M{},
			},
		},
		{
			&pb.MapMessage{
				Strings: map[string]string{"hello": "world", "cruel": "place"},
			},
			bson.M{
				"Strings": bson.M{"hello": "world", "cruel": "place"},
				"Strings2": bson.M{},
			},
		},
	}
	
	for _, tc := range tests {
		t.Log(tc.src)
	
		got := pb.MapMessageToBSON(tc.src)
		t.Log(got)
		
		t.Log(tc.want)
	
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatal()
		}
	}
}

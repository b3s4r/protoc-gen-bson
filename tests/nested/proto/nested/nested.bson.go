package nested

import (
	bson "go.mongodb.org/mongo-driver/bson"
)

func FooToBSON(msg *Foo) bson.M {
	res := bson.M{}
	if msg == nil {
		msg = &Foo{}
	}
	res["Text"] = msg.GetText() //a
	return res
}
func BarToBSON(msg *Bar) bson.M {
	res := bson.M{}
	if msg == nil {
		msg = &Bar{}
	}
	res["Foo"] = FooToBSON(msg.Foo)
	return res
}
func NestedMessageToBSON(msg *NestedMessage) bson.M {
	res := bson.M{}
	if msg == nil {
		msg = &NestedMessage{}
	}
	res["Foo"] = FooToBSON(msg.Foo)
	res["Bar"] = BarToBSON(msg.Bar)
	return res
}

package repeated

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
func RepeatedMessage_Tests_ArrayToBSON(values []string) bson.A {
	res := bson.A{}
	for _, v := range values {
		res = append(res, v)
	}
	return res
}

func RepeatedMessage_Foos_ArrayToBSON(values []*Foo) bson.A {
	res := bson.A{}
	for _, v := range values {
		res = append(res, FooToBSON(v))
	}
	return res
}

func RepeatedMessage_Bars_ArrayToBSON(values []*Bar) bson.A {
	res := bson.A{}
	for _, v := range values {
		res = append(res, BarToBSON(v))
	}
	return res
}

func RepeatedMessageToBSON(msg *RepeatedMessage) bson.M {
	res := bson.M{}
	if msg == nil {
		msg = &RepeatedMessage{}
	}
	res["Tests"] = RepeatedMessage_Tests_ArrayToBSON(msg.GetTests())
	res["Foos"] = RepeatedMessage_Foos_ArrayToBSON(msg.GetFoos())
	res["Bars"] = RepeatedMessage_Bars_ArrayToBSON(msg.GetBars())
	return res
}

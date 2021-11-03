package _map

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
func MapMessage_Strings_MapToBSON(values map[string]string) bson.M {
	res := bson.M{}
	for k, v := range values {
		res[k] = v
	}
	return res
}

func MapMessage_Strings2_MapToBSON(values map[string]int32) bson.M {
	res := bson.M{}
	for k, v := range values {
		res[k] = v
	}
	return res
}

func MapMessageToBSON(msg *MapMessage) bson.M {
	res := bson.M{}
	if msg == nil {
		msg = &MapMessage{}
	}
	res["Strings"] = MapMessage_Strings_MapToBSON(msg.GetStrings())
	res["Strings2"] = MapMessage_Strings2_MapToBSON(msg.GetStrings2())
	return res
}

package scalar

import (
	bson "go.mongodb.org/mongo-driver/bson"
)

func ScalarMessageToBSON(msg *ScalarMessage) bson.M {
	res := bson.M{}
	if msg == nil {
		msg = &ScalarMessage{}
	}
	res["Double"] = msg.GetDouble() //a
	res["Float"] = msg.GetFloat()   //a
	res["Int32"] = msg.GetInt32()   //a
	res["Int64"] = msg.GetInt64()   //a
	return res
}

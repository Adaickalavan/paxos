package document

import (
	"gopkg.in/mgo.v2/bson"
)

//Message properties
type Message struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Message string        `bson:"message" json:"message"`
	Digest  string        `bson:"digest" json:"digest"`
}

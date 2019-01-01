package document

import "gopkg.in/mgo.v2/bson"

//Message properties
type Message struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Value string        `bson:"value" json:"value"`
	Hash  string        `bson:"hash" json:"hash"`
}

package database

import (
	"document"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Product contains details of Product and its Database & Collection settings
type Product struct {
	connection
}

// EnsureIndex creates an index field in the collection
func (prod *Product) EnsureIndex(fields []string) {
	index := mgo.Index{
		Key:        fields, // Index key fields; prefix name with (-) dash for descending order
		Unique:     true,   // Prevent two documents from having the same key
		DropDups:   true,   // Drop documents with same index
		Background: true,   // Build index in background and return immediately
		Sparse:     true,   // Only index documents containing the Key fields
	}
	err := prod.c.EnsureIndex(index)
	checkError(err)
}

// FindByValue retrieves the Documents by its Value from Product
func (prod *Product) FindByValue(value string) (document.Message, error) {
	var doc document.Message
	err := prod.c.Find(bson.M{"digest": value}).One(&doc)
	return doc, err
}

// Insert inserts the Document into the Product
func (prod *Product) Insert(doc document.Message) error {
	err := prod.c.Insert(&doc)
	return err
}

//checkError prints non-nil error and returns true. For nil error, returns false.
func checkError(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}

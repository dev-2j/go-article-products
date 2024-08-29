package tox

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func BsonM(v interface{}) (doc *bson.M, err error) {

	if v == nil {
		return nil, fmt.Errorf(`tox.BsonM, v is nil`)
	}

	data, err := bson.Marshal(v)
	if err != nil {
		err = fmt.Errorf(`bson.Marshal: %v`, err.Error())
		return
	}

	err = bson.Unmarshal(data, &doc)
	if err != nil {
		err = fmt.Errorf(`bson.Unmarshal: %v`, err.Error())
		return
	}

	return

}

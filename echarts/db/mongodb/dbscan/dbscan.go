package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Dbscan struct {
	Id        				bson.ObjectId	`bson:"_id"`
	Point  					string			`bson:"Point"`
	HorizontalCoordinates	[]int64        	`bson:"HorizontalCoordinates"`
	VerticalCoordinates	    []int64 		`bson:"VerticalCoordinates"`
}

func (dbscan Dbscan) ToString() string {
	return fmt.Sprintf("%#v", dbscan)
}



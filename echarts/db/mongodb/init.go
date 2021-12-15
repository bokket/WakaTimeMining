package mongodb

import (
	"gopkg.in/mgo.v2"
)

const URL string = "127.0.0.1:27017"

var (
	c *mgo.Collection
	session *mgo.Session
)

func init() {
	session, _ = mgo.Dial(URL)
	//切换到数据库
	db := session.DB("wakaTime")
	//切换到collection
	c = db.C("dbscan")
}
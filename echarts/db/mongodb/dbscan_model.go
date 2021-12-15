package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	mongodb "wakever/charts/db/mongodb/dbscan"
)

/*
//新增数据
func add() {
	//    defer session.Close()
	stu1 := new(User)
	stu1.Id = bson.NewObjectId()
	stu1.Username = "stu1_name"
	stu1.Pass = "stu1_pass"

	stu1.Regtime = time.Now().Unix()
	stu1.Interests = []string{"象棋", "游泳", "跑步"}
	err := c.Insert(stu1)

	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}
}
*/



//查询
func Find(idStr string)(dbscan *mongodb.Dbscan,err error) {
	/*var users []User

	c.Find(bson.M{"name": "stu1_name"}).All(&users)
	for _, value := range users {
		fmt.Println(value.ToString())
	}*/
	//根据ObjectId进行查询

	//idStr := "61b18bbe931c8c5510d52ed7"
	objectId := bson.ObjectIdHex(idStr)
	dbscan = new(mongodb.Dbscan)
	err=c.Find(bson.M{"_id": objectId}).One(dbscan)
	return
}

/*
//根据id进行修改
func update() {
	interests := []string{"象棋", "游泳", "跑步"}
	err := c.Update(bson.M{"_id": bson.ObjectIdHex("577fb2d1cde67307e819133d")}, bson.M{"$set": bson.M{
		"name":      "修改后的name",
		"pass":      "修改后的pass",
		"regtime":   time.Now().Unix(),
		"interests": interests,
	}})
	if err != nil {
		fmt.Println("修改失败")
	} else {
		fmt.Println("修改成功")
	}
}

//删除
func del() {
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex("577fb2d1cde67307e819133d")})
	if err != nil {
		fmt.Println("删除失败" + err.Error())
	} else {
		fmt.Println("删除成功")
	}
}
func main() {
	add()
	find()
	update()
	del()
}
*/

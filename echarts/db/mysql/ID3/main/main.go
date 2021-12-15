package main

import (
	"github.com/kelseyhightower/confd/log"
	"wakever/charts/db/mysql"
)

func main() {
	i, err := mysql.GetAllID3()
	if err != nil {
		log.Error("%v", err)
	}
	for _, u := range i {
		println(u.Id)

		println(u.Lchild)
		println(u.Rchild)
		println(u.Lvalue)
		println(u.Rvalue)
		println(u.TableName())
	}
}

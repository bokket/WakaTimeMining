package main

import (
	"wakever/charts/db/mysql"

	"github.com/kelseyhightower/confd/log"
)

func main() {
	//type Result=map[int64]string
	////re=make(Result,8)
	//
	//re := make(Result, 0)

	r, err := mysql.GetAllKmeans()
	if err != nil {
		log.Error("%v", err)
	}
	for _, u := range r {
		println(u.Id)
		println(u.HorizontalCoordinates)
		println(u.VerticalCoordinates)
		println(u.ThreeDimensionalCoordinates)
		println(u.TableName())
	}
	//println(re)
}
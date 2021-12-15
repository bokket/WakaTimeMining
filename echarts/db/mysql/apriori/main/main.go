package main

import (
	"wakever/charts/db/mysql"

	util "wakever/utils"

	"github.com/kelseyhightower/confd/log"
)

func main() {
	//type Result=map[int64]string
	////re=make(Result,8)
	//
	//re := make(Result, 0)

	r, err := mysql.GetAllApriori()
	if err != nil {
		log.Error("%v", err)
	}
	for _, u := range r {
		println(u.Id)
		//f:=strings.TrimPrefix(u.FrequentSet,"frozenset({")
		//f=strings.TrimSuffix(f,"})")

		f:=util.GetAbsString(u.FrequentSet,"frozenset({","})")
		c:=util.GetAbsString(u.ConsequentSet,"frozenset({","})")

		println("FrequentSet:",f)
		println("ConsequentSet:",c)
		println(u.Credibility)
		println(u.TableName())
	}
	//println(re)
}

package main

import (
	"github.com/kelseyhightower/confd/log"
	"wakever/charts/db/mysql"
)

func main()  {
	type Result=map[int64]string
	//re=make(Result,8)

	re := make(Result, 0)

	r,err:=mysql.GetAllUsers()
	if err!=nil {
		log.Error("%v",err)
	}
	for _, u := range r {
		println(u.SSE)
		println(u.Id)
		println(u.TableName())

		re[u.Id]=u.SSE
	}
	println(re)

	//type Result=map[int64]string

	/*re,err:=mysql.ScanUser()
	if err!=nil {
		log.Error("%v",err)
	}
	println(re)

	 */
}


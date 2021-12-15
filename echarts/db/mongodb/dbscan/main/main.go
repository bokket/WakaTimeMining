package main

import (
	"fmt"

	"wakever/charts/db/mongodb"
)

func main()  {
	user,err:=mongodb.Find("61b6aac00ba62f26fc26c713")
	if err!=nil {
		panic(err)
	}

	fmt.Println(user.Id)
	fmt.Println(user.Point)
	fmt.Println(user.HorizontalCoordinates)
	fmt.Println(user.VerticalCoordinates)
	fmt.Println(user.ToString())
}
package mysql

import (
	"fmt"

	"wakever/charts/db/mysql/SSE"
)

func CreateUser(sse *mysql.SSE) (err error) {
	err = DB().Create(sse).Error

	return
}

func GetUserById(userId int64) (sse *mysql.SSE, err error) {
	sse = new(mysql.SSE)
	err = DB().Where("id = ?", userId).First(sse).Error

	return
}

func GetAllSSEs() (sses []*mysql.SSE, err error) {
	err = DB().Find(&sses).Error
	return
}

//func GetUserBySSEAndId (SSE string) (user *table.User, err error) {
func GetUserBySSEAndId (sseStr string,SSEId  int64) (user *mysql.SSE, err error) {
	user = new(mysql.SSE)
	err = DB().Where("SSE = ? AND id=?", sseStr,SSEId).
		  First(&user).Error
	fmt.Printf("%#v\n", user)

	return
}

func UpdateUserNameById(sseStr string, SSEId int64) (err error) {
	user := new(mysql.SSE)
	err = DB().Where("id = ?", SSEId).First(user).Error
	if err != nil {
		return
	}

	user.SSE = sseStr
	err = DB().Save(user).Error

	return
}

func DeleteUserById(userId int64) (err error) {
	user := new(mysql.SSE)
	err = DB().Where("id = ?", userId).First(user).Error
	if err != nil {
		return
	}
	err = DB().Delete(user).Error

	return
}

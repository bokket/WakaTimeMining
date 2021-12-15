package mysql

import mysql "wakever/charts/db/mysql/ID3"

func CreateID3(ID3 *mysql.ID3) (err error) {
	err=DB().Create(ID3).Error

	return
}

func GetID3ById(userId int64) (ID3 *mysql.ID3, err error) {
	ID3 = new(mysql.ID3)
	err = DB().Where("id = ?", userId).First(ID3).Error

	return
}

func GetAllID3() (ID3 []*mysql.ID3, err error) {
	err = DB().Find(&ID3).Error
	return
}

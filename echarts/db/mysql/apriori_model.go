package mysql

import (
	"wakever/charts/db/mysql/apriori"
)

func CreateApriori(apriori *mysql.Apriori) (err error) {
	err = DB().Create(apriori).Error

	return
}

func GetAprioriById(userId int64) (apriori *mysql.Apriori, err error) {
	apriori = new(mysql.Apriori)
	err = DB().Where("id = ?", userId).First(apriori).Error

	return
}

func GetAllApriori() (apriori []*mysql.Apriori, err error) {
	err = DB().Find(&apriori).Error
	return
}



package mysql

import (
	"wakever/charts/db/mysql/kmeans"
)

func CreateKmeans(kmeans *mysql.Kmeans) (err error) {
	err = DB().Create(kmeans).Error

	return
}

func GetKmeansById(userId int64) (kmean *mysql.Kmeans, err error) {
	kmean = new(mysql.Kmeans)
	err = DB().Where("id = ?", userId).First(kmean).Error

	return
}

func GetAllKmeans() (kmeans []*mysql.Kmeans, err error) {
	err = DB().Find(&kmeans).Error
	return
}
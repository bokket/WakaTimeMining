package mysql

type Kmeans struct {
	Id   						int64     	`gorm:"column:id;primary_key"`
	HorizontalCoordinates		string    	`gorm:"column:HorizontalCoordinates"`
	VerticalCoordinates			string   	`gorm:"column:VerticalCoordinates"`
	ThreeDimensionalCoordinates	string   	`gorm:"column:ThreeDimensionalCoordinates"`
}

// TableName sets the insert table name for this struct type
func (m *Kmeans) TableName() string {
	return "kmeans"
}

package mysql

type SSE struct {
	Id   int64     `gorm:"column:id;primary_key"`
	SSE  string    `gorm:"column:SSE"`
	//Secret    string    `gorm:"column:secret;type:varchar(1000)"`
	//CreatedAt time.Time `gorm:"column:created_at"`
	//UpdatedAt time.Time `gorm:"column:updated_at"`
}


// TableName sets the insert table name for this struct type
func (m *SSE) TableName() string {
	return "SSE"
}
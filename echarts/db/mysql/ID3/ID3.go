package mysql


type ID3 struct {
	Id   		int64     	`gorm:"column:id;primary_key"`
	Root		string    	`gorm:"column:root"`
	Rchild		string   	`gorm:"column:rchild"`
	Lchild		string   	`gorm:"column:lchild"`
	Rvalue  	int			`gorm:"column:rvalue"`
	Lvalue  	int			`gorm:"column:lvalue"`
}


// TableName sets the insert table name for this struct type
func (m *ID3) TableName() string {
	return "ID3"
}

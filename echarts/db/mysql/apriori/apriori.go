package mysql

type Apriori struct {
	Id   int64     			`gorm:"column:id;primary_key"`
	FrequentSet  string    	`gorm:"column:frequentSet"`
	ConsequentSet string   	`gorm:"column:consequentSet"`
	Credibility   string   	`gorm:"column:credibility"`
}

// TableName sets the insert table name for this struct type
func (m *Apriori) TableName() string {
	return "rule"
}

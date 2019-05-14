package models

// Ghost -> represents our test model
type Ghost struct {
	ID   int    `gorm:"type:smallint;primary key" json:"id"`
	Name string `gorm:"type:varchar(36);not null" json:"name"`
	Age  int    `gorm:"type:smallint;not null" json:"age"`
}

// TableName -> Overrides orm's default behaviour of using plurals for table names
func (ghost *Ghost) TableName() string {
	return "ghost"
}

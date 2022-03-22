package model

import "time"

func (User) TableName() string {
	return "user"
}

type User struct {
	Id         int64  `gorm:"column:id;auto_increment;primary_key"`
	Name       string `gorm:"column:name;type:varchar(50)"`
	Mobile     string `gorm:"column:mobile;type:varchar(50)"`
	Age        int    `gorm:"column:age;type:int(11)"`
	CreateTime time.Time
	State      int
}

type UserParams struct {
	Id, Age, State int
	Name, Mobile   string
}

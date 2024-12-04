package entity

type TaskDetail struct {
	Id          string `gorm:"type:varcahr(100) NOT NULL;primary_key;" json:"id" validate:"required,min=5,max=40"`
	Description string `gorm:"type:varcahr(100)" json:"description"`
	Status      string `gorm:"type:varcahr(100)" json:"status"`
	DateTime    string `gorm:"type:varcahr(100)" json:"dateTime"`
}
type TaskLog struct {
	Id          string `gorm:"type:varcahr(100) NOT NULL;primary_key;" json:"id"`
	Description string `gorm:"type:varcahr(100)" json:"description"`
	DateTime    string `gorm:"type:varcahr(100)" json:"dateTime"`
}

func (TaskDetail) TableName() string {
	return "task_detail"
}
func (TaskLog) TableName() string {
	return "task_log"
}

var TaskChan chan TaskDetail = make(chan TaskDetail, 1000)

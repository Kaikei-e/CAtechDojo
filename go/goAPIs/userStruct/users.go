package userstruct

type Users struct{
	Id int `gorm:"column:id"`
	Name string `json:"name"`
	Token string `json:"token"`
}

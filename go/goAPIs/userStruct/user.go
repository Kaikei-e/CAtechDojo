package userstruct

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Token string `json:"token"`
}

package entity

type User struct {
	Id         int64
	Name       string   `json:"name" binding:"required"`
	Password   string   `json:"password" binding:"required"`
	Role       string   `json:"role" binding:"required"`
	ResultTest []string `json:"result_test"`
	Students   []string `json:"students"`
}

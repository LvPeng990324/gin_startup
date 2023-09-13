package model

type User struct {
	Id       uint   `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`

	IsDeleted bool `json:"is_deleted" db:"is_deleted"`

	CreatedAt string `json:"created_at" db:"created_at"`
}

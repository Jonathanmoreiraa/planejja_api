package domain

type Users struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

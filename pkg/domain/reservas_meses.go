package domain

type ReservasMeses struct {
	ID    uint   `json:"id" gorm:"unique;not null"`
	Valor string `json:"valor"`
}

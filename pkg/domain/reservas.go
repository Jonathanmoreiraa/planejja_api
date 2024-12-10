package domain

type Reservas struct {
	ID   uint   `json:"id" gorm:"unique;not null"`
	Meta string `json:"description"`
	Tipo uint   `json:"tipo"`
}

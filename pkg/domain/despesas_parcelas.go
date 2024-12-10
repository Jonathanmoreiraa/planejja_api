package domain

type DespesasParcelas struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	DespesaId uint   `json:"despesa_id"`
	Mes       string `json:"mes"`
	Valor     string `json:"valor"`
}

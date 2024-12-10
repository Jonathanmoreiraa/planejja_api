package domain

type Despesas struct {
	ID             uint   `json:"id" gorm:"unique;not null"`
	Descricao      string `json:"description"`
	Valor          string `json:"valor"`
	Status         string `json:"status"`
	DataVencimento string `json:"data_vencimento"`
}

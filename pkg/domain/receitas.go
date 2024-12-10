package domain

type Receitas struct {
	ID              uint   `json:"id" gorm:"unique;not null"`
	Descricao       string `json:"description"`
	Valor           string `json:"valor"`
	Status          uint   `json:"status"`
	DataRecebimento string `json:"data_recebimento"`
}

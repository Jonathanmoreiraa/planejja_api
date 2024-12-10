package domain

type Categorias struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	Descricao string `json:"description"`
	Tipo      string `json:"tipo"`
}

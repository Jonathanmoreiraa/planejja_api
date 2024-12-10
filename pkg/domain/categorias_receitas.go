package domain

type CategoriasReceitas struct {
	ID          uint `json:"id" gorm:"unique;not null"`
	ReceitaId   uint `json:"receita_id"`
	CategoriaId uint `json:"categoria_id"`
}

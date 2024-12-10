package domain

type CategoriasDespesas struct {
	ID          uint `json:"id" gorm:"unique;not null"`
	DespesaId   uint `json:"despesa_id"`
	CategoriaId uint `json:"categoria_id"`
}

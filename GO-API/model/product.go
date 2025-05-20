package model

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Estoque   int     `json:"estoque"`
	Categoria string  `json:"categoria"`
	Descricao string  `json:"descricao"`
	Ativo     bool    `json:"ativo"`
}

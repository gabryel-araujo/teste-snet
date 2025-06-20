package models

type Lojas struct {
	ID int `json:"id"`
	Numero_loja int `json:"numero_loja"`
	Nome string `json:"nome"`
	Razao_social string `json:"razao_social"`
	Endereco string `json:"endereco"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Cep string `json:"cep"`
	Numero string `json:"numero"`
	Estabelecimento_id int `json:"estabelecimento_id"`
}
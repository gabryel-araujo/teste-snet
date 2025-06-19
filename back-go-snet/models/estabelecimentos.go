package models

type Estabelecimentos struct {
	ID int `json:"id"`
	Numero_estabelecimento int `json:"numero_estabelecimento"`
	Nome string `json:"nome"`
	Razao_social string `json:"razao_social"`
	Endereco string `json:"endereco"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Cep string `json:"cep"`
	Numero string `json:"numero"`
}
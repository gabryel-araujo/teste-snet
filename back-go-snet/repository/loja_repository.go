package repository

import (
	"database/sql"

	"github.com/gabryel-araujo/back-go-snet/models"
)

type StoreRepository struct{
	DB *sql.DB
}

func (r *StoreRepository) CreateStore(l *models.Lojas) error{
	err := r.DB.QueryRow("INSERT INTO lojas (nome,numero_loja,razao_social,endereco,cidade,estado,cep,numero,estabelecimento_id) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id",l.Nome,l.Numero_loja,l.Razao_social,l.Endereco,l.Cidade,l.Estado,l.Cep,l.Numero,l.Estabelecimento_id).Scan(&l.ID)

	return err
}

func (r *StoreRepository) GetAllStores() ([]models.Lojas, error) {
	rows, err := r.DB.Query(`
		SELECT id, nome, numero_loja, razao_social, endereco, cidade, estado, cep, numero, estabelecimento_id
		FROM lojas
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.Lojas

	for rows.Next() {
		var l models.Lojas
		if err := rows.Scan(&l.ID, &l.Nome, &l.Numero_loja, &l.Razao_social, &l.Endereco, &l.Cidade, &l.Estado, &l.Cep, &l.Numero, &l.Estabelecimento_id); err != nil {
			return nil, err
		}
		stores = append(stores, l)
	}

	return stores, nil
}

func (r *StoreRepository) GetOneStore(id string) (*models.Lojas, error) {
	var l models.Lojas
	err := r.DB.QueryRow(`
		SELECT id, nome, numero_loja, razao_social, endereco, cidade, estado, cep, numero, estabelecimento_id
		FROM lojas WHERE id = $1
	`, id).Scan(&l.ID, &l.Nome, &l.Numero_loja, &l.Razao_social, &l.Endereco, &l.Cidade, &l.Estado, &l.Cep, &l.Numero, &l.Estabelecimento_id)

	if err == sql.ErrNoRows {
		return nil, nil 
	} else if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *StoreRepository) UpdateStore(id string, l *models.Lojas) (bool, error) {
	res, err := r.DB.Exec(`
		UPDATE lojas 
		SET nome=$1, numero_loja=$2, razao_social=$3, endereco=$4, cidade=$5, estado=$6, cep=$7, numero=$8, estabelecimento_id=$9
		WHERE id=$10
	`, l.Nome, l.Numero_loja, l.Razao_social, l.Endereco, l.Cidade, l.Estado, l.Cep, l.Numero, l.Estabelecimento_id, id)

	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *StoreRepository) DeleteStore(id string) (bool, error) {
	res, err := r.DB.Exec("DELETE FROM lojas WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
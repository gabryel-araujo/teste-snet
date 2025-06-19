package repository

import (
	"database/sql"

	"github.com/gabryel-araujo/back-go-snet/models"
)

type EstabRepository struct{
	DB *sql.DB
}

func (r *EstabRepository) CreateEstab(e *models.Estabelecimentos) error{
	err := r.DB.QueryRow("INSERT INTO estabelecimentos (nome,numero_estabelecimento,razao_social,endereco,cidade,estado,cep,numero) values ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id",e.Nome,e.Numero_estabelecimento,e.Razao_social,e.Endereco,e.Cidade,e.Estado,e.Cep,e.Numero).Scan(&e.ID)

	return err
}

func (r *EstabRepository) GetAllEstabs() ([]models.Estabelecimentos, error) {
	rows, err := r.DB.Query("SELECT * FROM estabelecimentos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var estabs []models.Estabelecimentos

	for rows.Next() {
		var e models.Estabelecimentos
		if err := rows.Scan(&e.ID, &e.Nome, &e.Numero_estabelecimento, &e.Razao_social, &e.Endereco, &e.Cidade, &e.Estado, &e.Cep, &e.Numero); err != nil {
			return nil, err
		}
		estabs = append(estabs, e)
	}

	return estabs, nil
}

func (r *EstabRepository) GetOneEstab(id string) (*models.Estabelecimentos, error) {
	var e models.Estabelecimentos
	err := r.DB.QueryRow(`
		SELECT id, nome, numero_estabelecimento, razao_social, endereco, cidade, estado, cep, numero FROM estabelecimentos WHERE id = $1
	`, id).Scan(&e.ID, &e.Nome, &e.Numero_estabelecimento, &e.Razao_social, &e.Endereco, &e.Cidade, &e.Estado, &e.Cep, &e.Numero)

	if err == sql.ErrNoRows {
		return nil, nil 
	} else if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *EstabRepository) UpdateEstab(id string, e *models.Estabelecimentos) (bool, error) {
	res, err := r.DB.Exec(`
		UPDATE estabelecimentos 
		SET nome=$1, numero_estabelecimento=$2, razao_social=$3, endereco=$4, cidade=$5, estado=$6, cep=$7, numero=$8
		WHERE id=$9
	`, e.Nome, e.Numero_estabelecimento, e.Razao_social, e.Endereco, e.Cidade, e.Estado, e.Cep, e.Numero, id)

	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *EstabRepository) DeleteEstab(id string) (bool, error) {
	res, err := r.DB.Exec("DELETE FROM estabelecimentos WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
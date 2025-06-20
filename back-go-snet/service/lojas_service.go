package service

import (
	"errors"
	"fmt"

	"github.com/gabryel-araujo/back-go-snet/models"
	"github.com/gabryel-araujo/back-go-snet/repository"
)

type StoreService struct {
	Repo *repository.StoreRepository
}

func (s *StoreService) CreateStore(l *models.Lojas) error {
	fmt.Println("Criando loja...")

	err := s.Repo.CreateStore(l)
	if err != nil {
		fmt.Println("Erro ao criar loja:", err)
		return err
	}

	fmt.Printf("Loja criada com sucesso: %+v\n", l)
	return nil
}

func (s *StoreService) GetAllStores() ([]models.Lojas, error) {
	fmt.Println("Listando todas as lojas...")
	return s.Repo.GetAllStores()
}

func (s *StoreService) GetOneStore(id string) (*models.Lojas, error) {
	fmt.Println("Buscando loja...")

	loja, err := s.Repo.GetOneStore(id)
	if err != nil {
		fmt.Println("Erro ao buscar loja:", err)
		return nil, err
	}
	if loja == nil {
		fmt.Println("Loja não encontrada.")
		return nil, errors.New("loja não encontrada")
	}

	fmt.Printf("Loja encontrada: %+v\n", loja)
	return loja, nil
}

func (s *StoreService) UpdateStore(id string, l *models.Lojas) error {
	fmt.Println("Atualizando loja...")

	updated, err := s.Repo.UpdateStore(id, l)
	if err != nil {
		fmt.Println("Erro ao atualizar loja:", err)
		return err
	}
	if !updated {
		fmt.Println("Loja não encontrada ou não alterada.")
		return errors.New("loja não encontrada ou não alterada")
	}

	fmt.Printf("Loja atualizada com sucesso: %+v\n", l)
	return nil
}

func (s *StoreService) DeleteStore(id string) error {
	fmt.Println("Removendo loja...")

	deleted, err := s.Repo.DeleteStore(id)
	if err != nil {
		fmt.Println("Erro ao remover loja:", err)
		return err
	}
	if !deleted {
		fmt.Println("Loja não encontrada para remoção.")
		return errors.New("loja não encontrada")
	}

	fmt.Println("Loja removida com sucesso.")
	return nil
}

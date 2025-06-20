package service

import (
	"errors"
	"fmt"

	"github.com/gabryel-araujo/back-go-snet/models"
	"github.com/gabryel-araujo/back-go-snet/repository"
)

type EstabService struct {
	Repo *repository.EstabRepository
}

func (s *EstabService) CreateEstab(e *models.Estabelecimentos) error {
	fmt.Println("Criando estabelecimento...")

	err := s.Repo.CreateEstab(e)
	if err != nil {
		fmt.Println("Erro ao criar estabelecimento:", err)
		return err
	}

	fmt.Printf("Estabelecimento criado com sucesso: %+v\n", e)
	return nil
}

func (s *EstabService) GetAllEstabs() ([]models.Estabelecimentos, error) {
	fmt.Println("Listando todos os estabelecimentos...")
	return s.Repo.GetAllEstabs()
}

func (s *EstabService) GetOneEstab(id string) (*models.Estabelecimentos, error) {
	fmt.Println("Buscando estabelecimento...")

	estab, err := s.Repo.GetOneEstab(id)
	if err != nil {
		fmt.Println("Erro ao buscar estabelecimento:", err)
		return nil, err
	}
	if estab == nil {
		fmt.Println("Estabelecimento não encontrado.")
		return nil, errors.New("estabelecimento não encontrado")
	}

	fmt.Printf("Estabelecimento encontrado: %+v\n", estab)
	return estab, nil
}

func (s *EstabService) UpdateEstab(id string, e *models.Estabelecimentos) error {
	fmt.Println("Atualizando estabelecimento...")

	updated, err := s.Repo.UpdateEstab(id, e)
	if err != nil {
		fmt.Println("Erro ao atualizar estabelecimento:", err)
		return err
	}
	if !updated {
		fmt.Println("Estabelecimento não encontrado ou não alterado.")
		return errors.New("estabelecimento não encontrado ou não alterado")
	}

	fmt.Printf("Estabelecimento atualizado com sucesso: %+v\n", e)
	return nil
}

func (s *EstabService) DeleteEstab(id string) error {
	fmt.Println("Removendo estabelecimento...")

	deleted, err := s.Repo.DeleteEstab(id)
	if err != nil {
		fmt.Println("Erro ao remover estabelecimento:", err)
		return err
	}
	if !deleted {
		fmt.Println("Estabelecimento não encontrado para remoção.")
		return errors.New("estabelecimento não encontrado")
	}

	fmt.Println("Estabelecimento removido com sucesso.")
	return nil
}

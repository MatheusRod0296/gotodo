package todo

import (
	"errors"
	"strconv"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) List() []Todo {
	return s.repo.GetAll()
}

func (s *Service) Create(title string) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("title é obrigatório")
	}

	todo := Todo{
		Title:     title,
		Completed: false,
	}

	return s.repo.Create(todo), nil
}

func (s *Service) Update(idStr, title string, completed bool) (Todo, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Todo{}, errors.New("id inválido")
	}

	updated := Todo{
		Title:     title,
		Completed: completed,
	}

	return s.repo.Update(id, updated)
}

func (s *Service) Delete(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("id inválido")
	}

	return s.repo.Delete(id)
}

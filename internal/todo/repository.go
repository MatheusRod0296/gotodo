package todo

import "errors"

type Repository struct {
	data   []Todo
	lastID int
}

func NewRepository() *Repository {
	return &Repository{
		data:   []Todo{},
		lastID: 0,
	}
}

func (r *Repository) GetByID(id int) (Todo, error) {
	for _, t := range r.data {
		if t.ID == id {
			return t, nil
		}
	}
	return Todo{}, errors.New("todo não encontrado")
}

func (r *Repository) GetAll() []Todo {
	return r.data
}

func (r *Repository) Create(todo Todo) Todo {
	r.lastID++
	todo.ID = r.lastID
	r.data = append(r.data, todo)
	return todo
}

func (r *Repository) Update(id int, updated Todo) (Todo, error) {
	for i, t := range r.data {
		if t.ID == id {
			updated.ID = id
			r.data[i] = updated
			return updated, nil
		}
	}
	return Todo{}, errors.New("todo não encontrado")
}

func (r *Repository) Delete(id int) error {
	for i, t := range r.data {
		if t.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return errors.New("todo não encontrado")
}

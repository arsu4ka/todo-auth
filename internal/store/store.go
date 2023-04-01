package store

import "gorm.io/gorm"

type Store struct {
	user *UserRepo
	todo *ToDoRepo
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		user: &UserRepo{db: db},
		todo: &ToDoRepo{db: db},
	}
}

func (s *Store) User() *UserRepo {
	return s.user
}

func (s *Store) ToDo() *ToDoRepo {
	return s.todo
}

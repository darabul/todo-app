package service

import "github.com/darabul/todo-app/pkg/repository"

//тут будет организована бизнес логика приложения
//сущность из доменной зоны(todo.go)
type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

//создаем конструктор нашего сервиса
//сервис будет взаимодействовать с БД, поэтому в него будет передавать указатель на хранилище данных
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

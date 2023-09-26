package repository

//тут будет организована логика работы с БД
//сущность из доменной зоны(todo.go)
type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

//создаем конструктор нашего хранилища
func NewRepository() *Repository {
	return &Repository{}
}

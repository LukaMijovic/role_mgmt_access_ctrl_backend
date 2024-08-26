package repository

type Repository interface {
	Save() error
	Read()
	Update()
	Delete()
}

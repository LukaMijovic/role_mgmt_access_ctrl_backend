package repository

type Repository interface {
	Save()
	Read()
	Update()
	Delete()
}

package storage

type Storage interface {
	Ping()
	Close()
	Create()
	GetByID()
	GetCollection()
	Update()
	Delete()
}

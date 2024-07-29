package storage

type Storager interface {
	AddUser(id int64, slug string) error
	GetSlug(id int64) (string, error)
}

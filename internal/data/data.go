package data

type RingStorage struct {
	rings map[string]string
}

type CrudForRing interface {
	Create(key string, value string) error
	Update(key string, newValue string) error
	Delete(key string) error
	Get(key string) (string, error)
}

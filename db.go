package pkauth

type DB interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
}
package interfaces

type Storage interface {
	Get(string) (string, error)
	Set(string, string, int) error
}

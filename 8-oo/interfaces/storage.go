package interfaces

type Storage interface {
	Get(string) string
	Set(string, string, int) error
}

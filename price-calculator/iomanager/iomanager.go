package iomanager

// IOManager is an interface that defines the methods that an IOManager should implement
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

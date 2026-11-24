package iomanager

type IOManager interface {
	WriteResult(data any) error
	ReadLines() ([]float64, error)
}

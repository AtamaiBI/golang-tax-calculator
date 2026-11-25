package iomanager

/*type IOManager interface {
	WriteResult(data any) error
	ReadLines() ([]float64, error)
}*/

// Split the interface into two as a core Go idiom often summarised as 'the bigger the interface, the weaker the abstraction. By separating reading from writing cmdmanager no longer needs to kknow anything about file or filemanager.

type LineReader interface {
	ReadLines() ([]float64, error)
}

type ResultWriter interface {
	WriteResult(data any) error
}

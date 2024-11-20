package csv

// AbstractExecutor ジェネリクス構造体
type AbstractExecutor[T any] struct {
	Data      []T
	rowsLimit int
}

// NewAbstractExecutor is constructor
func NewAbstractExecutor[T any](data []T, opt ...*Option) *AbstractExecutor[T] {
	a := &AbstractExecutor[T]{
		Data:      data,
		rowsLimit: defaultCsvRowsLimit,
	}
	if len(opt) > 0 {
		o := opt[0]
		if o.RowsLimit > 0 {
			a.rowsLimit = o.RowsLimit
		}
	}
	return a
}

// Option is AbstractExecutor Option
type Option struct {
	RowsLimit int
}

const (
	defaultCsvRowsLimit int = 5000
)

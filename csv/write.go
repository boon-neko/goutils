package csv

import (
	"bytes"

	"github.com/jszwec/csvutil"
)

// WriteCsv is return csv format byte
func (a *AbstractExecutor[T]) WriteCsv() ([]byte, error) {
	var buf bytes.Buffer

	w := NewCustomCSVWriter(&buf)
	w.UseCRLF = true
	enc := csvutil.NewEncoder(w)

	for _, u := range a.Data {
		if err := enc.Encode(u); err != nil {
			return nil, err
		}
	}

	if err := w.Flush(); err != nil {
		return nil, err
	}
	if err := w.Error(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

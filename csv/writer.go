package csv

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"unicode/utf8"
)

// CustomCSVWriter is custom csv writer
type CustomCSVWriter struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

// NewCustomCSVWriter is constructor
func NewCustomCSVWriter(w io.Writer) *CustomCSVWriter {
	return &CustomCSVWriter{Comma: ',', w: bufio.NewWriter(w)}
}

// Write is flush enclose double quote record to buffer.
// custom "encoding/csv" package Writer struct method Write
func (w *CustomCSVWriter) Write(record []string) error { // nolint:funlen,gocognit
	if !validDelim(w.Comma) {
		return errors.New("csv: invalid field or comment delimiter")
	}
	for n, field := range record {
		if n > 0 {
			if _, err := w.w.WriteRune(w.Comma); err != nil {
				return err
			}
		}
		if err := w.w.WriteByte('"'); err != nil {
			return err
		}
		for len(field) > 0 {
			// Search for special characters.
			i := strings.IndexAny(field, "\"\r\n")
			if i < 0 {
				i = len(field)
			}
			// Copy verbatim everything before the special character.
			if _, err := w.w.WriteString(field[:i]); err != nil {
				return err
			}
			field = field[i:]
			// Encode the special character.
			if len(field) > 0 {
				var err error
				switch field[0] {
				case '"':
					_, err = w.w.WriteString(`""`)
				case '\r':
					if !w.UseCRLF {
						err = w.w.WriteByte('\r')
					}
				case '\n':
					if w.UseCRLF {
						_, err = w.w.WriteString("\r\n")
					} else {
						err = w.w.WriteByte('\n')
					}
				}
				field = field[1:]
				if err != nil {
					return err
				}
			}
		}
		if err := w.w.WriteByte('"'); err != nil {
			return err
		}
	}
	var err error
	if w.UseCRLF {
		_, err = w.w.WriteString("\r\n")
	} else {
		err = w.w.WriteByte('\n')
	}
	return err
}

// Flush writes any buffered data to the underlying io.Writer.
func (w *CustomCSVWriter) Flush() error { return w.w.Flush() }

// Error reports any error that has occurred during a previous Write or Flush.
func (w *CustomCSVWriter) Error() error {
	_, err := w.w.Write(nil)
	return err
}
func validDelim(r rune) bool {
	return r != 0 && r != '"' && r != '\r' && r != '\n' && utf8.ValidRune(r) && r != utf8.RuneError
}

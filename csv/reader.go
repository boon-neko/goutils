package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"reflect"

	"github.com/jszwec/csvutil"
)

// NewIgnoreDoubleQuoteByteReader is ignored double quote from byte
func NewIgnoreDoubleQuoteByteReader(b []byte) io.Reader {
	nb := make([]byte, len(b))
	i := 0
	for _, v := range b {
		if v == 34 {
			continue
		}
		nb[i] = v
		i++
	}
	return bytes.NewReader(nb[:i])
}

// ReadCsv is csv file decode struct from io.Reader
func (a *AbstractExecutor[T]) ReadCsv(r io.Reader) error {
	// *csv.Readerに変換
	reader := csv.NewReader(r)
	reader.LazyQuotes = false // ダブルクオートを厳密にチェックする。
	// *csvutil.Decoderに変換
	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		return err
	}
	i := 0
	for {
		i++
		var data T
		// 1行ずつデコード
		if err := dec.Decode(&data); err == io.EOF {
			fmt.Println("Completed read csv data.")
			break
		} else if err != nil {
			// パースエラーが起きた場合
			if e, ok := err.(*csv.ParseError); ok {
				errStr := "データ形式に誤りがあります"
				if e.Err == csv.ErrFieldCount {
					errStr += "指定されたカラム数と実際のデータのカラム数が合っていません"
				}
				return fmt.Errorf("エラー内容: %s,%s StartLine: %d, Line: %d, Column:%d", errStr, e.Err, e.StartLine, e.Line, e.Column)
			}
		}
		setRowNumber(data, i)
		a.Data = append(a.Data, data)
	}
	if a.rowsLimit > 0 && len(a.Data) > a.rowsLimit {
		return fmt.Errorf("CSVの行数が制限を超えています。upload_file_rows:%d, limit:%d", len(a.Data), a.rowsLimit)
	}
	return nil
}

func setRowNumber(v any, num int) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return
	}

	e := val.Elem()

	if e.Kind() != reflect.Struct {
		return
	}
	fv := e.FieldByName("RowNumber")
	if !fv.IsValid() {
		return
	}
	if !fv.CanSet() {
		return
	}
	if fv.Kind() != reflect.Int {
		return
	}
	fv.SetInt(int64(num))
}

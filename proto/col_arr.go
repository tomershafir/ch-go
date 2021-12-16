package proto

import (
	"github.com/go-faster/errors"
)

type Column interface {
	Result
	Input
}

// ColArr represents Array[T].
type ColArr struct {
	Offsets ColUInt64
	Data    Column
}

// Compile-time assertions for ColArr.
var (
	_ Input  = ColArr{}
	_ Result = (*ColArr)(nil)
	_ Column = (*ColArr)(nil)
)

func (c ColArr) EncodeColumn(b *Buffer) {
	c.Offsets.EncodeColumn(b)
	c.Data.EncodeColumn(b)
}

func (c ColArr) Type() ColumnType {
	return c.Data.Type().Array()
}

func (c ColArr) Rows() int {
	return len(c.Offsets)
}

func (c *ColArr) DecodeColumn(r *Reader, rows int) error {
	if err := c.Offsets.DecodeColumn(r, rows); err != nil {
		return errors.Wrap(err, "read offsets")
	}
	var size int
	if l := len(c.Offsets); l > 0 {
		// Pick last offset as total size of "elements" column.
		size = int(c.Offsets[l-1])
	}
	if err := c.Data.DecodeColumn(r, size); err != nil {
		return errors.Wrap(err, "decode data")
	}

	return nil
}

func (c *ColArr) Reset() {
	c.Offsets = c.Offsets[:0]
	c.Data.Reset()
}
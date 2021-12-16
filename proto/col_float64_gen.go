// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"github.com/go-faster/errors"
	"math"
)

// ColFloat64 represents Float64 column.
type ColFloat64 []float64

// Compile-time assertions for ColFloat64.
var (
	_ Input  = ColFloat64{}
	_ Result = (*ColFloat64)(nil)
	_ Column = (*ColFloat64)(nil)
)

// Type returns ColumnType of Float64.
func (ColFloat64) Type() ColumnType {
	return ColumnTypeFloat64
}

// Rows returns count of rows in column.
func (c ColFloat64) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColFloat64) Reset() {
	*c = (*c)[:0]
}

// NewArrFloat64 returns new Array(Float64).
func NewArrFloat64() *ColArr {
	return &ColArr{
		Data: new(ColFloat64),
	}
}

// AppendFloat64 appends slice of float64 to Array(Float64).
func (c *ColArr) AppendFloat64(data []float64) {
	d := c.Data.(*ColFloat64)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Float64 rows to *Buffer.
func (c ColFloat64) EncodeColumn(b *Buffer) {
	const size = 64 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		bin.PutUint64(
			b.Buf[offset:offset+size],
			math.Float64bits(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Float64 rows from *Reader.
func (c *ColFloat64) DecodeColumn(r *Reader, rows int) error {
	const size = 64 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			math.Float64frombits(bin.Uint64(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
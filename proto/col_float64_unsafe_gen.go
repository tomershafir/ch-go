//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColFloat64 represents Float64 column.
type ColFloat64 []float64

// Compile-time assertions for ColFloat64.
var (
	_ ColInput  = ColFloat64{}
	_ ColResult = (*ColFloat64)(nil)
	_ Column    = (*ColFloat64)(nil)
)

// Type returns ColumnType of Float64.
func (ColFloat64) Type() ColumnType {
	return ColumnTypeFloat64
}

// Rows returns count of rows in column.
func (c ColFloat64) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColFloat64) Row(i int) float64 {
	return c[i]
}

// Append float64 to column.
func (c *ColFloat64) Append(v float64) {
	*c = append(*c, v)
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
		binary.LittleEndian.PutUint64(
			b.Buf[offset:offset+size],
			math.Float64bits(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Float64 rows from *Reader.
func (c *ColFloat64) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 64 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	v = append(v, make([]float64, rows)...)
	s := *(*struct {
		Data unsafe.Pointer
		Len  uintptr
		Cap  uintptr
	})(unsafe.Pointer(&v))
	s.Len *= size
	s.Cap *= size
	dst := *(*[]byte)(unsafe.Pointer(&s))
	copy(dst, data)
	*c = v
	return nil
}

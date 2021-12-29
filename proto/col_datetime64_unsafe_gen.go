//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"unsafe"

	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColDateTime64 represents DateTime64 column.
type ColDateTime64 []DateTime64

// Compile-time assertions for ColDateTime64.
var (
	_ ColInput  = ColDateTime64{}
	_ ColResult = (*ColDateTime64)(nil)
	_ Column    = (*ColDateTime64)(nil)
)

// Type returns ColumnType of DateTime64.
func (ColDateTime64) Type() ColumnType {
	return ColumnTypeDateTime64
}

// Rows returns count of rows in column.
func (c ColDateTime64) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDateTime64) Row(i int) DateTime64 {
	return c[i]
}

// Append DateTime64 to column.
func (c *ColDateTime64) Append(v DateTime64) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDateTime64) Reset() {
	*c = (*c)[:0]
}

// NewArrDateTime64 returns new Array(DateTime64).
func NewArrDateTime64() *ColArr {
	return &ColArr{
		Data: new(ColDateTime64),
	}
}

// AppendDateTime64 appends slice of DateTime64 to Array(DateTime64).
func (c *ColArr) AppendDateTime64(data []DateTime64) {
	d := c.Data.(*ColDateTime64)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes DateTime64 rows to *Buffer.
func (c ColDateTime64) EncodeColumn(b *Buffer) {
	const size = 64 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint64(
			b.Buf[offset:offset+size],
			uint64(v),
		)
		offset += size
	}
}

// DecodeColumn decodes DateTime64 rows from *Reader.
func (c *ColDateTime64) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 64 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	v = append(v, make([]DateTime64, rows)...)
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

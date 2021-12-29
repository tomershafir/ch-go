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

// ColEnum16 represents Enum16 column.
type ColEnum16 []Enum16

// Compile-time assertions for ColEnum16.
var (
	_ ColInput  = ColEnum16{}
	_ ColResult = (*ColEnum16)(nil)
	_ Column    = (*ColEnum16)(nil)
)

// Type returns ColumnType of Enum16.
func (ColEnum16) Type() ColumnType {
	return ColumnTypeEnum16
}

// Rows returns count of rows in column.
func (c ColEnum16) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColEnum16) Row(i int) Enum16 {
	return c[i]
}

// Append Enum16 to column.
func (c *ColEnum16) Append(v Enum16) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColEnum16) Reset() {
	*c = (*c)[:0]
}

// NewArrEnum16 returns new Array(Enum16).
func NewArrEnum16() *ColArr {
	return &ColArr{
		Data: new(ColEnum16),
	}
}

// AppendEnum16 appends slice of Enum16 to Array(Enum16).
func (c *ColArr) AppendEnum16(data []Enum16) {
	d := c.Data.(*ColEnum16)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Enum16 rows to *Buffer.
func (c ColEnum16) EncodeColumn(b *Buffer) {
	const size = 16 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint16(
			b.Buf[offset:offset+size],
			uint16(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Enum16 rows from *Reader.
func (c *ColEnum16) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 16 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	v = append(v, make([]Enum16, rows)...)
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

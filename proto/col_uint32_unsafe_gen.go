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

// ColUInt32 represents UInt32 column.
type ColUInt32 []uint32

// Compile-time assertions for ColUInt32.
var (
	_ ColInput  = ColUInt32{}
	_ ColResult = (*ColUInt32)(nil)
	_ Column    = (*ColUInt32)(nil)
)

// Type returns ColumnType of UInt32.
func (ColUInt32) Type() ColumnType {
	return ColumnTypeUInt32
}

// Rows returns count of rows in column.
func (c ColUInt32) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColUInt32) Row(i int) uint32 {
	return c[i]
}

// Append uint32 to column.
func (c *ColUInt32) Append(v uint32) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt32) Reset() {
	*c = (*c)[:0]
}

// NewArrUInt32 returns new Array(UInt32).
func NewArrUInt32() *ColArr {
	return &ColArr{
		Data: new(ColUInt32),
	}
}

// AppendUInt32 appends slice of uint32 to Array(UInt32).
func (c *ColArr) AppendUInt32(data []uint32) {
	d := c.Data.(*ColUInt32)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes UInt32 rows to *Buffer.
func (c ColUInt32) EncodeColumn(b *Buffer) {
	const size = 32 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint32(
			b.Buf[offset:offset+size],
			v,
		)
		offset += size
	}
}

// DecodeColumn decodes UInt32 rows from *Reader.
func (c *ColUInt32) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 32 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	v = append(v, make([]uint32, rows)...)
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

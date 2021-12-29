//go:build !amd64 || nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"

	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColDate32 represents Date32 column.
type ColDate32 []Date32

// Compile-time assertions for ColDate32.
var (
	_ ColInput  = ColDate32{}
	_ ColResult = (*ColDate32)(nil)
	_ Column    = (*ColDate32)(nil)
)

// Type returns ColumnType of Date32.
func (ColDate32) Type() ColumnType {
	return ColumnTypeDate32
}

// Rows returns count of rows in column.
func (c ColDate32) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDate32) Row(i int) Date32 {
	return c[i]
}

// Append Date32 to column.
func (c *ColDate32) Append(v Date32) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDate32) Reset() {
	*c = (*c)[:0]
}

// NewArrDate32 returns new Array(Date32).
func NewArrDate32() *ColArr {
	return &ColArr{
		Data: new(ColDate32),
	}
}

// AppendDate32 appends slice of Date32 to Array(Date32).
func (c *ColArr) AppendDate32(data []Date32) {
	d := c.Data.(*ColDate32)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Date32 rows to *Buffer.
func (c ColDate32) EncodeColumn(b *Buffer) {
	const size = 32 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint32(
			b.Buf[offset:offset+size],
			uint32(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Date32 rows from *Reader.
func (c *ColDate32) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 32 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			Date32(binary.LittleEndian.Uint32(data[i:i+size])),
		)
	}
	*c = v
	return nil
}

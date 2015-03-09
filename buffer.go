// lifo bytes buffer
//
//     /*   For exapmle:   */
//     b := NewBuffer(nil)
//     b.WriteByte(byte(1))
//     b.WriteByte(byte(2))
//     b.WriteByte(byte(3))
//     b.ReadByte // => 3
//     b.ReadByte // => 2
//     b.ReadByte // => 1
package lifo

import (
	"errors"
	"io"
)

// LIFO bytes buffer
type Buffer []byte

// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("lifo.Buffer: to large")

// ErrInvalidWriteCount is returned when WriteTo's writer returns wrong count
var ErrInvalidWriteCount = errors.New("lifo.Buffer.WriteTo: invalid Write count")

// NewBuffer returns new buffer with pre-defined buffer or nil
func NewBuffer(p []byte) *Buffer {
	b := new(Buffer)
	*b = p
	return b
}

// Read reads the next len(p) bytes from the buffer or until the buffer is drained.
// The return value n is the number of bytes read. If the buffer has no data to
// return, err is io.EOF (unless len(p) is zero); otherwise it is nil.
func (b *Buffer) Read(p []byte) (n int, err error) {
	if x := len(*b) - len(p); x >= 0 {
		n = copy(p, (*b)[x:])
		*b = (*b)[:x]
		return
	}
	n = copy(p, *b)
	*b = nil
	return n, io.EOF
}

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p. If the
// buffer becomes too large, err will be ErrTooLarge.
func (b *Buffer) Write(p []byte) (n int, err error) {
	(*b) = append(*b, p...)
	n = len(p)
	defer func() {
		if x := recover(); x != nil {
			if x == ErrTooLarge {
				err = x.(error)
			} else {
				panic(x)
			}
		}
	}()
	return
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of bytes written; it always fits into an
// int, but it is int64 to match the io.WriterTo interface. Any error
// encountered during the write is also returned. If w returns invalid count
// err will be ErrInvalidWriteCount
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	if lenb := len(*b); lenb > 0 {
		m, e := w.Write(*b)
		n = int64(m)
		if m > lenb {
			return n, ErrInvalidWriteCount
		}
		*b = (*b)[:lenb-m]
		if e != nil {
			return n, e
		}
		if m != lenb {
			return n, io.ErrShortWrite
		}
	}
	return
}

// ReadByte reads and returns the next byte from the buffer.
// If no byte is available, it returns error io.EOF.
func (b *Buffer) ReadByte() (c byte, err error) {
	if len(*b) > 0 {
		c, *b = (*b)[len(*b)-1], (*b)[:len(*b)-1]
		return
	}
	return c, io.EOF
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// If the buffer becomes too large, err will be ErrTooLarge.
func (b *Buffer) WriteByte(c byte) (err error) {
	*b = append(*b, c)
	defer func() {
		if x := recover(); x != nil {
			if x == ErrTooLarge {
				err = x.(error)
			} else {
				panic(x)
			}
		}
	}()
	return
}

// Len returns the length of the buffer
func (b *Buffer) Len() int {
	return len(*b)
}

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
func (b *Buffer) Next(n int) (p []byte) {
	if x := len(*b) - n; x >= 0 {
		p = make([]byte, n)
		n = copy(p, (*b)[x:])
		*b = (*b)[:x]
		return p
	} else {
		p = *b
		*b = nil
		return p
	}
}

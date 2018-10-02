//
// Copyright (c) 2016 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

// Package lifo implements bytes buffer with LIFO order (last-in-first-out)
// part-by-part. So, if you write `[]byte("Hello")` and then
// `[]byte("World")`. After that you can read it back: `"World"`
// then `"Hello"`.
package lifo

import (
	"errors"
	"io"
)

// Buffer is LIFO bytes buffer type
type Buffer []byte

// ErrInvalidWriteCount is returned when WriteTo's writer returns wrong
// count
var ErrInvalidWriteCount = errors.New("lifo.Buffer.WriteTo:" +
	" invalid Write count")

// NewBuffer returns new buffer with pre-defined buffer or nil
func NewBuffer(p []byte) *Buffer {
	b := new(Buffer)
	*b = p
	return b
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the buffer
// has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
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
// needed. The return value n is the length of p.
// The err will be always nil
func (b *Buffer) Write(p []byte) (n int, err error) {
	(*b) = append(*b, p...)
	n = len(p)
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
// It always returns nil
func (b *Buffer) WriteByte(c byte) (err error) {
	*b = append(*b, c)
	return
}

// Len returns the length of the buffer
func (b *Buffer) Len() int {
	return len(*b)
}

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next
// returns the entire buffer.
func (b *Buffer) Next(n int) (p []byte) {
	if x := len(*b) - n; x >= 0 {
		p = make([]byte, n)
		copy(p, (*b)[x:])
		*b = (*b)[:x]
		return p
	}
	p = *b
	*b = nil
	return p
}

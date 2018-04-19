lifo
====

[![GoDoc](https://godoc.org/github.com/logrusorgru/lifo?status.svg)](https://godoc.org/github.com/logrusorgru/lifo)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/lifo.svg)](https://travis-ci.org/logrusorgru/lifo)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/lifo/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/lifo?branch=master)
[![GoReportCard](https://goreportcard.com/badge/logrusorgru/lifo)](https://goreportcard.com/report/logrusorgru/lifo)
[![Gitter](https://img.shields.io/badge/chat-on_gitter-46bc99.svg?logo=data:image%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGhlaWdodD0iMTQiIHdpZHRoPSIxNCI%2BPGcgZmlsbD0iI2ZmZiI%2BPHJlY3QgeD0iMCIgeT0iMyIgd2lkdGg9IjEiIGhlaWdodD0iNSIvPjxyZWN0IHg9IjIiIHk9IjQiIHdpZHRoPSIxIiBoZWlnaHQ9IjciLz48cmVjdCB4PSI0IiB5PSI0IiB3aWR0aD0iMSIgaGVpZ2h0PSI3Ii8%2BPHJlY3QgeD0iNiIgeT0iNCIgd2lkdGg9IjEiIGhlaWdodD0iNCIvPjwvZz48L3N2Zz4%3D&logoWidth=10)](https://gitter.im/logrusorgru/lifo?utm_source=share-link&utm_medium=link&utm_campaign=share-link)

Golang lifo reader/writer buffer like truncated bytes.Buffer. The LIFO
means part-by-part. If you write `[]byte("Hello")` and then `[]byte("World")`.
After that you can read it back: `"World"` then `"Hello"`.

### Install

```bash
go get github.com/logrusorgru/lifo
cd $GOPATH/src/github.com/logrusorgru/lifo
go test
```

### Methods

- Read
- Write
- WriteTo
- ReadByte
- WriteByte
- Len
- Next

### Example

```go
package main

import (
	"fmt"
	"github.com/logrusorgru/lifo"
)

func main() {
	b := lifo.NewBuffer(nil)

	b.Write([]byte("bye"))
	b.Write([]byte("hello"))

	hello := make([]byte, 5)
	bye := make([]byte, 3)

	b.Read(hello)
	b.Read(bye)

	fmt.Println(string(hello))
	fmt.Println(string(bye))

	// Output
	// hello
	// bye
}
```

### Licensing

Copyright &copy; 2016 Konstantin Ivanov <kostyarin.ivanov@gmail.com>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
